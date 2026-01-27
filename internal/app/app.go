package app

import (
	"os"
	"slices"

	"github.com/Archiker-715/cache-server/internal/cache"
	"github.com/Archiker-715/cache-server/internal/flags"
	"github.com/Archiker-715/cache-server/internal/server"
)

func Run() {

	if slices.Contains(os.Args, "--port") || slices.Contains(os.Args, "--method") || slices.Contains(os.Args, "--url") || slices.Contains(os.Args, "--body") {
		var port, method, url, body string
		flags.InitStartingServer(&port, &method, &url, &body)

		server.Start(port, url)

		// respBody, err := httpSender.Send(port, url, method, body)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// var buf bytes.Buffer
		// if err := json.Indent(&buf, respBody, "", "  "); err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(buf.String())
	} else if slices.Contains(os.Args, "--clear-cache") {
		var clearCache string
		flags.InitClearCache(&clearCache)
		cache.ClearCache()
	}

}
