package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"path/filepath"

	"github.com/joho/godotenv"
)

var encryptionKey []byte

// InitCrypto reads the encryption key from the environment and decodes it
func InitCrypto() error {
	// Load .env file one level up
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	encodedKey := os.Getenv("ENCRYPTION_KEY")
	if encodedKey == "" {
		return fmt.Errorf("ENCRYPTION_KEY not set")
	}

	// Decode the base64 encoded key
	encryptionKey, err = base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		return fmt.Errorf("failed to decode encryption key: %w", err)
	}

	// Check that the key length is 32 bytes (AES-256)
	if len(encryptionKey) != 32 {
		return fmt.Errorf("invalid encryption key length: expected 32 bytes, got %d bytes", len(encryptionKey))
	}

	return nil
}

// EncryptFile encrypts the uploaded file and writes to the server storage
func EncryptFile(in io.Reader, out *os.File) error {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create an initialization vector (IV)
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return fmt.Errorf("failed to generate IV: %w", err)
	}

	// Write the IV to the output file first (needed for decryption later)
	if _, err := out.Write(iv); err != nil {
		return fmt.Errorf("failed to write IV: %w", err)
	}

	// Create the CFB encrypter stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Wrap the writer with the encryption stream writer
	writer := &cipher.StreamWriter{S: stream, W: out}

	// Copy the data from the input file to the output file through the encryption stream
	if _, err := io.Copy(writer, in); err != nil {
		return fmt.Errorf("failed to encrypt and write data: %w", err)
	}

	return nil
}

// DecryptFile decrypts a file and streams it to the response writer
func DecryptFile(in io.Reader, out io.Writer) error {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	// Read the initialization vector (IV) from the encrypted file
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(in, iv); err != nil {
		return fmt.Errorf("failed to read IV: %w", err)
	}

	// Create the CFB decrypter stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Wrap the reader with the decryption stream reader
	reader := &cipher.StreamReader{S: stream, R: in}

	// Copy the decrypted data from the encrypted file to the output writer
	if _, err := io.Copy(out, reader); err != nil {
		return fmt.Errorf("failed to decrypt and write data: %w", err)
	}

	return nil
}
