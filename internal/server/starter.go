package server

import (
	"fmt"
	"log"
	"net/http"
)

func Start(port string) {
	go func() {
		fmt.Printf("Server starting on :%s\n", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatalf("Server 1 error: %v", err)
		}
	}()
}
