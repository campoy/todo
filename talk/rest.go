package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const address = "127.0.0.1:8080"

func main() {
	http.HandleFunc("/", addHandler)

	// Server started in a different goroutine.
	go func() {
		log.Fatal(http.ListenAndServe(address, nil))
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("%v + %v = %v\n", i, 2*i, add(i, 2*i))
	}
}

type AddRequest struct {
	A, B int
}

type AddResponse struct {
	Result int
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, r.Method+" not allowed", http.StatusMethodNotAllowed)
		return
	}

	req := AddRequest{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	res := AddResponse{req.A + req.B}
	enc := json.NewEncoder(w)
	err = enc.Encode(res)
	if err != nil {
		log.Println(err)
		http.Error(w, "oops", http.StatusInternalServerError)
	}
}

func add(a, b int) int {
	req := AddRequest{a, b}
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(req)
	if err != nil {
		log.Fatalf("encoding request: %v", err)
	}

	r, err := http.Post("http://"+address, "application/json", buf)
	if err != nil {
		log.Fatalf("sending request: %v", err)
	}
	defer r.Body.Close()

	res := AddResponse{}
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&res)
	if err != nil {
		log.Fatalf("decoding response: %v", err)
	}
	return res.Result
}
