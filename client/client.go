package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/mitchellh/go-homedir"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
	"github.com/peekaboo-labs/peekaboo/pkg/storage"
	"github.com/peekaboo-labs/peekaboo/pkg/user"
)

type config struct {
	NoTLS    bool
	MTLS     bool // TBD
	CertFile string
	KeyFile  string
	CAFile   string
	Format   string
}

type envelope struct {
	Address  string      `json:"address"`
	Response interface{} `json:"response,omitempty"`
	Error    error       `json:"error,omitempty"`
}

var version = "undefined"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] <resource> <address...>\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `  resource
    	Resource to query [system, users, groups, filesystems]
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
	flag.StringVar(&conf.Format, "format", "json", "Output format [json,csv,table]")
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

	for i, addr := range addresses {
		if !strings.Contains(addr, ":") {
			addr += ":17711"
		}

		if err := dialAgent(i, resource, conf.Format, addr, opts); err != nil {
			log.Print(err)
		}
	}
}

func dialAgent(index int, resource string, format string, addr string, opts []grpc.DialOption) error {
	// Connect to gRPC server.
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Initialize new client.
	client := services.NewSystemServiceClient(conn)

	// Create context for gRPC calls.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch resource {
	case "system":
		//		e.Response, e.Error = client.GetSystem(ctx, &services.GetSystemRequest{})
	case "users":
		v, err := client.ListUsers(ctx, &services.ListUsersRequest{})
		if err != nil {
			return err
		}

		switch format {
		case "json":
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		case "csv":
			w := csv.NewWriter(os.Stdout)
			h, t := user.UsersToStringTable(v.Hostname, v.Users)
			if index == 0 {
				if err := w.Write(h); err != nil {
					return err
				}
			}
			if err := w.WriteAll(t); err != nil {
				return err
			}
		case "table":
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			h, t := user.UsersToStringTable(v.Hostname, v.Users)
			if index == 0 {
				fmt.Fprintln(w, strings.Join(h, "\t"))
			}
			for _, r := range t {
				fmt.Fprintln(w, strings.Join(r, "\t"))
			}
			w.Flush()
		}
	case "groups":
		v, err := client.ListGroups(ctx, &services.ListGroupsRequest{})
		if err != nil {
			return err
		}

		switch format {
		case "json":
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		case "csv":
			w := csv.NewWriter(os.Stdout)
			h, t := user.GroupsToStringTable(v.Hostname, v.Groups)
			if index == 0 {
				if err := w.Write(h); err != nil {
					return err
				}
			}
			if err := w.WriteAll(t); err != nil {
				return err
			}
		case "table":
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			h, t := user.GroupsToStringTable(v.Hostname, v.Groups)
			if index == 0 {
				fmt.Fprintln(w, strings.Join(h, "\t"))
			}
			for _, r := range t {
				fmt.Fprintln(w, strings.Join(r, "\t"))
			}
			w.Flush()
		}
	case "filesystems":
		v, err := client.ListFilesystems(ctx, &services.ListFilesystemsRequest{})
		if err != nil {
			return err
		}

		switch format {
		case "json":
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		case "csv":
			w := csv.NewWriter(os.Stdout)
			h, t := storage.FilesystemsToStringTable(v.Hostname, v.Filesystems)
			if index == 0 {
				if err := w.Write(h); err != nil {
					return err
				}
			}
			if err := w.WriteAll(t); err != nil {
				return err
			}
		case "table":
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
			h, t := storage.FilesystemsToStringTable(v.Hostname, v.Filesystems)
			if index == 0 {
				fmt.Fprintln(w, strings.Join(h, "\t"))
			}
			for _, r := range t {
				fmt.Fprintln(w, strings.Join(r, "\t"))
			}
			w.Flush()
		}
	}

	return nil
}
