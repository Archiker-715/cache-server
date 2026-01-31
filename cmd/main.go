package main

import (
	"bufio"
	"os"
	"slices"
	"strings"

	"github.com/Archiker-715/cache-server/internal/cache"
	"github.com/Archiker-715/cache-server/internal/entity"
	"github.com/Archiker-715/cache-server/internal/flags"
	proxyserver "github.com/Archiker-715/cache-server/internal/proxy-server"
)

var cch = cache.InitCache()

func main() {
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
			go proxyserver.Start(fillRequest(port, method, url, body, cch))
		} else if clearCacheCommand(args) {
			var clearCache string
			flags.InitClearCache(args, &clearCache)
			cch.ClearCache()
		}
	}
}

func startingCommand(args []string) bool {
	if slices.Contains(args, "--port") || slices.Contains(args, "--method") || slices.Contains(args, "--origin") || slices.Contains(args, "--body") {
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

func fillRequest(port, method, url, body string, cch *cache.Cache) *entity.Request {
	return &entity.Request{
		Port:   port,
		Method: method,
		Url:    url,
		Body:   body,
		Cache:  cch,
	}
}
