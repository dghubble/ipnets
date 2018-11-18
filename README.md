# ipnets [![Build Status](https://travis-ci.org/dghubble/ipnets.svg?branch=master)](https://travis-ci.org/dghubble/ipnets) [![Coverage](https://gocover.io/_badge/github.com/dghubble/ipnets)](https://gocover.io/github.com/dghubble/ipnets) [![GoDoc](https://godoc.org/github.com/dghubble/ipnets?status.svg)](https://godoc.org/github.com/dghubble/ipnets)

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
