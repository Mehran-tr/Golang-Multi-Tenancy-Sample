package controllers

import (
	"encoding/json"
	"multi-tenant-go-app/config"
	"multi-tenant-go-app/middleware"
	"multi-tenant-go-app/utils"
	"net/http"
	"sync"
)

// User represents the structure of a user
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Fetch additional data for each user (simulate with a function)
func fetchAdditionalUserData(userID int) string {
	// Simulate fetching additional data
	return "Additional data for user"
}

// GetUsers handles the API request to fetch users for a tenant
func GetUsers(w http.ResponseWriter, r *http.Request) {
	tenant := middleware.GetTenantFromContext(r.Context())
	if tenant == nil {
		http.Error(w, "No tenant found", http.StatusInternalServerError)
		return
	}

	// Fetch users for the current tenant
	query := "SELECT id, name, email FROM users WHERE tenant_id=$1"
	rows, err := config.DB.Query(query, tenant.ID)
	if err != nil {
		http.Error(w, "Failed to query users", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			http.Error(w, "Failed to scan user", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Concurrently fetch additional data for each user
	var wg sync.WaitGroup
	additionalDataCh := make(chan string, len(users))
	for _, user := range users {
		wg.Add(1)
		go func(user User) {
			defer wg.Done()
			additionalDataCh <- fetchAdditionalUserData(user.ID)
		}(user)
	}

	// Wait for all goroutines to finish and close the channel
	go func() {
		wg.Wait()
		close(additionalDataCh)
	}()

	// Collect the additional data
	additionalDataResults := []string{}
	for data := range additionalDataCh {
		additionalDataResults = append(additionalDataResults, data)
	}

	// Combine users and additional data into the response
	response := struct {
		Users          []User   `json:"users"`
		AdditionalData []string `json:"additional_data"`
	}{
		Users:          users,
		AdditionalData: additionalDataResults,
	}

	// Send the JSON response
	json.NewEncoder(w).Encode(response)

	// Queue background jobs to send welcome emails
	for _, user := range users {
		utils.JobQueue <- utils.Job{UserID: user.ID, UserEmail: user.Email}
	}
}
