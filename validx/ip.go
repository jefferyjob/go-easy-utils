package validx

import "net"

// IsIPv4 是否为ipv4地址
func IsIPv4(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() != nil
}

// IsIPv6 是否为ipv6地址
func IsIPv6(input string) bool {
	ip := net.ParseIP(input)
	return ip != nil && ip.To4() == nil
}
