package proxyserver

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	neturl "net/url"

	"github.com/Archiker-715/cache-server/internal/entity"
)

func createProxy(request *entity.Request) *httputil.ReverseProxy {
	urlToDirect, err := neturl.Parse(request.Url)
	if err != nil {
		log.Fatal(err)
	}

	proxy := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = urlToDirect.Scheme
			r.URL.Host = urlToDirect.Host
			r.Host = urlToDirect.Host
		},
		ModifyResponse: func(resp *http.Response) error {
			urlPath := resp.Request.URL.Path + "?" + resp.Request.URL.RawQuery
			fmt.Println("proxy", urlPath)
			respBodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(respBodyBytes))

			request.Cache.SaveCache(request.Cache.ReflectReqPort(request.Port), request.Cache.ReflectReqURL(urlPath), request.Body, request.Method, respBodyBytes)
			resp.Header.Add("X-Cache", "MISS")
			return nil
		},
	}
	return &proxy
}
