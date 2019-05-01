package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strings"

	remoteexecution "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	"github.com/buildbarn/bb-storage/pkg/ac"
	"github.com/buildbarn/bb-storage/pkg/blobstore/configuration"
	"github.com/buildbarn/bb-storage/pkg/builder"
	"github.com/buildbarn/bb-storage/pkg/cas"
	"github.com/buildbarn/bb-storage/pkg/util"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	prometheus_exporter "contrib.go.opencensus.io/exporter/prometheus"
	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"go.opencensus.io/zpages"

	"google.golang.org/genproto/googleapis/bytestream"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	var (
		blobstoreConfig      = flag.String("blobstore-config", "/config/blobstore.conf", "Configuration for blob storage")
		webListenAddress     = flag.String("web.listen-address", ":80", "Port on which to expose metrics")
		agentEndpointURI     = flag.String("jaeger.agent-endpoint", "127.0.0.1:6831", "jaeger agent address")
		collectorEndpointURI = flag.String("jaeger.collector-endpoint", "http://127.0.0.1:14268/api/traces", "jaeger collector endpoint")
		serviceName          = flag.String("service-name", "bb_storage", "service name for tracing")
		alwaysSample         = flag.Bool("always-sample", false, "record all traces. warning! do not use this in production!")
	)
	var schedulersList util.StringList
	flag.Var(&schedulersList, "scheduler", "Backend capable of executing build actions. Example: debian8|hostname-of-debian8-scheduler:8981")
	var allowActionCacheUpdatesForInstancesList util.StringList
	flag.Var(&allowActionCacheUpdatesForInstancesList, "allow-ac-updates-for-instance", "Allow clients to write into the action cache for this instance")
	flag.Parse()

	// Web server for metrics and profiling.
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(*webListenAddress, nil))
	}()

	pe, err := prometheus_exporter.NewExporter(prometheus_exporter.Options{
		Namespace: "bb_storage",
		Registry: prometheus.DefaultRegisterer.(*prometheus.Registry),
	})
	if err != nil {
		log.Fatalf("Failed to create the Prometheus stats exporter: %v", err)
	}
	view.RegisterExporter(pe)
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatalf("Failed to register ocgrpc server views: %v", err)
	}
	zpages.Handle(nil, "/debug")
	je, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint:          *agentEndpointURI,
		CollectorEndpoint:      *collectorEndpointURI,
		ServiceName:            *serviceName,
	})
	if err != nil {
		log.Fatalf("Failed to create the Jaeger exporter: %v", err)
	}

	trace.RegisterExporter(je)
	if *alwaysSample {
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	}

	// Storage access.
	contentAddressableStorageBlobAccess, actionCacheBlobAccess, err := configuration.CreateBlobAccessObjectsFromConfig(*blobstoreConfig)
	if err != nil {
		log.Fatal("Failed to create blob access: ", err)
	}
	actionCache := ac.NewBlobAccessActionCache(actionCacheBlobAccess)

	// Let GetCapabilities() work, even for instances that don't
	// have a scheduler attached to them, but do allow uploading
	// results into the Action Cache.
	schedulers := map[string]builder.BuildQueue{}
	allowActionCacheUpdatesForInstances := map[string]bool{}
	if len(allowActionCacheUpdatesForInstancesList) > 0 {
		fallback := builder.NewNonExecutableBuildQueue()
		for _, instance := range allowActionCacheUpdatesForInstancesList {
			schedulers[instance] = fallback
			allowActionCacheUpdatesForInstances[instance] = true
		}
	}

	// Backends capable of compiling.
	for _, schedulerEntry := range schedulersList {
		components := strings.SplitN(schedulerEntry, "|", 2)
		if len(components) != 2 {
			log.Fatal("Invalid scheduler entry: ", schedulerEntry)
		}
		scheduler, err := grpc.Dial(
			components[1],
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
			grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor))
		if err != nil {
			log.Fatal("Failed to create scheduler RPC client: ", err)
		}
		schedulers[components[0]] = builder.NewForwardingBuildQueue(scheduler)
	}
	buildQueue := builder.NewDemultiplexingBuildQueue(func(instance string) (builder.BuildQueue, error) {
		scheduler, ok := schedulers[instance]
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "Unknown instance name")
		}
		return scheduler, nil
	})

	// RPC server.
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
		grpc.StatsHandler(&ocgrpc.ServerHandler{}),
	)
	remoteexecution.RegisterActionCacheServer(s, ac.NewActionCacheServer(actionCache, allowActionCacheUpdatesForInstances))
	remoteexecution.RegisterContentAddressableStorageServer(s, cas.NewContentAddressableStorageServer(contentAddressableStorageBlobAccess))
	bytestream.RegisterByteStreamServer(s, cas.NewByteStreamServer(contentAddressableStorageBlobAccess, 1<<16))
	remoteexecution.RegisterCapabilitiesServer(s, buildQueue)
	remoteexecution.RegisterExecutionServer(s, buildQueue)
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(s)

	sock, err := net.Listen("tcp", ":8980")
	if err != nil {
		log.Fatal("Failed to create listening socket: ", err)
	}
	if err := s.Serve(sock); err != nil {
		log.Fatal("Failed to serve RPC server: ", err)
	}
}
