package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	problem "github.com/mrsmileg/transport"
	"github.com/mrsmileg/transport/logger"
)

func postSolve(w http.ResponseWriter, r *http.Request) {
	var problem problem.Problem
	err := json.NewDecoder(r.Body).Decode(&problem)

	if err != nil {
		json.NewEncoder(w).Encode(logger.Error{
			Code:    214,
			Message: "Not valid task data. Please, check your input",
		})
	} else {
		json.NewEncoder(w).Encode(problem.GetSolution())
	}
}

func getSolve(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to the Transportation Problem Solver. Please give me JSON task to my body with POST method")
}

func main() {
	logger := logger.NewLogger()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8008"
	}
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET"})
	router := mux.NewRouter()
	router.HandleFunc("/", getSolve).Methods("GET")
	router.HandleFunc("/", postSolve).Methods("POST")
	logger.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
