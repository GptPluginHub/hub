package apisever

import (
	"context"
	"net"
	"net/http"
	"os"
	"path"

	apisplugin "github.com/GptPluginHub/hub/pkg/apis/plugin/v1alpha1"
	"github.com/GptPluginHub/hub/pkg/config"

	"github.com/cockroachdb/cmux"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"k8s.io/klog"
)

const (
	SwaggerOpenAPIURLPrefix = "/swagger/"
	SwaggerRootPath         = "api/assets/swagger"
)

type Server struct {
	// debug mode
	Debug bool
	// config
	Config *config.Config
	// http server
	HTTPServer *http.Server
	// grpc server
	GRPCServer *grpc.Server
	// grpc gateway server
	GatewayServerMux *runtime.ServeMux
	// cmux
	CMux cmux.CMux
	// Router
	Router *mux.Router
}

func (s *Server) AddSwaggerHandler() {
	dir, err := os.Getwd()
	if err != nil {
		klog.Errorf("get current directory: %v", err)
		panic(err)
	}
	assetsDir := path.Join(dir, SwaggerRootPath)
	klog.Infof("Serving swagger files in path: %s.", assetsDir)
	openapiServer := http.FileServer(http.Dir(assetsDir))
	s.Router.PathPrefix(SwaggerOpenAPIURLPrefix).Handler(http.StripPrefix(SwaggerOpenAPIURLPrefix, openapiServer))
}

func (s *Server) InstallHubApis(ctx context.Context) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	apisplugin.RegisterGrpcService(ctx, *s.Config, s.GRPCServer)
	apisplugin.RegisterGatewayService(ctx, s.GatewayServerMux, s.HTTPServer.Addr, opts)
}

func (s *Server) Run(ctx context.Context) error {
	klog.Infof("Starting server on %s", s.HTTPServer.Addr)

	// Register reflection service on gRPC server.
	reflection.Register(s.GRPCServer)

	l, err := net.Listen("tcp", s.HTTPServer.Addr)
	if err != nil {
		return err
	}
	s.CMux = cmux.New(l)

	grpcListener := s.CMux.Match(cmux.HTTP2())
	httpLister := s.CMux.Match(cmux.HTTP1Fast())

	go func() {
		if err = s.GRPCServer.Serve(grpcListener); err != nil {
			klog.Errorf("grpc server serve error: %v", err)
			panic(err)
		}
	}()
	go func() {
		if err = s.HTTPServer.Serve(httpLister); err != nil {
			klog.Errorf("http server serve error: %v", err)
			panic(err)
		}
	}()

	return s.CMux.Serve()
}
