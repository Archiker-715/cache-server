package proxyserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Archiker-715/cache-server/internal/entity"
)

var Servers = make(map[string]*http.Server)

func Start(request *entity.Request) {
	if !alreadyStarted(request.Port) {
		fmt.Printf("Server starting on :%s\n", request.Port)
		Servers[request.Port] = &http.Server{Addr: request.Port}
		if err := http.ListenAndServe(":"+request.Port, createHandler(request)); err != nil {
			log.Println(err)
			delete(Servers, request.Port)
		}
	} else {
		fmt.Printf("Server on :%s already started\n", request.Port)
	}
}

func alreadyStarted(port string) bool {
	if _, ok := Servers[port]; ok {
		return true
	}
	return false
}
