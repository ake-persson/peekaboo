% PEEKABOO(1)
% Åke Persson
% Jun. 2020

# NAME

peekaboo - service for exposing system and hardware information.

# SYNOPSIS

**peekaboo** query [**-h**] [**-version**] [**-ca-file** *FILE*] [**-cert-file** *FILE*] [**-key-file** *FILE*] [**-no-tls**] [**-no-verify**] [**-mtls**]  <**resource**> <**address...**>
**peekaboo** agent [**-h**] [**-version**] [**-bind** *ADDR*] [**-ca-file** *FILE*] [**-cert-file** *FILE*] [**-key-file** *FILE*] [**-no-tls**] [**-no-verify**] [**-mtls**]

# DESCRIPTION

Service for exposing system and hardware information.

# OPTIONS

**-h**
:	Display a help message.

**-version**
:	Display version.

**-ca-file**
:	CA certificate file, required for Mutual TLS.

**-cert-file**
:	Server TLS certificate.

**-key-file**
:	Server TLS key file.

**-no-tls**
:	No TLS (testing).

**-no-verify**
:       No verify TLS certificate, usefull for self-signed certificates (testing).

**-mtls**
:	Use Mutual TLS, client and server certificate needs to be signed by the same CA authority to establish trust.

**-colors**
:	Comma separated list of output colors [black,red,green,yellow,blue,magenta,cyan,light-gray,
:	dark-gray,light-red,light-green,light-yellow,light-blue,light-magenta,light-cyan,white]
:
:	hostname header,hostname content,headers,content (default "light-cyan,light-yellow,cyan,yellow")

**-fields**
:	Comma separated list of fields to output

**-fmt**
:	Output format [json,csv,table,vtable] (default "json")

**-no-color**
:	No color output

# COMMANDS

**server**
:	Run peekaboo as an server on a machine.

**client**
:	Run peekaboo as a client and query one or multiple servers.

# ARGUMENTS

**resource**
:	Resource to query [system,users,groups,filesystems]

**address**
:	Address to agent specified as <address[:port]> (default port 17711)
