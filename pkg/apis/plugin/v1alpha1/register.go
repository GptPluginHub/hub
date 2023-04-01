package v1alpha1

import (
	"context"

	"github.com/GptPluginHub/hub/pkg/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pluginv1alpha1 "hub.io/api/plugin/v1alpha1"
	"k8s.io/klog"
)

func RegisterGrpcService(ctx context.Context, cfg config.Config, grpcServer *grpc.Server) {
	klog.V(3).Info("register plugin grpc services")
	handler := NewPluginHandler(cfg)
	pluginv1alpha1.RegisterPluginServiceServer(grpcServer, handler)
}

func RegisterGatewayService(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	err := pluginv1alpha1.RegisterPluginServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		panic(err)
	}
}
