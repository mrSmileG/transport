package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mrsmileg/transp/problem"
)

func getSolve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var problem problem.Problem
	err := json.NewDecoder(r.Body).Decode(&problem)

	if err != nil {

	}

	json.NewEncoder(w).Encode(problem.GetSolution())
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}
	router := mux.NewRouter()
	router.HandleFunc("/transport/", getSolve).Methods("POST")
	problem.Logger.Fatal(http.ListenAndServe(":"+port, router))
}
