package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/zimash/hrenovina/api/health"
)

var (
	httpPort int
)

func main() {
	flag.IntVar(&httpPort, "port", 8080, "Port to serve")
	flag.Parse()

	http.HandleFunc("/healthz", health.HandlerHealth)
	log.Printf("Server starting on %d port", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil); err != nil {
		log.Fatal(err)
	}
}
