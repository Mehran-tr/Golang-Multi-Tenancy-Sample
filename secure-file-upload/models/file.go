package models

// File represents metadata of an uploaded file
type File struct {
	ID            int
	OriginalName  string // The original file name
	EncryptedPath string // The encrypted file path on the server
}
