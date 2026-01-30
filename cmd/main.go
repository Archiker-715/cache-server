package main

import (
	"os"
	"slices"

	"github.com/Archiker-715/cache-server/internal/cache"
	"github.com/Archiker-715/cache-server/internal/flags"
	proxyserver "github.com/Archiker-715/cache-server/internal/proxy-server"
)

var cch = cache.InitCache()

func main() {

	// for dbg
	os.Args = []string{
		"main.exe",
		"--port", "8080",
		"--origin", "test",
	}

	if startingCommand() {
		var port, method, url, body string
		flags.InitStartingServer(&port, &method, &url, &body)
		proxyserver.Start(port, url, cch)
	} else if clearCacheCommand() {
		var clearCache string
		flags.InitClearCache(&clearCache)
		cch.ClearCache()
	}

	select {}
}

func startingCommand() bool {
	if slices.Contains(os.Args, "--port") || slices.Contains(os.Args, "--method") || slices.Contains(os.Args, "--url") || slices.Contains(os.Args, "--body") {
		return true
	}
	return false
}

func clearCacheCommand() bool {
	if slices.Contains(os.Args, "--clear-cache") {
		return true
	}
	return false
}
