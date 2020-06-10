package serve

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

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

func Run(args []string, stdout io.Writer, opts *Options) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.SetOutput(stdout)
	flags.Usage = usage(flags, stdout)
	var (
		addr = flags.String("addr", "localhost:17711", "Server address")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	// Setup logger.
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Setup server options.
	srvrOpts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	}

	// Setup TLS server options.
	if !opts.NoTLS {
		creds, err := credentials.NewServerTLSFromFile(opts.CertFile, opts.KeyFile)
		if err != nil {
			logger.Fatal("failed to load certificates",
				zap.Error(err),
				zap.String("certificate", opts.CertFile),
				zap.String("key", opts.KeyFile))
		}

		srvrOpts = append(srvrOpts, grpc.Creds(creds))
	}

	// New gRPC server.
	s := grpc.NewServer(srvrOpts...)
	services.RegisterSystemServiceServer(s, &server{logger: logger})

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		logger.Fatal("failed to listen",
			zap.Error(err),
			zap.String("address", *addr))
	}

	logger.Debug("serve grpc",
		zap.String("address", *addr))
	logger.Fatal("serve grpc", zap.Error(s.Serve(lis)))

	return nil
}
