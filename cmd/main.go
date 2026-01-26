package main

import (
	"flag"
	"os"

	"github.com/Archiker-715/cache-server/internal/server"
)

// описать флаги --port и --origin
// создать динамическое создание серверов через горутину
// проверять стартовал ли уже такой сервер
// описать основную логику переадресации и кеширования (через мапу ключ - url, значение - response)
func main() {

	os.Args = []string{
		"main.exe",
		"--port", "8080",
		"--url", "test",
	}

	var port, url string
	flag.StringVar(&port, "port", "8080", "starts server on port")
	flag.StringVar(&url, "url", "8080", "url which will request")
	flag.Parse()

	server.Start(port)

	select {}

	// var startedServers = make([]string, 0)
	// startedServers = append(startedServers, port)
}
