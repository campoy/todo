package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	req := struct{ Name string }{}
	dec := json.NewDecoder(os.Stdin)
	err := dec.Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
