package controllers

import (
	"encoding/json"
	"multi-tenant-go-app/config"
	"multi-tenant-go-app/middleware"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// GetUsers returns all users for the current tenant.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	tenant := middleware.GetTenantFromContext(r.Context())
	if tenant == nil {
		http.Error(w, "No tenant found", http.StatusInternalServerError)
		return
	}

	rows, err := config.DB.Query("SELECT id, name, email FROM users WHERE tenant_id=$1", tenant.ID)
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, "Error scanning user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}
