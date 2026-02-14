package handlers

import "net/http"

func HealthyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy"))
}
