// Package ipv4utils provides utilities
// for working with IPv4 addresses.
package ipv4utils

import (
	"fmt"
	"net"
)

// Merge merges an IP to a value.
//
// Merging an IP allows for the use of operations on the IP.
func Merge(ip []byte) (merged uint32, err error) {
	if len(ip) != 4 {
		return 0, fmt.Errorf("%s is not an IPv4 address.", ip)
	}
	merged = uint32(0)
	// Takes the most significant octet,
	// shifts it to the left end of merged,
	// takes the next lesser significant octet,
	// shifts it to the start of the last shifted octet in merged, etc.
	for i, octet := range ip {
		merged |= uint32(octet) << (8 * (3 - uint32(i)))
	}
	return merged, nil
}

// Split splits a value into an IP.
//
// After merging an IP to perform operations on it, Split
// can be used to put the IP back into a slice.
func Split(merged uint32) []byte {
	ip := make([]byte, 4)
	for i := 3; i >= 0; i-- {
		// Takes the current least significant octet of merged,
		// inserts it into ip, disposes the least significant
		// octet of merged and moves the next more significant octet
		// into the position of the least significant octet.
		ip[i] = byte(merged & 0xFF)
		merged >>= 8
	}
	return ip
}

// Subnet subnets a network with the provided amount of bits.
//
// The network IP may have host bits set, as long as
// this address is a valid network address within the new subnet mask.
//
// This means that you can provide an offset for subnetting through the network IP:
//
// Network IP: 192.168.1.128/24, bits: 1 -> First subnet is 192.168.1.128/25, not 192.168.1.0/25
//
// The channel to provide the subnets is closed once subnetting is completed.
func Subnet(network net.IPNet, bits uint) (subnets chan net.IPNet, err error) {
	unchecked := network.IP
	originalMask := network.Mask
	if bits > 31 {
		return nil, fmt.Errorf("%d exceeds the maximum amount of subnettable bits (31).", bits)
	}
	maskLen, _ := originalMask.Size()
	newMaskLen := uint(maskLen) + bits
	if newMaskLen > 31 {
		return nil, fmt.Errorf("/%d subnetted with %d bits to /%d exceeds the size of valid IPv4 subnets (31).",
			maskLen, bits, newMaskLen)
	}
	uncheckedVal, err := Merge(unchecked)
	if err != nil {
		return nil, err
	}
	ip := uint64(uncheckedVal)
	hostBits := 32 - newMaskLen
	hosts := uint64(1 << hostBits)
	if ip&(hosts-1) != 0 {
		return nil, fmt.Errorf("%s has host bits set in /%d.", unchecked, newMaskLen)
	}

	subnets = make(chan net.IPNet)
	go func() {
		maskVal, _ := Merge(originalMask)
		baseIP := ip & uint64(maskVal)
		totalSubnets := uint64(1 << bits)
		totalHosts := totalSubnets * hosts
		nextIP := baseIP + totalHosts
		lastIP := nextIP - hosts
		newMask := net.CIDRMask(int(newMaskLen), 32)
		for ip <= lastIP {
			subnets <- net.IPNet{
				IP:   Split(uint32(ip)),
				Mask: newMask,
			}
			ip += hosts
		}
		close(subnets)
	}()
	return subnets, nil
}
