# hr - Hostname Resolver

Simple tool to generate a delimited list of hostname -> ip address given any number of hostnames from STDIN

## Usage

```bash
hr [-d=";"] hostname1 hostname2 hostname3
```

Example output:

```
hostname1;10.0.0.22
hostname2;10.0.0.23
hostname3;10.0.0.45
```

## Installation

```bash
go build -o /usr/local/bin/hr cmd/hr/main.go
```
