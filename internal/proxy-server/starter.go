package proxyserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	neturl "net/url"

	"github.com/Archiker-715/cache-server/internal/cache"
)

var servers = make(map[string]struct{})

func Start(port, urlToDirect string, c *cache.Cache) {
	if !alreadyStarted(port) {
		go func() {
			fmt.Printf("Server starting on :%s\n", port)

			proxy := createProxy(port, urlToDirect, c)

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				urlPath := r.URL.Path + "?" + r.URL.RawQuery
				reqPort, reqURL := c.ReflectReqPort(port), c.ReflectReqURL(urlPath)
				if c.Cached(reqPort, reqURL) {
					w.Write(c.GetCache(reqPort, reqURL))
					w.Header().Add("X-Cache", "HIT")
					return
				}

				proxy.ServeHTTP(w, r)
			})
			if err := http.ListenAndServe(":"+port, handler); err != nil {
				log.Fatalf("Server on port :%s error: %v", port, err)
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

func createProxy(port, url string, c *cache.Cache) *httputil.ReverseProxy {
	urlToDirect, err := neturl.Parse(url)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = urlToDirect.Scheme
			r.URL.Host = urlToDirect.Host
			r.Host = urlToDirect.Host
		},
		ModifyResponse: func(r *http.Response) error {
			urlPath := r.Request.URL.Path + "?" + r.Request.URL.RawQuery
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				return err
			}
			r.Body.Close()

			c.SaveCache(c.ReflectReqPort(port), c.ReflectReqURL(urlPath), bodyBytes)
			r.Header.Add("X-Cache", "MISS")
			return nil
		},
	}
	return &proxy
}
