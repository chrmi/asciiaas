package main

import (
	"errors"
	"fmt"
	"localmods/asciibuilder"
	"log"
	"net/http"
)

const PATH = "./data/data.json"

func init() {
	log.Println("Main init")
}

func main() {
	http.HandleFunc("/text", handler)

	port := ":8099"
	fmt.Println("Server is running on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		handleGet(w, r)
	case "POST":
		handlePost(w, r)
	case "PUT":
		handlePut(w, r)
	default:
		handleDefault(w)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue("key")

	value, err := asciibuilder.Text(value, PATH)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, value)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	value := r.FormValue("value")
	err := asciibuilder.Add(key, value, PATH)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, value)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "POST not yet implemented", http.StatusInternalServerError)
}

func handleDefault(w http.ResponseWriter) {
	err := errors.New("Sorry, only GET and PUT methods are supported.")
	http.Error(w, err.Error(), http.StatusBadRequest)
}
