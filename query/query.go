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

	"github.com/mickep76/color"
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

var colorNames = map[string]color.Code{
	"black":         color.Black,
	"red":           color.Red,
	"green":         color.Green,
	"yellow":        color.Yellow,
	"blue":          color.Blue,
	"magenta":       color.Magenta,
	"cyan":          color.Cyan,
	"light-gray":    color.LightGray,
	"dark-gray":     color.DarkGray,
	"light-red":     color.LightRed,
	"light-green":   color.LightGreen,
	"light-yellow":  color.LightYellow,
	"light-blue":    color.LightBlue,
	"light-magenta": color.LightMagenta,
	"light-cyan":    color.LightCyan,
	"white":         color.White,
}

func usage(flags *flag.FlagSet, stdout io.Writer) func() {
	return func() {
		fmt.Fprintf(stdout, `Usage: %s query [OPTIONS]

Query API from one or multiple servers

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
		colors = flags.String("colors", "light-cyan,light-yellow,cyan,yellow",
			"Comma separated list of output colors [black,red,green,yellow,blue,magenta,cyan,light-gray,\n"+
				"dark-gray,light-red,light-green,light-yellow,light-blue,light-magenta,light-cyan,white]\n\n"+
				"hostname header,hostname content,headers,content")
		noColor = flags.Bool("no-color", false, "No color output")
		fields  = flags.String("fields", "", "Comma separated list of fields to output")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	// Check resource and address(es).
	if len(flags.Args()) < 2 {
		usage(flags, stdout)()
		return fmt.Errorf("no resource or address specified")
	}
	resource := flags.Args()[0]
	addresses := flags.Args()[1:]

	log.Println("flags", flags.Args())

	// Check resource.
	// TODO
	// - Resources should register on import.
	if !text.InList(resource, []string{"system", "users", "groups", "filesystems"}) {
		return fmt.Errorf("unknown resource: %s", resource)
	}

	// Check format.
	if !text.InList(*format, []string{"json", "csv", "table", "vtable"}) {
		return fmt.Errorf("unknown format: %s", *format)
	}

	// Check colors.
	// TODO
	// - Improve color notation.
	fmtColors := []string{}
	for _, c := range text.Split(*colors, ",") {
		if v, ok := colorNames[c]; ok {
			fmtColors = append(fmtColors, v.String())
		} else {
			log.Fatalf("invalid color: %s", c)
		}
	}

	if len(fmtColors) != 4 {
		log.Fatalf("you need to specify 4 colors")
	}

	grpcOpts := []grpc.DialOption{grpc.WithBlock()}
	if opts.NoTLS {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	} else {
		// Load CA certificate file.
		creds, err := credentials.NewClientTLSFromFile(opts.CAFile, "")
		if err != nil {
			return fmt.Errorf("failed to load ca certificate: %v", err)
		}

		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(creds))
	}

	rows := [][]string{}
	responses := []interface{}{}
	for _, addr := range addresses {
		// Set default port.
		if !strings.Contains(addr, ":") {
			addr += ":17711"
		}

		switch *format {
		case "json":
			resp, err := dialAgent(resource, addr, grpcOpts)
			if err != nil {
				log.Print(err)
			}
			responses = append(responses, resp)
		case "csv", "table", "vtable":
			r, err := dialAgentTable(resource, addr, grpcOpts)
			if err != nil {
				log.Print(err)
			}
			rows = append(rows, r...)
		default:
			return fmt.Errorf("unknown format: %s", *format)
		}
	}

	switch *format {
	case "json":
		b, _ := json.MarshalIndent(responses, "", "  ")
		fmt.Print(string(b))
	case "csv":
		text.PrintCSV(os.Stdout, text.Split(*fields, ","), rows)
	case "table":
		text.PrintTable(os.Stdout, text.Split(*fields, ","), *noColor, fmtColors, rows)
	case "vtable":
		text.PrintVertTable(os.Stdout, text.Split(*fields, ","), *noColor, fmtColors, rows)
	}

	return nil
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
