package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"

	"github.com/Archiker-715/cache-server/internal/cache"
	"github.com/Archiker-715/cache-server/internal/http"
	"github.com/Archiker-715/cache-server/internal/server"
)

func main() {

	// for dbg
	os.Args = []string{
		"main.exe",
		"--port", "8080",
		"--origin", "test",
	}

	// TODO: на стартовавшем порте нужно чтобы можно было постучаться по роуту локалхост:порт
	// и тогда в зависимости от порта шла переадресация на тот роут что был настроен через --порт --ориджин
	// в общем нужнон намутить редиректор

	// TODO: вывести в отдельный файл
	var port, method, url, body, clearCache string
	if slices.Contains(os.Args, "--port") || slices.Contains(os.Args, "--method") || slices.Contains(os.Args, "--url") || slices.Contains(os.Args, "--body") {
		flag.StringVar(&port, "port", "8080", "starts server on port")
		flag.StringVar(&method, "method", "8080", "http method")
		flag.StringVar(&url, "origin", "", "url which will request")
		flag.StringVar(&body, "body", "{}", "request JSON-body")
		flag.Parse()

		server.Start(port)
		respBody, err := http.Send(port, url, method, body)
		if err != nil {
			fmt.Println(err)
		}

		var buf bytes.Buffer
		if err := json.Indent(&buf, respBody, "", "  "); err != nil {
			fmt.Println(err)
		}
		fmt.Println(buf.String())
	} else if slices.Contains(os.Args, "--clear-cache") {
		flag.StringVar(&clearCache, "clear-cache", "", "command for clearing cache")
		flag.Parse()
		cache.ClearCache()
	}

	select {}
}
