# ipnets
[![GoDoc](https://pkg.go.dev/badge/github.com/dghubble/ipnets.svg)](https://pkg.go.dev/github.com/dghubble/ipnets)
[![Workflow](https://github.com/dghubble/ipnets/actions/workflows/test.yaml/badge.svg)](https://github.com/dghubble/ipnets/actions/workflows/test.yaml?query=branch%3Amain)
[![Sponsors](https://img.shields.io/github/sponsors/dghubble?logo=github)](https://github.com/sponsors/dghubble)
[![Mastodon](https://img.shields.io/badge/follow-news-6364ff?logo=mastodon)](https://fosstodon.org/@typhoon)

<img align="right" src="https://storage.googleapis.com/dghubble/gopher-ipnets.png">

Package `ipnets` divides `net.IPNet` networks into subnets.

Originally forked from [mhuisi/ipv4utils](https://github.com/mhuisi/ipv4utils), but now repurposed.

## Install

```
go get github.com/dghubble/ipnets
```

## Docs

Read [GoDoc](https://godoc.org/github.com/dghubble/ipnets)

## Usage

Divide a [net.IPNet](https://golang.org/pkg/net/#IPNet) by shifting the subnet bitmask with `SubnetShift`.

```
import (
    "net"
    "github.com/dghubble/ipnets"
)
...

ip, network, err := net.ParseCIDR("10.2.0.0/16")
subnets, err := ipnets.SubnetShift(network, 1)

fmt.Println(subnets)     // [10.2.0.0/17 10.2.128.0/17]
```

Divide a [net.IPNet](https://golang.org/pkg/net/#IPNet) into a desired number of subnets (or more) with `SubnetInto`. If a number is given that is not a power of 2, the returned number of subnets will be the next power of 2.

```
ip, network, err := net.ParseCIDR("10.2.0.0/16")
subnets, err := ipnets.SubnetInto(network, 4)

fmt.Println(subnets)     // [10.2.0.0/18 10.2.64.0/18 10.2.128.0/18 10.2.192.0/18]
```

## Contributing

Bug fixes are welcome. Please note that we are not accepting feature requests at this time.
