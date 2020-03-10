package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

type config struct {
	Addr     string
	NoTLS    bool
	MTLS     bool // TBD
	CertFile string
	KeyFile  string
	CAFile   string
}

var version = "undefined"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] <resource>\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	// Setup config and flags.
	conf := &config{}
	var printVersion bool
	flag.Usage = usage
	flag.StringVar(&conf.Addr, "addr", "localhost:17711", "Catalog address")
	flag.BoolVar(&conf.NoTLS, "no-tls", false, "No TLS (testing)")
	flag.BoolVar(&conf.MTLS, "mtls", false, "MTLS") // TBD
	flag.StringVar(&conf.CertFile, "cert-file", "~/certs/srv.crt", "Server TLS certificate file")
	flag.StringVar(&conf.KeyFile, "key-file", "~/certs/srv.key", "Server TLS key file")
	flag.StringVar(&conf.CAFile, "ca-file", "~/certs/root_ca.crt", "CA certificate file, required for Mutual TLS")
	flag.BoolVar(&printVersion, "version", false, "Version")
	flag.Parse()

	if printVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	// Get resource.
	if len(flag.Args()) < 1 {
		usage()
		os.Exit(1)
	}
	resource := flag.Args()[0]

	// Replace tilde with home directory.
	conf.CertFile, _ = homedir.Expand(conf.CertFile)
	conf.KeyFile, _ = homedir.Expand(conf.KeyFile)
	conf.CAFile, _ = homedir.Expand(conf.CAFile)

	opts := []grpc.DialOption{grpc.WithBlock()}
	if conf.NoTLS {
		opts = append(opts, grpc.WithInsecure())
	} else {
		// Load CA certificate file.
		creds, err := credentials.NewClientTLSFromFile(conf.CAFile, "")
		if err != nil {
			log.Fatalf("failed to load ca certificate: %v", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	// Connect to gRPC server.
	conn, err := grpc.Dial(conf.Addr, opts...)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Initialize new client.
	client := services.NewSystemServiceClient(conn)

	// Create context for gRPC calls.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var v interface{}
	switch resource {
	case "system":
		v, err = client.GetSystem(ctx, &services.GetSystemRequest{})
	case "users":
		v, err = client.ListUsers(ctx, &services.ListUsersRequest{})
	}
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
