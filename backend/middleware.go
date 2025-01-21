package main

import (
	"context" // Import the context package
	"net"
	"net/http"
	"strings"
)

func IPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if strings.Contains(ip, ":") {
			ip, _, _ = net.SplitHostPort(ip)
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx,"ip", ip)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
