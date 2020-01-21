# k3sd

Management daemon and CLIs for k3s servers and agents.

## Overview

`k3sd` is a collection of a management daemon and CLIs for the k3s Kubernetes distribution. k3s is built of two main components; servers and agents. In a similar way, `k3sd` is built of multiple components. The components are:

- `k3sd`, a k3s management daemon with a gRPC interface
- `k3sserverctl`, a CLI for `k3sd` that manages k3s servers
- `k3sagentctl`, a CLI for `k3sd` that manages k3s agents

`k3sd` bundles the `k3s` binary into it's own binary and extracts it on startup, so there is no need to install it manually.

## Installation

A Go package [is available](https://godoc.org/github.com/pojntfx/k3sd). In order to use it, you have to `go generate` it first.

## Usage

### Daemon

The daemon requires root priviledges.

You may also set the flags by setting env variables in the format `K3SD_[FLAG]` (i.e. `K3SD_K3SD_CONFIGFILE=examples/k3sd.yaml`) or by using a [configuration file](examples/k3sd.yaml).

```bash
% k3sd --help
k3sd is the k3s management daemon.

Find more information at:
https://pojntfx.github.io/k3sd/

Usage:
  k3sd [flags]

Flags:
  -h, --help                          help for k3sd
  -f, --k3sdd.configFile string       Configuration file to use.
  -l, --k3sdd.listenHostPort string   TCP listen host:port. (default "localhost:1340")
```

### Client CLIs

There are two client CLIs, `k3sserverctl` and `k3sagentctl`.

#### `k3sserverctl`

You may also set the flags by setting env variables in the format `K3SSERVER_[FLAG]` (i.e. `K3SSERVER_K3SSERVER_CONFIGFILE=examples/k3sserver.yaml`) or by using a [configuration file](examples/k3sserver.yaml).

```bash
% k3sserverctl --help
k3sserverctl manages servers on k3sd, the k3s management daemon.

Find more information at:
https://pojntfx.github.io/k3sd/

Usage:
  k3sserverctl [command]

Available Commands:
  cleanup     Clean up a server
  get         Get a resource
  help        Help about any command
  start       Start a server
  stop        Stop a server

Flags:
  -h, --help   help for k3sserverctl

Use "k3sserverctl [command] --help" for more information about a command.
```

#### `k3sagentctl`

You may also set the flags by setting env variables in the format `K3SAGENT_[FLAG]` (i.e. `K3SAGENT_K3SAGENT_CONFIGFILE=examples/k3sagent.yaml`) or by using a [configuration file](examples/k3sagent.yaml).

```bash
% k3sagentctl --help
k3sagentctl manages agents on k3sd, the k3s management daemon.

Find more information at:
https://pojntfx.github.io/k3sd/

Usage:
  k3sagentctl [command]

Available Commands:
  cleanup     Clean up an agent
  help        Help about any command
  start       Start an agent
  stop        Stop an agent

Flags:
  -h, --help   help for k3sagentctl

Use "k3sagentctl [command] --help" for more information about a command.
```

## License

k3sd (c) 2020 Felix Pojtinger

SPDX-License-Identifier: AGPL-3.0
