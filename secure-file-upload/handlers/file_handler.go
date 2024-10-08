package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"secure-file-upload/database"
	"secure-file-upload/models"
	"secure-file-upload/utils"
)

// UploadFile handles secure file upload
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Limit the size of the uploaded file to 10MB
	r.ParseMultipartForm(10 << 20) // 10MB

	// Retrieve the file from the form
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate the file (e.g., check extension)
	if !utils.ValidateFileExtension(fileHeader.Filename) {
		http.Error(w, "Invalid file type", http.StatusUnsupportedMediaType)
		return
	}

	// Generate a random file name for secure storage
	randomName := make([]byte, 32)
	if _, err := rand.Read(randomName); err != nil {
		http.Error(w, "Failed to generate secure file name", http.StatusInternalServerError)
		return
	}
	encryptedFileName := base64.URLEncoding.EncodeToString(randomName)

	// Create the encrypted file path
	encryptedFilePath := filepath.Join("storage/encrypted_files", encryptedFileName)

	// Create the file on the server
	outFile, err := os.Create(encryptedFilePath)
	if err != nil {
		http.Error(w, "Failed to store the file", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Encrypt and save the file
	err = utils.EncryptFile(file, outFile)
	if err != nil {
		log.Printf("Encryption error: %v", err) // Log the error for debugging
		http.Error(w, "Failed to encrypt file", http.StatusInternalServerError)
		return
	}

	// Store file metadata in the PostgreSQL database using a prepared statement
	stmt, err := database.DB.Prepare("INSERT INTO files(original_name, encrypted_path) VALUES($1, $2)")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(fileHeader.Filename, encryptedFilePath)
	if err != nil {
		http.Error(w, "Error storing file metadata", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("File uploaded and encrypted successfully"))
}

// DownloadFile handles secure file download
func DownloadFile(w http.ResponseWriter, r *http.Request) {
	// Get file ID from query params
	fileID := r.URL.Query().Get("id")
	if fileID == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	// Retrieve file metadata from the PostgreSQL database using a prepared statement
	var file models.File
	err := database.DB.QueryRow("SELECT id, original_name, encrypted_path FROM files WHERE id = $1", fileID).Scan(&file.ID, &file.OriginalName, &file.EncryptedPath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Open the encrypted file
	encryptedFile, err := os.Open(file.EncryptedPath)
	if err != nil {
		http.Error(w, "Error opening the file", http.StatusInternalServerError)
		return
	}
	defer encryptedFile.Close()

	// Decrypt the file and send it to the client
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", file.OriginalName))
	w.Header().Set("Content-Type", "application/octet-stream")

	err = utils.DecryptFile(encryptedFile, w)
	if err != nil {
		http.Error(w, "Error decrypting file", http.StatusInternalServerError)
		return
	}
}
