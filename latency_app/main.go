package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/latency/{latency:[0-9]+}", latency).Methods("GET", "HEAD")

	http.Handle("/", rtr)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func latency(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	latency, _ := strconv.Atoi(params["latency"])
	time.Sleep(time.Duration(latency) * time.Millisecond)
	w.WriteHeader(200)
	w.Write([]byte(params["latency"]))
}
