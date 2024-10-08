package utils

import (
	"path/filepath"
	"strings"
)

// ValidateFileExtension ensures the uploaded file has a safe extension (e.g., .pdf, .jpg)
func ValidateFileExtension(filename string) bool {
	allowedExtensions := []string{".pdf", ".jpg", ".jpeg", ".png"}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}
	return false
}
