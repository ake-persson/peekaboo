package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	//	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/peekaboo-labs/peekaboo/query"
	"github.com/peekaboo-labs/peekaboo/serve"
)

var version = "undefined"

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

// Abstract main so we can test it as a regular function.
func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.SetOutput(stdout)
	flags.Usage = usage(flags, stdout)
	var (
		noTLS        = flags.Bool("no-tls", false, "No TLS (testing)")
		noMTLS       = flags.Bool("no-mtls", true, "No Mutual TLS, client and server certificate needs to be signed by the same CA authority to establish trust")
		noVerify     = flags.Bool("no-verify", false, "No TLS verify, for self-signed certificates (testing)")
		certFile     = flags.String("cert-file", "~/certs/srv.crt", "Server TLS certificate file")
		keyFile      = flag.String("key-file", "~/certs/srv.key", "Server TLS key file")
		caFile       = flag.String("ca-file", "~/certs/root_ca.crt", "CA certificate file, required for Mutual TLS")
		printVersion = flag.Bool("version", false, "Version")
	)
	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	// Print version.
	if *printVersion {
		fmt.Fprintf(stdout, "%s\n", version)
		return nil
	}

	// Replace tilde with home directory.
	*certFile, _ = homedir.Expand(*certFile)
	*keyFile, _ = homedir.Expand(*keyFile)
	*caFile, _ = homedir.Expand(*caFile)

	// Check command.
	if len(flags.Args()) < 1 {
		usage(flags, stdout)()
		return fmt.Errorf("no command specified")
	}
	command := flags.Args()[0]

	// Run command.
	args = append([]string{os.Args[0]}, flags.Args()[1:]...)
	switch command {
	case "serve":
		return serve.Run(args, stdout, &serve.Options{
			NoTLS:    *noTLS,
			NoMTLS:   *noMTLS,
			NoVerify: *noVerify,
			CertFile: *certFile,
			KeyFile:  *keyFile,
			CAFile:   *caFile})
	case "query":
		return query.Run(args, stdout, &query.Options{
			NoTLS:    *noTLS,
			NoMTLS:   *noMTLS,
			NoVerify: *noVerify,
			CertFile: *certFile,
			KeyFile:  *keyFile,
			CAFile:   *caFile})
	}

	return fmt.Errorf("unknown command: %s", command)
}
