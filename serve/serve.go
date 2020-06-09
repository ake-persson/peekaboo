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

type server struct {
	logger *zap.Logger
}

type Options struct {
	NoTLS    bool
	NoVerify bool
	NoMTLS   bool
	CertFile string
	KeyFile  string
	CAFile   string
}

func usage(flags *flag.FlagSet, stdout io.Writer) func() {
        return func() {
                fmt.Fprintf(stdout, `Usage: %s serve [OPTIONS]

Serve API

Options:
`, os.Args[0])
                flags.PrintDefaults()
        }
}

func Run(args []string, opts *Options) error {
	var addr := flags.String("addr", "localhost:17711", "Server address")

        flags := flag.NewFlagSet(args, flag.ExitOnError)
        flags.Usage = usage(flags, stdout)
        var (
	addr := flags.String("addr", "localhost:17711", "Server address")
        )
        flags.Usage = usage(flags, stdout)
        if err := flags.Parse(args); err != nil {
                return err
        }

	// Replace tilde with home directory.
	conf.CertFile, _ = homedir.Expand(conf.CertFile)
	conf.KeyFile, _ = homedir.Expand(conf.KeyFile)
	conf.CAFile, _ = homedir.Expand(conf.CAFile)

	// Setup logger.
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

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
