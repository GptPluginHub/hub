package options

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GptPluginHub/hub/pkg/apisever"
	"github.com/GptPluginHub/hub/pkg/config"
	"github.com/GptPluginHub/hub/pkg/middleware"

	"github.com/gorilla/mux"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Options struct {
	// config file
	ConfigFile string
	// server options
	ServerOptions *ServerOptions
	// config
	Conf *config.Config
}

type ServerOptions struct {
	// server bind address
	BindAddress string
	// insecure port number
	InsecurePort int
	// secure port number
	SecurePort int
	// tls cert file
	TLSCertFile string
	// tls private key file
	TLSPrivateKey string
}

func NewDefaultServerOption() *ServerOptions {
	return &ServerOptions{
		BindAddress:   "0.0.0.0",
		InsecurePort:  8000,
		SecurePort:    0,
		TLSCertFile:   "",
		TLSPrivateKey: "",
	}
}

func (s *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.BindAddress, "bind-address", s.BindAddress, "server bind address")
	fs.IntVar(&s.InsecurePort, "insecure-port", s.InsecurePort, "insecure port number")
	fs.IntVar(&s.SecurePort, "secure-port", s.SecurePort, "secure port number")
	fs.StringVar(&s.TLSCertFile, "tls-cert-file", s.TLSCertFile, "tls cert file")
	fs.StringVar(&s.TLSPrivateKey, "tls-private-key", s.TLSPrivateKey, "tls private key")
}

func (o *Options) Flags() (fs *pflag.FlagSet) {
	flagSet := pflag.NewFlagSet("", pflag.ExitOnError)
	flagSet.StringVar(&o.ConfigFile, "config", o.ConfigFile, "config file")
	o.ServerOptions.AddFlags(flagSet)
	o.Conf.AddFlags(flagSet)
	return flagSet
}

func NewDefaultOption() *Options {
	return &Options{
		ServerOptions: NewDefaultServerOption(),
		Conf:          config.NewConfig(),
	}
}

// NewServer creates a new server.
func (o *Options) NewServer(ctx context.Context) *apisever.Server {
	address := fmt.Sprintf("%s:%d", o.ServerOptions.BindAddress, o.ServerOptions.InsecurePort)
	s := &apisever.Server{
		Config: o.Conf,
		HTTPServer: &http.Server{
			Addr: address,
		},
		GRPCServer: grpc.NewServer(grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			grpcrecovery.StreamServerInterceptor(),
		)),
			grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer())),
		GatewayServerMux: runtime.NewServeMux(),
		Router:           mux.NewRouter(),
	}
	// Register reflection service on gRPC server.
	reflection.Register(s.GRPCServer)
	// add swagger handler
	s.AddSwaggerHandler()
	// install apis
	s.InstallHubApis(ctx)
	// add openapi handler
	s.AddOpenapiHandler(ctx)
	// add http cost handler
	s.Router.Use(middleware.LogRequestAndResponse)
	s.HTTPServer.Handler = s.Router
	return s
}
