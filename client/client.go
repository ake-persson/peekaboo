package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/peekaboo-labs/peekaboo/pkg/filesystem"
	"github.com/peekaboo-labs/peekaboo/pkg/group"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
	"github.com/peekaboo-labs/peekaboo/pkg/system"
	"github.com/peekaboo-labs/peekaboo/pkg/text"
	"github.com/peekaboo-labs/peekaboo/pkg/user"
)

// TODO
// - Set field order using flag
// - No color option
// - Define color option

type config struct {
	NoTLS    bool
	MTLS     bool // TBD
	CertFile string
	KeyFile  string
	CAFile   string
	Format   string
	Fields   string
}

type envelope struct {
	Address  string      `json:"address"`
	Response interface{} `json:"response,omitempty"`
	Error    error       `json:"error,omitempty"`
}

var version = "undefined"

func splitOmitEmpty(s string, del string) []string {
	out := []string{}
	for _, v := range strings.Split(s, del) {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] <resource> <address...>\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `  resource
    	Resource to query [system,users,groups,filesystems]
  address
        Address to agent specified as <address[:port]> (default port 17711)
`)
}

func main() {
	// Setup config and flags.
	conf := &config{}
	var printVersion bool
	flag.Usage = usage
	flag.BoolVar(&conf.NoTLS, "no-tls", false, "No TLS (testing)")
	flag.BoolVar(&conf.MTLS, "mtls", false, "Use MTLS") // TBD
	flag.StringVar(&conf.CertFile, "cert-file", "~/certs/srv.crt", "Server TLS certificate file")
	flag.StringVar(&conf.KeyFile, "key-file", "~/certs/srv.key", "Server TLS key file")
	flag.StringVar(&conf.CAFile, "ca-file", "~/certs/root_ca.crt", "CA certificate file, required for Mutual TLS")
	flag.StringVar(&conf.Format, "fmt", "json", "Output format [json,csv,table,vtable]")
	flag.StringVar(&conf.Fields, "fields", "", "Comma separate list of fields to output")
	flag.BoolVar(&printVersion, "version", false, "Version")
	flag.Parse()

	if printVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	if len(flag.Args()) < 2 {
		usage()
		os.Exit(1)
	}

	// Positional arguments.
	resource := flag.Args()[0]
	addresses := flag.Args()[1:]

	// Replace tilde with home directory.
	conf.CertFile, _ = homedir.Expand(conf.CertFile)
	conf.KeyFile, _ = homedir.Expand(conf.KeyFile)
	conf.CAFile, _ = homedir.Expand(conf.CAFile)

	// Check resource.
	if !text.InList(resource, []string{"system", "users", "groups", "filesystems"}) {
		log.Fatalf("unknown resource: %s", resource)
	}

	// Check format.
	if !text.InList(conf.Format, []string{"json", "csv", "table", "vtable"}) {
		log.Fatalf("unknown format: %s", conf.Format)
	}

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

	tables := text.Tables{}
	responses := []interface{}{}
	for _, addr := range addresses {
		if !strings.Contains(addr, ":") {
			addr += ":17711"
		}

		switch conf.Format {
		case "json":
			resp, err := dialAgent(resource, addr, opts)
			if err != nil {
				log.Print(err)
			}
			responses = append(responses, resp)
		case "csv", "table", "vtable":
			t, err := dialAgentTable(resource, addr, opts)
			if err != nil {
				log.Print(err)
			}
			tables = append(tables, t)
		default:
			log.Fatalf("unknown output format: %s", conf.Format)
		}
	}

	switch conf.Format {
	case "json":
		b, _ := json.MarshalIndent(responses, "", "  ")
		fmt.Print(string(b))
	case "csv":
		tables.PrintCSV(os.Stdout, splitOmitEmpty(conf.Fields, ","))
	case "table":
		tables.PrintTable(os.Stdout, splitOmitEmpty(conf.Fields, ","))
	case "vtable":
		tables.PrintVertTable(os.Stdout, splitOmitEmpty(conf.Fields, ","))
	}
}

func dialAgent(resource string, addr string, opts []grpc.DialOption) (interface{}, error) {
	// Connect to gRPC server.
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Initialize new client.
	client := services.NewSystemServiceClient(conn)

	// Create context for gRPC calls.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch resource {
	case "system":
		return client.GetSystem(ctx, &services.GetSystemRequest{})
	case "users":
		return client.ListUsers(ctx, &services.ListUsersRequest{})
	case "groups":
		return client.ListGroups(ctx, &services.ListGroupsRequest{})
	case "filesystems":
		return client.ListFilesystems(ctx, &services.ListFilesystemsRequest{})
	}

	return nil, fmt.Errorf("unknown resource: %s", resource)
}

func dialAgentTable(resource string, addr string, opts []grpc.DialOption) (*text.Table, error) {
	// Connect to gRPC server.
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Initialize new client.
	client := services.NewSystemServiceClient(conn)

	// Create context for gRPC calls.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch resource {
	case "system":
		resp, err := client.GetSystem(ctx, &services.GetSystemRequest{})
		if err != nil {
			return nil, err
		}
		return system.ToTable(resp), nil
	case "users":
		resp, err := client.ListUsers(ctx, &services.ListUsersRequest{})
		if err != nil {
			return nil, err
		}
		return user.ToTable(resp.Hostname, resp.Users), nil
	case "groups":
		resp, err := client.ListGroups(ctx, &services.ListGroupsRequest{})
		if err != nil {
			return nil, err
		}
		return group.ToTable(resp.Hostname, resp.Groups), nil
	case "filesystems":
		resp, err := client.ListFilesystems(ctx, &services.ListFilesystemsRequest{})
		if err != nil {
			return nil, err
		}
		return filesystem.ToTable(resp.Hostname, resp.Filesystems), nil
	}

	return nil, fmt.Errorf("unknown resource: %s", resource)
}
