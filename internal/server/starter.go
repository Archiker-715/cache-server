package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	neturl "net/url"
)

var servers = make(map[string]struct{})

func Start(port, urlToDirect string) {
	if !alreadyStarted(port) {
		go func() {
			fmt.Printf("Server starting on :%s\n", port)
			if err := http.ListenAndServe(":"+port, createProxy(urlToDirect)); err != nil {
				log.Fatalf("Server 1 error: %v", err)
			}
		}()
	} else {
		fmt.Printf("Server on :%s already started\n", port)
	}
}

func alreadyStarted(port string) bool {
	if _, ok := servers[port]; ok {
		return true
	}
	return false
}

func createProxy(url string) *httputil.ReverseProxy {
	targetURL, err := neturl.Parse(url)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = targetURL.Scheme
			r.URL.Host = targetURL.Host
			r.Host = targetURL.Host
		},
	}
	return &proxy
}
