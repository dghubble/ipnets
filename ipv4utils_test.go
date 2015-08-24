package ipv4utils

import (
	"bytes"
	"net"
	"testing"
)

func BenchmarkSubnet(t *testing.B) {
	net := ipv4Net("172.16.0.0/16")
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		ch, _ := Subnet(net, 4)
		for _ = range ch {
		}
	}
}

func TestMerge(t *testing.T) {
	cases := []struct {
		in       []byte
		expected uint32
	}{
		{[]byte{0, 0, 0, 0}, 0},
		{[]byte{127, 0, 0, 1}, 2130706433},
		{[]byte{172, 16, 0, 0}, 2886729728},
		{[]byte{192, 168, 1, 0}, 3232235776},
		{[]byte{255, 255, 255, 255}, 4294967295},
	}
	for _, c := range cases {
		actual, _ := Merge(c.in)
		e := c.expected
		if actual != e {
			t.Errorf("expected: %d, got: %d", e, actual)
		}
	}
}

func TestSplit(t *testing.T) {
	cases := []struct {
		in       uint32
		expected []byte
	}{
		{0, []byte{0, 0, 0, 0}},
		{2130706433, []byte{127, 0, 0, 1}},
		{2886729728, []byte{172, 16, 0, 0}},
		{3232235776, []byte{192, 168, 1, 0}},
		{4294967295, []byte{255, 255, 255, 255}},
	}
	for _, c := range cases {
		actual := Split(c.in)
		e := c.expected
		if !bytes.Equal(actual, e) {
			t.Errorf("expected: %d, got: %d", e, actual)
		}
	}
}

func ipv4Net(cidr string) net.IPNet {
	ip, n, _ := net.ParseCIDR(cidr)
	n.IP = ip.To4()
	return *n
}

func collectSubnets(n net.IPNet, bits uint, amount int) []net.IPNet {
	s := make([]net.IPNet, amount)
	receiver, _ := Subnet(n, bits)
	for i := 0; i < amount; i++ {
		s[i] = <-receiver
	}
	return s
}

func equalSubnets(actual []net.IPNet, expected []net.IPNet) bool {
	for i, e := range expected {
		a := actual[i]
		if !e.IP.Equal(a.IP) || !bytes.Equal(e.Mask, a.Mask) {
			return false
		}
	}
	return true
}

func TestSubnet(t *testing.T) {
	cases := []struct {
		network  net.IPNet
		bits     uint
		expected []net.IPNet
	}{
		// Normal case
		{
			ipv4Net("172.16.0.0/16"), 8,
			[]net.IPNet{
				ipv4Net("172.16.0.0/24"),
				ipv4Net("172.16.1.0/24"),
				ipv4Net("172.16.2.0/24"),
			},
		},
		// Offset case
		{
			ipv4Net("172.16.0.128/24"), 3,
			[]net.IPNet{
				ipv4Net("172.16.0.128/27"),
				ipv4Net("172.16.0.160/27"),
				ipv4Net("172.16.0.192/27"),
			},
		},
		// Edge cases
		{
			ipv4Net("0.0.0.0/0"), 2,
			[]net.IPNet{
				ipv4Net("0.0.0.0/2"),
				ipv4Net("64.0.0.0/2"),
				ipv4Net("128.0.0.0/2"),
				ipv4Net("192.0.0.0/2"),
				net.IPNet{IP: nil, Mask: nil},
			},
		},
		{
			ipv4Net("255.255.255.250/0"), 31,
			[]net.IPNet{
				ipv4Net("255.255.255.250/31"),
				ipv4Net("255.255.255.252/31"),
				ipv4Net("255.255.255.254/31"),
				net.IPNet{IP: nil, Mask: nil},
			},
		},
		// 0-cases
		{
			ipv4Net("192.168.1.0/24"), 0,
			[]net.IPNet{ipv4Net("192.168.1.0/24")},
		},
		{
			ipv4Net("0.0.0.0/0"), 0,
			[]net.IPNet{ipv4Net("0.0.0.0/0")},
		},
	}
	for _, c := range cases {
		e := c.expected
		actual := collectSubnets(c.network, c.bits, len(e))
		if !equalSubnets(actual, e) {
			t.Errorf("expected: %s, got: %s", e, actual)
		}
	}
}
