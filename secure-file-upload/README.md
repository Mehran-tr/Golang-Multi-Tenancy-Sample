# Secure File Upload and Download Service in Go (PostgreSQL + AES-256 Encryption)

This project demonstrates how to build a secure file upload and download service in Go. It uses **PostgreSQL** to securely store metadata for uploaded files and handles **AES-256** file encryption, validation, and secure file retrieval. It reads the encryption key from an `.env` file for enhanced security.

## Key Features

- **Secure File Upload**: Files are validated, encrypted using AES-256, and securely stored.
- **File Download**: Encrypted files are decrypted and securely streamed to the user.
- **SQL Injection Prevention**: Uses prepared statements to prevent SQL injection.
- **Rate Limiting**: Limits the number of requests per minute per user to prevent abuse.
- **AES-256 Encryption**: Files are encrypted with a strong 256-bit key, which is stored securely in a `.env` file.

## Prerequisites

- **PostgreSQL**: Make sure you have a PostgreSQL database set up and running.
- **Go**: Ensure Go is installed.

## Environment Setup

1. Set up a PostgreSQL database and create a `.env` file **one level above** the project directory with the following content:

   ```plaintext
   DB__HOST=localhost
   DB__PORT=5432
   DB__USER=yourusername
   DB__PASSWORD=yourpassword
   DB__DB=secure_file_service
   ENCRYPTION_KEY=your_base64_encoded_32_byte_key
    ```

2. To generate a secure random 32-byte encryption key for AES-256, use the following command:
    ```
    openssl rand -base64 32
    ``` 

Example output:

    ```
      GzMBE8HKkHV6vNMEAXyVD0Ks8ls+9CNGoXoVrx/Y2Io=
    ```
Add this to the .env file as the ENCRYPTION_KEY.

## Run the application

```
    go run main.go
```

## Test the endpoints

 Upload a file

```
    curl -X POST -F "file=@example.pdf" http://localhost:8080/upload
```

Download a file

```
   curl "http://localhost:8080/download?id=1"
```

