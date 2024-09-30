package models

import (
	"database/sql"
	"log"
)

type Tenant struct {
	ID     int
	Name   string
	Domain string
}

// GetTenantByDomain fetches tenant details using domain.
func GetTenantByDomain(db *sql.DB, domain string) (*Tenant, error) {
	var tenant Tenant
	query := "SELECT id, name, domain FROM tenants WHERE domain=$1"
	err := db.QueryRow(query, domain).Scan(&tenant.ID, &tenant.Name, &tenant.Domain)
	if err != nil {
		log.Printf("Error fetching tenant: %v", err)
		return nil, err
	}
	return &tenant, nil
}
