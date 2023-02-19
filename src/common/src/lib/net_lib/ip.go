package net_lib

import (
	"net"
)

func IsIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && ip.To4() != nil
}

func IsIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && ip.To4() == nil
}
