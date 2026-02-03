package proxyserver

import (
	"context"
	"log"
	"time"
)

func Shutdown(port string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if server, ok := Servers[port]; ok {
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Shutdown err: %v", err)
			return
		}
		delete(Servers, port)
	}
}
