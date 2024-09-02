package handlers

import (
	services "api/internal/service"
	"api/pkg/models"
	"encoding/json"
	"net/http"
)

func HandleUserData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var userData models.DataUser
	
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	
	if err = services.SendEmail(userData); err != nil {
		http.Error(w, "Bad Request, SendEmail", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}

func HandleCors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "POST" {
            return
        }

        next.ServeHTTP(w, r)
    })
}
