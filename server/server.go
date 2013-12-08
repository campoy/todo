package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/campoy/todo/todo"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/task", listTasks).Methods("GET")
	r.HandleFunc("/task", newTask).Methods("POST")
	r.HandleFunc("/task/{id}", getTask).Methods("GET")
	r.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	http.ListenAndServe(":8080", r)
}

var tasks = todo.NewTaskManager()

func listTasks(w http.ResponseWriter, r *http.Request) {
	var res struct{ Tasks []*todo.Task }
	res.Tasks = tasks.All()
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
		log.Println(err)
	}
}

func newTask(w http.ResponseWriter, r *http.Request) {
	req := struct{ Title string }{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t, err := todo.NewTask(req.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tasks.Save(t)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	txt := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(txt, 10, 0)
	if err != nil {
		http.Error(w, "task ID is not a number", http.StatusBadRequest)
		return
	}
	t, ok := tasks.Find(id)
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := json.NewEncoder(w).Encode(t); err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
		log.Println(err)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	txt := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(txt, 10, 0)
	if err != nil {
		http.Error(w, "task ID is not a number", http.StatusBadRequest)
		return
	}
	var t todo.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "decode failed", http.StatusBadRequest)
		return
	}
	if t.ID != id {
		http.Error(w, "inconsistent task ID", http.StatusBadRequest)
		return
	}
	if _, ok := tasks.Find(t.ID); !ok {
		http.NotFound(w, r)
		return
	}
	tasks.Save(&t)
}
