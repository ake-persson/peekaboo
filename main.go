package main

import (
	"flag"
	"fmt"
	"os"

	//	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	//	"github.com/peekaboo-labs/client"
	//	"github.com/peekaboo-labs/server"
)

var version = "undefined"

type config struct {
	Addr         string
	NoTLS        bool
	NoVerify     bool
	MTLS         bool // TBD
	CertFile     string
	KeyFile      string
	CAFile       string
	Format       string
	FormatColors string
	NoColor      bool
	Fields       string
	Level        zapcore.Level // TBD
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
	flag.StringVar(&conf.Addr, "addr", "localhost:17711", "Server address")
	flag.BoolVar(&conf.NoTLS, "no-tls", false, "No TLS (testing)")
	flag.BoolVar(&conf.NoVerify, "no-verify", false, "No verify TLS (testing)")
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
}
