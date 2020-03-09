[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://raw.githubusercontent.com/pekaboo-io/peekaboo/master/LICENSE)

<p align="center">
  <img src="img/peekaboo.png" width="50%">
</p>

# Peekaboo

Expose system info using gRPC

# Example

Start agent.

```bash
make build
./agent/agent -no-tls
```

Query agent.

```bash
./client/client -no-tls
```

