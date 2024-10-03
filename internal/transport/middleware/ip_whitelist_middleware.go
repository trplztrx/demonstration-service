package middleware

import (
	"net"
	"net/http"
)
func ipWhitelistMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Unable to parse IP", http.StatusForbidden)
			return
		}

		if ip != "127.0.0.1" && ip != "::1" {
			http.Error(w, "Forbidden IP not allowed", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}