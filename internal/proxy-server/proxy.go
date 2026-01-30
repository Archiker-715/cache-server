package proxyserver

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	neturl "net/url"

	"github.com/Archiker-715/cache-server/internal/cache"
)

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
