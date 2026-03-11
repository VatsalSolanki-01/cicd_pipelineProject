package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go CI/CD Pipeline Application")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {

	response := Message{
		Service: "go-cicd-app",
		Status:  "running",
		Time:    time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api", apiHandler)

	fmt.Println("Server running on port 8081")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}