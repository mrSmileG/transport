package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mrsmileg/transp/problem"
)

func postSolve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var problem problem.Problem
	err := json.NewDecoder(r.Body).Decode(&problem)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(problem.GetSolution())
}

func getSolve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to the Transportation Problem Solver. Please give me JSON task to my body with POST method")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET"})
	router := mux.NewRouter()
	router.HandleFunc("/", getSolve).Methods("GET")
	router.HandleFunc("/transport/", postSolve).Methods("POST")
	problem.Logger.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
