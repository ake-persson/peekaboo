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

	"github.com/mickep76/color"
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

var version = "undefined"

type envelope struct {
	Address  string      `json:"address"`
	Response interface{} `json:"response,omitempty"`
	Error    error       `json:"error,omitempty"`
}

func usage(flags *flag.FlagSet, stdout io.Writer) func() {
	return func() {
		fmt.Fprintf(stdout, `Usage: %s [OPTIONS] COMMAND

Micro-service for exposing system and hardware information

Options:
`, os.Args[0])
		flags.PrintDefaults()
		fmt.Fprintf(stdout, `
Commands:
  serve     Serve API.
  query     Query API from one or more servers.
`)
	}
}

func main() {
	// Setup config and flags.
	conf := &config{}
	var printVersion bool
	flag.Usage = usage
	flag.BoolVar(&conf.NoTLS, "no-tls", false, "No TLS (testing)")
	flag.BoolVar(&conf.MTLS, "mtls", false, "Use Mutual TLS, client and server certificate needs to be signed by the same CA authority to establish trust ...TBD...") // TBD
	flag.StringVar(&conf.CertFile, "cert-file", "~/certs/srv.crt", "Server TLS certificate file")
	flag.StringVar(&conf.KeyFile, "key-file", "~/certs/srv.key", "Server TLS key file")
	flag.StringVar(&conf.CAFile, "ca-file", "~/certs/root_ca.crt", "CA certificate file, required for Mutual TLS")
	flag.StringVar(&conf.Format, "fmt", "json", "Output format [json,csv,table,vtable]")
	flag.StringVar(&conf.FormatColors, "colors", "light-cyan,light-yellow,cyan,yellow",
		"Comma separated list of output colors [black,red,green,yellow,blue,magenta,cyan,light-gray,\n"+
			"dark-gray,light-red,light-green,light-yellow,light-blue,light-magenta,light-cyan,white]\n\n"+
			"hostname header,hostname content,headers,content")
	flag.BoolVar(&conf.NoColor, "no-color", false, "No color output")
	flag.StringVar(&conf.Fields, "fields", "", "Comma separated list of fields to output")
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

	// Check colors.
	fmtColors := []string{}
	for _, c := range text.Split(conf.FormatColors, ",") {
		if v, ok := colorNames[c]; ok {
			fmtColors = append(fmtColors, v.String())
		} else {
			log.Fatalf("invalid color: %s", c)
		}
	}

	if len(fmtColors) != 4 {
		log.Fatalf("you need to specify 4 colors")
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

	rows := [][]string{}
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
			r, err := dialAgentTable(resource, addr, opts)
			if err != nil {
				log.Print(err)
			}
			rows = append(rows, r...)
		default:
			log.Fatalf("unknown output format: %s", conf.Format)
		}
	}

	switch conf.Format {
	case "json":
		b, _ := json.MarshalIndent(responses, "", "  ")
		fmt.Print(string(b))
	case "csv":
		text.PrintCSV(os.Stdout, text.Split(conf.Fields, ","), rows)
	case "table":
		text.PrintTable(os.Stdout, text.Split(conf.Fields, ","), conf.NoColor, fmtColors, rows)
	case "vtable":
		text.PrintVertTable(os.Stdout, text.Split(conf.Fields, ","), conf.NoColor, fmtColors, rows)
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
		// Customize fields for table
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

func dialAgentTable(resource string, addr string, opts []grpc.DialOption) ([][]string, error) {
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
		return [][]string{
			append([]string{"addr"}, system.Headers...),
			append([]string{addr}, system.StringSlice(resp)...),
		}, nil
	case "users":
		resp, err := client.ListUsers(ctx, &services.ListUsersRequest{})
		if err != nil {
			return nil, err
		}
		rows := [][]string{
			append([]string{"addr"}, user.Headers...),
		}
		for _, u := range resp.Users {
			r := []string{addr}
			rows = append(rows, append(r, user.StringSlice(u)...))
		}
		return rows, nil
	case "groups":
		resp, err := client.ListGroups(ctx, &services.ListGroupsRequest{})
		if err != nil {
			return nil, err
		}
		rows := [][]string{
			append([]string{"addr"}, group.Headers...),
		}
		for _, g := range resp.Groups {
			r := []string{addr}
			rows = append(rows, append(r, group.StringSlice(g)...))
		}
		return rows, nil
	case "filesystems":
		resp, err := client.ListFilesystems(ctx, &services.ListFilesystemsRequest{})
		if err != nil {
			return nil, err
		}
		rows := [][]string{
			append([]string{"addr"}, filesystem.Headers...),
		}
		for _, f := range resp.Filesystems {
			r := []string{addr}
			rows = append(rows, append(r, filesystem.StringSlice(f)...))
		}
		return rows, nil
	}

	return nil, fmt.Errorf("unknown resource: %s", resource)
}
