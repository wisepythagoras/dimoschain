package utils

import "net"

// IsIPAddressValid checks whether a provided IP address is valid.
func IsIPAddressValid(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	}

	return true
}
