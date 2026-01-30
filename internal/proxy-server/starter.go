package proxyserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Archiker-715/cache-server/internal/cache"
)

var servers = make(map[string]struct{})

func Start(port, urlToDirect string, c *cache.Cache) {
	if !alreadyStarted(port) {
		fmt.Printf("Server starting on :%s\n", port)
		log.Println(http.ListenAndServe(":"+port, createHandler(port, urlToDirect, c)))
	} else {
		fmt.Printf("Server on :%s already started\n", port)
	}
}

func alreadyStarted(port string) bool {
	if _, ok := servers[port]; ok {
		return true
	}
	return false
}
