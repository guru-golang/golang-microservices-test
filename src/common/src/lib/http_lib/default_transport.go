package http_lib

import (
	"car-rent-platform/backend/common/src/lib/net_lib"
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

var resolver = New(time.Second * 5)

func DT() *http.Transport {
	return &http.Transport{
		MaxIdleConnsPerHost: 64,
		DialContext: func(context context.Context, network string, address string) (net.Conn, error) {
			separator := strings.LastIndex(address, ":")
			ipAddr, _ := resolver.FetchOneString(address[:separator])
			if net_lib.IsIPv6(ipAddr) {
				ipAddr = fmt.Sprintf(`[%s]`, ipAddr)
			}
			return net.Dial("tcp", ipAddr+address[separator:])
		},
		DialTLSContext: func(context context.Context, network string, address string) (net.Conn, error) {
			separator := strings.LastIndex(address, ":")
			ip, _ := resolver.FetchOneString(address[:separator])
			return tls.Dial("tcp", ip+address[separator:], &tls.Config{})
		},
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}
