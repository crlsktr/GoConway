package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	// "strconv"

	"github.com/gorilla/mux"
)

func startServer() {
	router := mux.NewRouter()

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/nextGen",nextGen).Methods("POST")
	server := &http.Server{
		Addr: ":4000",
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())
}

func nextGen(w http.ResponseWriter, r *http.Request) {
	body,_ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
	response,_ := json.Marshal("gotcha")
	w.Write(response)
}

func home(w http.ResponseWriter, r *http.Request) {
	response,_ := json.Marshal(size)
	w.Write(response)
}