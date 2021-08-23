# objectid-gen

Simple cli tool for working with ObjectID 

# Features

* Generating ObjectID from current timestamp;
* Converting existing timestamp (in RFC3339 format) to ObjectID form (padded with zeros to the right);
* Generate ObjectID from relative point in time in the past (e.g. 1 hour 21 minutes and 42 seconds ago);
* Extracting timestamp from existing ObjectID string;

# Installation

For now, `objectid-gen` only available for installation as a Golang module:

```shell
go install gitlab.com/MakeMeLaugh/objectid-gen
```

If you want to install specific version of the tool just use `@v<version>`

You need to make sure that `$GOBIN` environment variable (usually it's `$GOPATH/bin` or `$HOME/go/bin`) is added to your `$PATH` environment variable.

---

Of course if you want to build it from source code yourself you are free to clone this repo or get it using Golang toolchain and build it manually:

```shell
$ go get gitlab.com/MakeMeLaugh/objectid-gen
$ cd objectid-gen && go build -o /your/desired/path/to/executables
```

Or download source code from [Releases](https://gitlab.com/MakeMeLaugh/objectid-gen/-/releases) page for a specific version.
Precompiled binary releases for Linux (amd64) and MacOS (arm64 and amd64) are also available under Releases page of this repo.

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
        Time to generate ObjectID from (in RFC3339 format)
```

# Examples

### Generate ObjectID from current timestamp

```shell
$ objectid-gen
6112e51d9357733c89798a85
```

### Extract timestamp from existing ObjectID

```shell
$ objectid-gen -o 6112e51d9357733c89798a85
2021-08-10T23:44:13+03:00
```

### Convert timestamp to ObjectID string

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