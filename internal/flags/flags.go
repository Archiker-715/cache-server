package flags

import "flag"

func InitStartingServer(args []string, port, method, url, body *string) {
	fs := flag.NewFlagSet("start", flag.ContinueOnError)
	fs.StringVar(port, "port", "8080", "starts server on port")
	fs.StringVar(method, "method", "8080", "http method")
	fs.StringVar(url, "origin", "", "url which will request")
	fs.StringVar(body, "body", "{}", "request JSON-body")
	fs.Parse(args)
}

func InitClearCache(args []string, clearCache *string) {
	fs := flag.NewFlagSet("clear-cache", flag.ContinueOnError)
	fs.StringVar(clearCache, "clear-cache", "", "command for clearing cache")
	fs.Parse(args)
}
