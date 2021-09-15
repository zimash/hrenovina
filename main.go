package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	httpPort int
)

type jsonPayload struct{}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s from %s\n", r.RequestURI, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	if _, err := fmt.Fprint(w, jsonPayload{}); err != nil {
		log.Printf("An error %s occured during execution\n", err)
	}
}

func main() {
	flag.IntVar(&httpPort, "port", 8080, "Port to serve")
	flag.Parse()

	http.HandleFunc("/healthz", healthCheck)
	log.Printf("Server starting on %d port", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), nil); err != nil {
		log.Fatal(err)
	}
}
