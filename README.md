[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/pekaboo-io/peekaboo/master/LICENSE)

<p align="center">
  <img src="img/peekaboo.png" width="50%">
</p>

# Peekaboo

Expose system info using gRPC

## Usage

```
Usage: ./client/client [options] <resource> <address...>
  -ca-file string
        CA certificate file, required for Mutual TLS (default "~/certs/root_ca.crt")
  -cert-file string
        Server TLS certificate file (default "~/certs/srv.crt")
  -fields string
        Comma separate list of fields to output
  -fmt string
        Output format [json,csv,table,vtable] (default "json")
  -key-file string
        Server TLS key file (default "~/certs/srv.key")
  -mtls
        Use MTLS
  -no-tls
        No TLS (testing)
  -version
        Version
  resource
        Resource to query [system,users,groups,filesystems]
  address
        Address to agent specified as <address[:port]> (default port 17711)
```

## Install Go

First install Go and then configure Go environment.

### Mac OS X

```bash
brew install go
```

### RedHat/CentOS/Fedora

```bash
yum install golang
```

### Setup Go environment

```bash
mkdir -p ~/go/{src,bin}
cat << EOF >>~/.bash_profile
export GOPATH=~/go
export PATH=\$PATH:\$GOPATH/bin
EOF
source ~/.bash_profile
```

## Clone code

```bash
mkdir -p $GOPATH/src/github.com/peekaboo-labs
cd $GOPATH/src/github.com/peekaboo-labs
git clone https://github.com/peekaboo-labs/peekaboo.git
```

## Build and run

```
make deps build
./agent/agent -no-tls
```

Query agent.

```bash
./client/client -no-tls -fmt vtable system localhost
```
