package ipnets

import (
	"bytes"
	"net"
	"testing"
)

func TestSubnet(t *testing.T) {
	cases := []struct {
		network  net.IPNet
		shift    int
		expected []net.IPNet
	}{
		// shift 1
		{
			mustIPNet(t, "172.16.0.0/16"), 1,
			[]net.IPNet{
				mustIPNet(t, "172.16.0.0/17"),
				mustIPNet(t, "172.16.128.0/17"),
			},
		},
		// shift 2
		{
			mustIPNet(t, "172.16.0.0/16"), 2,
			[]net.IPNet{
				mustIPNet(t, "172.16.0.0/18"),
				mustIPNet(t, "172.16.64.0/18"),
				mustIPNet(t, "172.16.128.0/18"),
				mustIPNet(t, "172.16.192.0/18"),
			},
		},
		// shift 3
		{
			mustIPNet(t, "172.16.0.0/16"), 3,
			[]net.IPNet{
				mustIPNet(t, "172.16.0.0/19"),
				mustIPNet(t, "172.16.32.0/19"),
				mustIPNet(t, "172.16.64.0/19"),
				mustIPNet(t, "172.16.96.0/19"),
				mustIPNet(t, "172.16.128.0/19"),
				mustIPNet(t, "172.16.160.0/19"),
				mustIPNet(t, "172.16.192.0/19"),
				mustIPNet(t, "172.16.224.0/19"),
			},
		},
		// edge cases
		{
			mustIPNet(t, "10.0.0.0/30"), 1,
			[]net.IPNet{
				mustIPNet(t, "10.0.0.0/31"),
				mustIPNet(t, "10.0.0.2/31"),
			},
		},
		{
			mustIPNet(t, "10.0.0.0/30"), 2,
			[]net.IPNet{
				mustIPNet(t, "10.0.0.0/32"),
				mustIPNet(t, "10.0.0.1/32"),
				mustIPNet(t, "10.0.0.2/32"),
				mustIPNet(t, "10.0.0.3/32"),
			},
		},
		// zeros
		{
			mustIPNet(t, "0.0.0.0/0"), 0,
			[]net.IPNet{
				mustIPNet(t, "0.0.0.0/0"),
			},
		},
		{
			mustIPNet(t, "0.0.0.0/0"), 2,
			[]net.IPNet{
				mustIPNet(t, "0.0.0.0/2"),
				mustIPNet(t, "64.0.0.0/2"),
				mustIPNet(t, "128.0.0.0/2"),
				mustIPNet(t, "192.0.0.0/2"),
			},
		},
		{
			mustIPNet(t, "10.0.0.0/16"), 0,
			[]net.IPNet{
				mustIPNet(t, "10.0.0.0/16"),
			},
		},
		{
			mustIPNet(t, "192.168.1.0/24"), 0,
			[]net.IPNet{
				mustIPNet(t, "192.168.1.0/24"),
			},
		},
	}
	for _, c := range cases {
		subnets, err := SubnetShift(&c.network, c.shift)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if !equalSubnets(subnets, c.expected) {
			t.Errorf("expected: %s, got: %s", c.expected, subnets)
		}
	}
}

func mustIPNet(t *testing.T, s string) net.IPNet {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	return *ipnet
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

func TestSubnetInto(t *testing.T) {
	cases := []struct {
		network  net.IPNet
		count    int
		expected []net.IPNet
	}{
		{
			mustIPNet(t, "10.0.0.0/16"), 1,
			[]net.IPNet{
				mustIPNet(t, "10.0.0.0/16"),
			},
		},
		{
			mustIPNet(t, "10.0.0.0/16"), 2,
			[]net.IPNet{
				mustIPNet(t, "10.0.0.0/17"),
				mustIPNet(t, "10.0.128.0/17"),
			},
		},
		{
			mustIPNet(t, "10.0.0.0/16"), 3,
			[]net.IPNet{
				mustIPNet(t, "10.0.0.0/18"),
				mustIPNet(t, "10.0.64.0/18"),
				mustIPNet(t, "10.0.128.0/18"),
				mustIPNet(t, "10.0.192.0/18"),
			},
		},
		{
			mustIPNet(t, "10.0.0.0/16"), 4,
			[]net.IPNet{
				mustIPNet(t, "10.0.0.0/18"),
				mustIPNet(t, "10.0.64.0/18"),
				mustIPNet(t, "10.0.128.0/18"),
				mustIPNet(t, "10.0.192.0/18"),
			},
		},
		{
			mustIPNet(t, "10.0.0.0/16"), 5,
			[]net.IPNet{
				mustIPNet(t, "10.0.0.0/19"),
				mustIPNet(t, "10.0.32.0/19"),
				mustIPNet(t, "10.0.64.0/19"),
				mustIPNet(t, "10.0.96.0/19"),
				mustIPNet(t, "10.0.128.0/19"),
				mustIPNet(t, "10.0.160.0/19"),
				mustIPNet(t, "10.0.192.0/19"),
				mustIPNet(t, "10.0.224.0/19"),
			},
		},
		{
			mustIPNet(t, "10.0.0.0/16"), 12,
			[]net.IPNet{
				mustIPNet(t, "10.0.0.0/20"),
				mustIPNet(t, "10.0.16.0/20"),
				mustIPNet(t, "10.0.32.0/20"),
				mustIPNet(t, "10.0.48.0/20"),
				mustIPNet(t, "10.0.64.0/20"),
				mustIPNet(t, "10.0.80.0/20"),
				mustIPNet(t, "10.0.96.0/20"),
				mustIPNet(t, "10.0.112.0/20"),
				mustIPNet(t, "10.0.128.0/20"),
				mustIPNet(t, "10.0.144.0/20"),
				mustIPNet(t, "10.0.160.0/20"),
				mustIPNet(t, "10.0.176.0/20"),
				mustIPNet(t, "10.0.192.0/20"),
				mustIPNet(t, "10.0.208.0/20"),
				mustIPNet(t, "10.0.224.0/20"),
				mustIPNet(t, "10.0.240.0/20"),
			},
		},
	}
	for _, c := range cases {
		subnets, err := SubnetInto(&c.network, c.count)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
		if !equalSubnets(subnets, c.expected) {
			t.Errorf("expected: %s, got: %s", c.expected, subnets)
		}
	}
}

func TestNumeric(t *testing.T) {
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
		num := numeric(c.in)
		if num != c.expected {
			t.Errorf("expected: %d, got: %d", c.expected, num)
		}
	}
}

func TestBytewise(t *testing.T) {
	cases := []struct {
		in       uint32
		expected net.IP
	}{
		{0, net.IP([]byte{0, 0, 0, 0})},
		{2130706433, net.IP([]byte{127, 0, 0, 1})},
		{2886729728, net.IP([]byte{172, 16, 0, 0})},
		{3232235776, net.IP([]byte{192, 168, 1, 0})},
		{4294967295, net.IP([]byte{255, 255, 255, 255})},
	}
	for _, c := range cases {
		ip := bytewise(c.in)
		if !ip.Equal(c.expected) {
			t.Errorf("expected: %v, got: %v", c.expected, ip)
		}
	}
}
