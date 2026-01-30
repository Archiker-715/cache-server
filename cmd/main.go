package main

import (
	"bufio"
	"os"
	"slices"
	"strings"

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

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "" {
			continue
		}
		args := strings.Fields(line)

		if startingCommand(args) {
			var port, method, url, body string
			flags.InitStartingServer(args, &port, &method, &url, &body)
			go proxyserver.Start(port, url, cch)
		} else if clearCacheCommand(args) {
			var clearCache string
			flags.InitClearCache(args, &clearCache)
			cch.ClearCache()
		}
	}

	select {}
}

func startingCommand(args []string) bool {
	if slices.Contains(args, "--port") || slices.Contains(args, "--method") || slices.Contains(args, "--url") || slices.Contains(args, "--body") {
		return true
	}
	return false
}

func clearCacheCommand(args []string) bool {
	if slices.Contains(args, "--clear-cache") {
		return true
	}
	return false
}
