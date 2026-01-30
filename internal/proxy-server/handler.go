package proxyserver

import (
	"net/http"

	"github.com/Archiker-715/cache-server/internal/cache"
)

func createHandler(port, urlToDirect string, c *cache.Cache) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path + "?" + r.URL.RawQuery
		reqPort, reqURL := c.ReflectReqPort(port), c.ReflectReqURL(urlPath)
		if c.Cached(reqPort, reqURL) {
			w.Write(c.GetCache(reqPort, reqURL))
			w.Header().Add("X-Cache", "HIT")
			return
		}
		proxy := createProxy(port, urlToDirect, c)
		proxy.ServeHTTP(w, r)
	})
}
