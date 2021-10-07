package health

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zimash/hrenovina/api"
)

func HandlerHealth(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s from %s\n", r.RequestURI, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	if _, err := fmt.Fprint(w, api.JSONPayload{}); err != nil {
		log.Printf("An error %s occurred during execution\n", err)
	}
}
