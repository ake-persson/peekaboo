package query

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/peekaboo-labs/peekaboo/pkg/filesystem"
	"github.com/peekaboo-labs/peekaboo/pkg/group"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
	"github.com/peekaboo-labs/peekaboo/pkg/system"
	"github.com/peekaboo-labs/peekaboo/pkg/text"
	"github.com/peekaboo-labs/peekaboo/pkg/user"
)

type Options struct {
	NoTLS    bool
	NoVerify bool
	NoMTLS   bool
	CertFile string
	KeyFile  string
	CAFile   string
}

type envelope struct {
	Address  string      `json:"address"`
	Response interface{} `json:"response,omitempty"`
	Error    error       `json:"error,omitempty"`
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
		format = flags.String("fmt", "json", "Output format [json,csv,table,vtable]")
		/*	colors = flags.String("colors", "light-cyan,light-yellow,cyan,yellow",
			"Comma separated list of output colors [black,red,green,yellow,blue,magenta,cyan,light-gray,\n"+
				"dark-gray,light-red,light-green,light-yellow,light-blue,light-magenta,light-cyan,white]\n\n"+
				"hostname header,hostname content,headers,content")
		noColor = flags.Bool("no-color", false, "No color output")*/
		fields = flags.String("fields", "", "Comma separated list of fields to output")
	)
	if err := flags.Parse(args); err != nil {
		return err
	}

	if len(flags.Args()) < 2 {
		usage(flags, stdout)()
		return fmt.Errorf("")
	}

	// Positional arguments.
	resource := flags.Args()[0]
	addresses := flags.Args()[1:]

	// Replace tilde with home directory.
	opts.CertFile, _ = homedir.Expand(opts.CertFile)
	opts.KeyFile, _ = homedir.Expand(opts.KeyFile)
	opts.CAFile, _ = homedir.Expand(opts.CAFile)

	// Check resource.
	if !text.InList(resource, []string{"system", "users", "groups", "filesystems"}) {
		return fmt.Errorf("unknown resource: %s", resource)
	}

	// Check format.
	if !text.InList(*format, []string{"json", "csv", "table", "vtable"}) {
		return fmt.Errorf("unknown format: %s", *format)
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
