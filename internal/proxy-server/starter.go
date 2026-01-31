package proxyserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Archiker-715/cache-server/internal/entity"
)

var servers = make(map[string]struct{})

func Start(request *entity.Request) {
	if !alreadyStarted(request.Port) {
		fmt.Printf("Server starting on :%s\n", request.Port)
		log.Println(http.ListenAndServe(":"+request.Port, createHandler(request)))
	} else {
		fmt.Printf("Server on :%s already started\n", request.Port)
	}
}

func alreadyStarted(port string) bool {
	if _, ok := servers[port]; ok {
		return true
	}
	return false
}
