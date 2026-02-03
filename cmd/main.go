package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/Archiker-715/cache-server/internal/cache"
	ch "github.com/Archiker-715/cache-server/internal/command-handler"
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
		ch.HandleCommand(args, cch)
	}
}
