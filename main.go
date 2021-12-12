package main

import (
	"errors"
	"fmt"
	"localmods/asciibuilder"
	"log"
	"net/http"
)

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
		handleGet(w)
	case "POST":
		handlePost(w, r)
	default:
		handleDefault(w)
	}
}

func handleGet(w http.ResponseWriter) {
	text := r.FormValue("text")

	text, err := asciibuilder.Text(text)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, text)

}

func handlePost(w http.ResponseWriter, r *http.Request) {
	/*
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = writeFile("data.txt", todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(todo)

	*/
}

func handleDefault(w http.ResponseWriter) {
	err := errors.New("Sorry, only GET and POST methods are supported.")
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
