package flags

import "flag"

func InitStartingServer(port, method, url, body *string) {
	flag.StringVar(port, "port", "8080", "starts server on port")
	flag.StringVar(method, "method", "8080", "http method")
	flag.StringVar(url, "origin", "", "url which will request")
	flag.StringVar(body, "body", "{}", "request JSON-body")
	flag.Parse()
}

func InitClearCache(clearCache *string) {
	flag.StringVar(clearCache, "clear-cache", "", "command for clearing cache")
	flag.Parse()
}
