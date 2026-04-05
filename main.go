package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func getIP(w http.ResponseWriter, r *http.Request) {
	var ip string

	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ip = strings.TrimSpace(strings.SplitN(xff, ",", 2)[0])
	} else {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err == nil {
			ip = host
		}
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%s\n", ip)
}

func main() {
	http.HandleFunc("/", getIP)
	http.ListenAndServe(":8080", nil)
}
