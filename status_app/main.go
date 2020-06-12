package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/status/{status:[0-9]+}", status).Methods("GET", "HEAD")

	http.Handle("/", rtr)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func status(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	status, _ := strconv.Atoi(params["status"])
	w.WriteHeader(status)
	w.Write([]byte(params["status"]))
}
