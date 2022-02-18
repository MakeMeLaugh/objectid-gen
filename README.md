# objectid-gen

Simple cli tool for working with ObjectID 

# Features

* Generate ObjectID from current timestamp;
* Convert existing timestamp (in RFC3339 format) to ObjectID form (padded with zeros to the right);
* Generate ObjectID from relative point in time in the past (e.g. 1 hour 21 minutes and 42 seconds ago);
* Extract timestamp from existing ObjectID string;

# Installation

## Precompiled binaries

Precompiled binary releases for Linux (amd64), MacOS (arm64 and amd64) and Windows (amd64) are available under [Releases](https://gitlab.com/MakeMeLaugh/objectid-gen/-/releases) page of this repo.

---

## Install using Golang toolchain

Installing `objectid-gen` using Golang toolchain:

```shell
go install gitlab.com/MakeMeLaugh/objectid-gen@<version>
```

Available versions match [releases](https://gitlab.com/MakeMeLaugh/objectid-gen/-/releases) names. Be aware that `@latest` version relates to the HEAD of the `main` branch and might be in "work in progress" state. Consider using only named tagged versions.

You need to make sure that `$GOBIN` environment variable (usually it's `$GOPATH/bin` or `$HOME/go/bin`) is added to your `$PATH` environment variable.

## Building manually using Golang

If you want to build it from source code yourself you are free to "go get" this repo and build it manually:

```shell
$ go get gitlab.com/MakeMeLaugh/objectid-gen
$ cd objectid-gen && go build -o /your/desired/path/to/executables
```

The `main` package provides some compile-time variables which can be set during compilation (there are no defaults for these variables) with `-ldflags '-X main.<variable>=<value>'`:

```go
var (
    // Used as an application name while generating help message
    applicationName string
    // Application version (used in -V flag)
    applicationVersion string
    // Time when application was built at (used in -V flag)
    buildAt string
    // Short commit hash that was used for building the application (used in -V flag)
    buildFrom string
)
```

## Building from source

Download source code from [Releases](https://gitlab.com/MakeMeLaugh/objectid-gen/-/releases) page for a specific version or just clone this repo and run `make`:

```shell
git clone git@gitlab.com:MakeMeLaugh/objectid-gen.git
# or git clone https://gitlab.com/MakeMeLaugh/objectid-gen.git
cd objectid-gen && make
```

By default, `make` will build binaries for every supported platform (`darwin/amd64`, `darwin/arm64`, `linux/amd64`,  and `windows/amd64`) which will be placed under `bin` directory inside the repo folder. If you want to override this behavior - run make with `BUILD_PLATFORMS` set to your platform:

```shell
make BUILD_PLATFORMS=darwin/amd64
```

Note that `make` recreates content of the `bin` directory (removes everything from it) on each call.
This behavior cannot be overridden.

# Usage

```shell
Usage of objectid-gen:
  -V    Show version number and exit
  -a string
        String representation of time in the past to generate ObjectID from (valid time units are "s", "m", "h")
  -h    Show  this message and exit
  -o string
        ObjectID to parse and return as datetime (in RFC3339 format)
  -t string
        Datetime to generate ObjectID from (in RFC3339 format)
```

# Examples

### Generate ObjectID from current timestamp

```shell
$ objectid-gen
6112e51d9357733c89798a85
```

### Extract datetime (RFC3339 string) from existing ObjectID

```shell
$ objectid-gen -o 6112e51d9357733c89798a85
2021-08-10T23:44:13+03:00
```

### Convert datetime (RFC3339 string) to ObjectID string

```shell
$ objectid-gen -t '2021-08-10T23:44:13+03:00'
6112e51d0000000000000000
```

* **Be aware**: timestamp string will be parsed as is: `2021-08-10T23:44:13+03:00` and `2021-08-10T23:44:13Z` will generate different ObjectID strings;
* due to ObjectID spec only first 8 characters represent time, so the string will be padded with zeros to the right;
   
### Generate ObjectID from a duration string

```shell
$ objectid-gen -a '13h29m42s'
611227830000000000000000
```

* "ns", "us" (or "µs") and "ms" time units are accepted but ignored due to ObjectID spec;
* due to ObjectID spec only first 8 characters represent time, so the string will be padded with zeros to the right;
