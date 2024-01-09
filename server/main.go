package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	userData = make(map[string]string)
	mu       sync.Mutex
)

func main() {
	r := mux.NewRouter()


	corsHandler := mux.CORSMethodMiddleware(r)
	r.Use(corsHandler)

	r.HandleFunc("/register", registerHandler).Methods("POST", "OPTIONS")

	port := 8080
	fmt.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		User string `json:"user"`
		Pwd  string `json:"pwd"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//  username and password validation
	if len(req.User) < 4 || len(req.User) > 24 || len(req.Pwd) < 8 || len(req.Pwd) > 24 {
		http.Error(w, "Invalid entry", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Check if the username is already taken
	if _, exists := userData[req.User]; exists {
		http.Error(w, "Username taken", http.StatusConflict)
		return
	}

	// Store user data 
	userData[req.User] = req.Pwd

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"message": "Registration successful"}
	json.NewEncoder(w).Encode(response)
}
