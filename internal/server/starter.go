package server

import (
	"fmt"
	"log"
	"net/http"
)

var servers = make(map[string]struct{})

func Start(port string) {
	if !alreadyStarted(port) {
		go func() {
			fmt.Printf("Server starting on :%s\n", port)
			if err := http.ListenAndServe(":"+port, nil); err != nil {
				log.Fatalf("Server 1 error: %v", err)
			}
		}()
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
