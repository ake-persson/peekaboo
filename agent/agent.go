package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/mitchellh/go-homedir"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/yaml.v2"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

var version = "undefined"

type server struct {
	logger *zap.Logger
}

type config struct {
	InfoFile    string
	Addr        string
	CatalogAddr string
	KeepAlive   int
	Register    bool
	NoTLS       bool
	MTLS        bool // TBD
	CertFile    string
	KeyFile     string
	CAFile      string
	Level       zapcore.Level // TBD
}

type info struct {
	Description  string `yaml:"description"`
	Site         string `yaml:"site"`
	Rack         string `yaml:"rack"`
	RackPosition int32  `yaml:"rackPosition"`
	RackSize     int32  `yaml:"rackSize"`
}

func main() {
	// Setup config and flags.
	conf := &config{}
	var printVersion bool
	flag.StringVar(&conf.InfoFile, "info-file", "~/etc/peekaboo.yml", "System info file")
	flag.StringVar(&conf.Addr, "addr", ":17711", "Server address")
	flag.StringVar(&conf.CatalogAddr, "ctlg-addr", "catalog:28657", "Catalog address")
	flag.IntVar(&conf.KeepAlive, "keepalive", 60*5, "Keepalive intervall in seconds")
	flag.BoolVar(&conf.Register, "register", false, "Register with Catalog")
	flag.BoolVar(&conf.NoTLS, "no-tls", false, "No TLS (testing)")
	flag.BoolVar(&conf.MTLS, "mtls", false, "Use MTLS") // TBD
	flag.StringVar(&conf.CertFile, "cert-file", "~/certs/srv.crt", "Server TLS certificate file")
	flag.StringVar(&conf.KeyFile, "key-file", "~/certs/srv.key", "Server TLS key file")
	flag.StringVar(&conf.CAFile, "ca-file", "~/certs/root_ca.crt", "CA certificate file, required for Mutual TLS")
	flag.BoolVar(&printVersion, "version", false, "Version")
	flag.Parse()

	if printVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	// Replace tilde with home directory.
	conf.InfoFile, _ = homedir.Expand(conf.InfoFile)
	conf.CertFile, _ = homedir.Expand(conf.CertFile)
	conf.KeyFile, _ = homedir.Expand(conf.KeyFile)
	conf.CAFile, _ = homedir.Expand(conf.CAFile)

	// Setup logger.
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Load info file if it exist's.
	i := &info{}
	if _, err := os.Stat(conf.InfoFile); !os.IsNotExist(err) {
		b, err := ioutil.ReadFile(conf.InfoFile)
		if err != nil {
			logger.Fatal("info file", zap.Error(err))
		}

		if err = yaml.Unmarshal(b, i); err != nil {
			logger.Fatal("unmarshal", zap.Error(err))
		}
	} else {
		logger.Debug("skipping info file since it's doesn't exist", zap.String("file", conf.InfoFile))
	}

	// Register with Catalog.
	if conf.Register {
		if err := registerSystem(conf); err != nil {
			logger.Fatal("failed to register",
				zap.Error(err))
		}
	}

	// Setup server options.
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	}

	// Setup TLS server options.
	if !conf.NoTLS {
		creds, err := credentials.NewServerTLSFromFile(conf.CertFile, conf.KeyFile)
		if err != nil {
			logger.Fatal("failed to load certificates",
				zap.Error(err),
				zap.String("certificate", conf.CertFile),
				zap.String("key", conf.KeyFile))
		}

		opts = append(opts, grpc.Creds(creds))
	}

	// New gRPC server.
	s := grpc.NewServer(opts...)
	services.RegisterSystemServiceServer(s, &server{logger: logger})

	lis, err := net.Listen("tcp", conf.Addr)
	if err != nil {
		logger.Fatal("failed to listen",
			zap.Error(err),
			zap.String("address", conf.Addr))
	}

	logger.Debug("serve grpc",
		zap.String("address", conf.Addr))
	logger.Fatal("serve grpc", zap.Error(s.Serve(lis)))
}
