package utils

import (
	"net"
	"net/http"
)

func GetIPAddress(r *http.Request) string {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)

	return host
}
