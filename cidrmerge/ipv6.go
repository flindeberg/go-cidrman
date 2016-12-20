package cidrmerge

import (
	"fmt"
	"math/big"
	"net"
)

// ipv6ToUInt128 converts an IPv6 address to an unsigned 128-bit integer.
func ipv6ToUInt128(ip net.IP) *big.Int {
	return big.NewInt(0).SetBytes(ip)
}

// uint128ToIPV6 converts an unsigned 128-bit integer to an IPv6 address.
func uint128ToIPV6(addr *big.Int) net.IP {
	return net.IP(addr.Bytes()).To16()
}

// netmask6 returns the netmask for the specified prefix.
func netmask6(prefix uint) *big.Int {
	if prefix == 0 {
		return big.NewInt(0)
	}

	n := big.NewInt(1)
	n.Lsh(n, 128-prefix)
	n.Sub(n, big.NewInt(1))
	return n.Not(n) // TODO Goes negative
}

// broadcast6 returns the broadcast address for the given address and prefix.
func broadcast6(addr *big.Int, prefix uint) *big.Int {
	n := netmask6(prefix)
	n.Not(n) // TODO Goes negative
	return addr.Or(addr, n)
}

// network6 returns the network address for the given address and prefix.
func network6(addr *big.Int, prefix uint) *big.Int {
	n := netmask6(prefix)
	return addr.And(addr, n)
}

// splitRange6 recursively computes the CIDR blocks to cover the range lo to hi.
func splitRange6(addr *big.Int, prefix uint, lo, hi *big.Int, cidrs *[]*net.IPNet) error {
	if prefix > 128 {
		return fmt.Errorf("Invalid mask size: %d", prefix)
	}

	bc := broadcast6(addr, prefix)
	if (lo.Cmp(addr) < 0) || (hi.Cmp(bc) > 0) {
		return fmt.Errorf("%v, %v out of range for network %v/%d, broadcast %v", lo, hi, addr, prefix, bc)
	}

	if (lo.Cmp(addr) == 0) && (hi.Cmp(bc) == 0) {
		cidr := net.IPNet{IP: uint128ToIPV6(addr), Mask: net.CIDRMask(int(prefix), 8*net.IPv6len)}
		*cidrs = append(*cidrs, &cidr)
		return nil
	}

	prefix++
	lowerHalf := addr
	upperHalf := big.NewInt(0).SetBit(addr, int(prefix), 1)
	if hi.Cmp(upperHalf) < 0 {
		return splitRange6(lowerHalf, prefix, lo, hi, cidrs)
	} else if lo.Cmp(upperHalf) >= 0 {
		return splitRange6(upperHalf, prefix, lo, hi, cidrs)
	} else {
		err := splitRange6(lowerHalf, prefix, lo, broadcast6(lowerHalf, prefix), cidrs)
		if err != nil {
			return err
		}
		return splitRange6(upperHalf, prefix, upperHalf, hi, cidrs)
	}
}