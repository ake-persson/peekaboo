package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type config struct {
	NoTLS    bool
	NoVerify bool
	NoMTLS   bool
	CertFile string
	KeyFile  string
	CAFile   string
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: %s [OPTIONS] COMMAND

Micro-service for exposing system and hardware information

Options:
`, os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, `
Commands:
  serve     Serve API.
  query     Query API from one or more servers.
`)
}

var version = "undefined"

func main() {
	// Setup config and flags.
	conf := &config{}
	var printVersion bool
	flag.BoolVar(&conf.NoTLS, "no-tls", false, "No TLS (testing)")
	flag.BoolVar(&conf.NoVerify, "no-verify", false, "No TLS verify, for self-signed certificates (testing)")
	flag.BoolVar(&conf.NoMTLS, "no-mtls", true, "No Mutual TLS, client and server certificate needs to be signed by the same CA authority to establish trust")
	flag.StringVar(&conf.CertFile, "cert-file", "~/certs/srv.crt", "Server TLS certificate file")
	flag.StringVar(&conf.KeyFile, "key-file", "~/certs/srv.key", "Server TLS key file")
	flag.StringVar(&conf.CAFile, "ca-file", "~/certs/root_ca.crt", "CA certificate file, required for Mutual TLS")
	flag.BoolVar(&printVersion, "version", false, "Version")
	flag.Usage = usage
	flag.Parse()

	// Peint version.
	if printVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	// Check command.
	if len(flag.Args()) < 1 {
		usage()
		os.Exit(1)
	}

	// Run command.
	switch os.Args[1] {
	case "serve":
	case "query":
	default:
		log.Fatalf("unknown command: %s", os.Args[1])
	}
}
