package apisever

import (
	"context"
	"net"
	"net/http"
	"os"
	"path"

	apisopenapi "github.com/GptPluginHub/hub/pkg/apis/openapi/v1alpha1"
	apisplugin "github.com/GptPluginHub/hub/pkg/apis/plugin/v1alpha1"
	"github.com/GptPluginHub/hub/pkg/config"

	"github.com/cockroachdb/cmux"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"k8s.io/klog"
)

const (
	SwaggerOpenAPIURLPrefix = "/swagger/"
	SwaggerRootPath         = "api/assets/swagger"
)

type Server struct {
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
	if !s.Config.Debug {
		klog.Warning("Swagger is not enabled in production mode")
		return
	}
	dir, err := os.Getwd()
	if err != nil {
		klog.Errorf("get current directory: %v", err)
		panic(err)
	}
	assetsDir := path.Join(dir, SwaggerRootPath)
	klog.Infof("Serving swagger files in path: %s", assetsDir)
	openapiServer := http.FileServer(http.Dir(assetsDir))
	s.Router.PathPrefix(SwaggerOpenAPIURLPrefix).Handler(http.StripPrefix(SwaggerOpenAPIURLPrefix, openapiServer))
}

func (s *Server) InstallHubApis(ctx context.Context) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	apisplugin.RegisterGrpcService(ctx, *s.Config, s.GRPCServer)
	apisplugin.RegisterGatewayService(ctx, s.GatewayServerMux, s.HTTPServer.Addr, opts)
	s.Router.PathPrefix("/apis/").Handler(s.GatewayServerMux)
	if s.Config.Debug {
		s.Router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)
	}
	s.Router.HandleFunc("/healthz", livenessProbe).Methods(http.MethodGet)
	s.Router.HandleFunc("/readyz", readinessProbe).Methods(http.MethodGet)
}

func (s *Server) AddOpenapiHandler(ctx context.Context) {
	openAPIHandler := apisopenapi.NewOpenAPIHandler(s.Config)
	s.Router.HandleFunc("/api/hub.io/v1alpha1/openapi", openAPIHandler.OpenAPIHandler)
	s.Router.HandleFunc("/api/hub.io/v1alpha1/openapi/data", openAPIHandler.OpenAPIHandlerData)
	// TODO this is a temporary solution
	s.Router.HandleFunc("/api/init", openAPIHandler.IninPluginMetadata)
}

func (s *Server) Run(ctx context.Context) error {
	klog.Infof("Starting server on %s", s.HTTPServer.Addr)

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

// With `livenessProbe`, apiserver would not be available if the back-end services are not ready.
func livenessProbe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

// With `readinessProbe`, apiserver would not accept traffics if the back-end services are not ready.
func readinessProbe(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
