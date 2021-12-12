package main

import (
	"fmt"
	"localmods/asciibuilder"
	"log"
	"net/http"
)

func init() {
	log.Println("Main init")
}

func main() {
	http.HandleFunc("/text", getText)

	port := ":8099"
	fmt.Println("Server is running on port" + port)

	log.Fatal(http.ListenAndServe(port, nil))
}

func getText(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")

	text, err := asciibuilder.Text(text)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, text)

}
