package proxyserver

import (
	"fmt"
	"net/http"

	"github.com/Archiker-715/cache-server/internal/entity"
)

func createHandler(request *entity.Request) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path + "?" + r.URL.RawQuery
		fmt.Println("from handler", urlPath)
		reqPort, reqURL := request.Cache.ReflectReqPort(request.Port), request.Cache.ReflectReqURL(request.Url)
		if request.Cache.Cached(reqPort, reqURL, request.Body, request.Method) {
			w.Write(request.Cache.GetCache(reqPort, reqURL, request.Body, request.Method))
			w.Header().Add("X-Cache", "HIT")
			return
		}
		proxy := createProxy(request)
		proxy.ServeHTTP(w, r)
	})
}
