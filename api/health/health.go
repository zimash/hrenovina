package health

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zimash/hrenovina/api/types"
)

func HandlerHealth(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s from %s\n", r.RequestURI, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	js, err := json.Marshal(types.HealthResponse{})
	if err != nil {
		log.Printf("An error %s occurred during execution\n", err)
	}
	if _, err := w.Write(js); err != nil {
		log.Printf("An error %s occurred during execution\n", err)
	}
}
