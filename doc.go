/*
Package ipnets provides net.IPNet division functions.

Divide a net.IPNet by shifting the subnet bitmask with `SubnetShift`.

	import (
		"net"
		"github.com/dghubble/ipnets"
	)
	...

	ip, network, err := net.ParseCIDR("10.2.0.0/16")
	subnets, err := ipnets.SubnetShift(network, 1)

	fmt.Println(subnets)     // [10.2.0.0/17 10.2.128.0/17]

Divide a net.IPNet into a desired number of subnets (or more) with `SubnetInto`. If a number is given that is not a power of 2, the returned number of subnets will be the next power of 2.

	ip, network, err := net.ParseCIDR("10.2.0.0/16")
	subnets, err := ipnets.SubnetInto(network, 4)

	fmt.Println(subnets)     // [10.2.0.0/18 10.2.64.0/18 10.2.128.0/18 10.2.192.0/18]
*/
package ipnets
