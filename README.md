[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/pekaboo-io/peekaboo/master/LICENSE)

<p align="center">
  <img src="img/peekaboo.png" width="50%">
</p>

# Peekaboo

Micro-service for exposing system and hardware information.

This is a re-factoring of the original [Peekaboo](https://github.com/imc-trading/peekaboo) to support gRPC and MTLS.

## Usage

```
Usage: ./client [options] <resource> <address...>
  -ca-file string
    	CA certificate file, required for Mutual TLS (default "~/certs/root_ca.crt")
  -cert-file string
    	Server TLS certificate file (default "~/certs/srv.crt")
  -colors string
    	Comma separated list of output colors [black,red,green,yellow,blue,magenta,cyan,light-gray,
    	dark-gray,light-red,light-green,light-yellow,light-blue,light-magenta,light-cyan,white]
    	
    	hostname header,hostname content,headers,content (default "light-cyan,light-yellow,cyan,yellow")
  -fields string
    	Comma separated list of fields to output
  -fmt string
    	Output format [json,csv,table,vtable] (default "json")
  -key-file string
    	Server TLS key file (default "~/certs/srv.key")
  -mtls
    	Use Mutual TLS, client and server certificate needs to be signed by the same CA authority to establish trust ...TBD...
  -no-color
    	No color output
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

### Ubuntu/Debian

```
apt-get install goland
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

Build and start agent.

```
cd $GOPATH/src/github.com/peekaboo-labs/peekaboo
make deps build
./agent/agent -no-tls
```

Query agent.

```bash
./client/client -no-tls -fmt vtable system localhost
```
