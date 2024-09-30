package middleware

import (
	"context"
	"multi-tenant-go-app/config"
	"multi-tenant-go-app/models"
	"net/http"
)

type contextKey string

const TenantKey = contextKey("tenant")

// TenantMiddleware checks the tenant based on the request's domain.
func TenantMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		domain := r.Host // Extract domain from the host header
		tenant, err := models.GetTenantByDomain(config.DB, domain)
		if err != nil {
			http.Error(w, "Tenant not found", http.StatusNotFound)
			return
		}

		// Add tenant to the context and pass to next handler
		ctx := context.WithValue(r.Context(), TenantKey, tenant)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetTenantFromContext retrieves the tenant from the request context.
func GetTenantFromContext(ctx context.Context) *models.Tenant { // Changed models.TTenant to models.Tenant
	tenant, ok := ctx.Value(TenantKey).(*models.Tenant)
	if !ok {
		return nil
	}
	return tenant
}
