
# Multi-Tenant Go Application

This project is a simple multi-tenant application built with Go (Golang) and PostgreSQL. It demonstrates how to handle multi-tenancy using domain-based tenant isolation, where each tenant's data is identified using a unique domain in the request.

## Features

- **Multi-Tenancy**: Isolates tenant data by the domain name provided in the HTTP request.
- **PostgreSQL Integration**: Uses PostgreSQL as the database for storing tenants and users.
- **Database Migrations**: Uses `golang-migrate` for database migrations to manage schema changes.
- **Middleware**: Injects tenant data into the request context based on the domain in the request.

## Project Structure

```
multi-tenant-go-app/
├── config/                 # Database configuration and initialization
├── controllers/            # HTTP handlers and API logic
├── middleware/             # Tenant context middleware
├── models/                 # Database models and queries
├── routes/                 # API routing
├── migrations/             # Database migration files
├── .env                    # Environment variables (DB configuration)
├── go.mod                  # Go module file (dependencies)
├── go.sum                  # Go dependency versions
└── main.go                 # Application entry point
```

## Prerequisites

- **Go (Golang)**: [Installation Instructions](https://golang.org/doc/install)
- **PostgreSQL**: [Installation Instructions](https://www.postgresql.org/download/)
- **golang-migrate**: CLI tool for database migrations.

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/multi-tenant-go-app.git
   cd multi-tenant-go-app
   ```

2. **Install Go dependencies**:

   ```bash
   go mod tidy
   ```

3. **Set up PostgreSQL**:
   - Install PostgreSQL on your system if you haven't already.
   - Create a database and user for the app:

     ```sql
     CREATE DATABASE multi_tenant_app;
     CREATE USER tenant_user WITH PASSWORD 'tenant_password';
     GRANT ALL PRIVILEGES ON DATABASE multi_tenant_app TO tenant_user;
     ```

4. **Configure environment variables**:
   - Create a `.env` file in the root of your project and configure your PostgreSQL credentials:

     ```bash
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=tenant_user
     DB_PASSWORD=tenant_password
     DB_NAME=multi_tenant_app
     DATABASE_URL=postgres://tenant_user:tenant_password@localhost:5432/multi_tenant_app?sslmode=disable
     ```

5. **Run Database Migrations**:
   - Install `golang-migrate` CLI if you haven't yet:

     ```bash
     brew install golang-migrate  # MacOS
     ```

     For Linux, follow the [instructions here](https://github.com/golang-migrate/migrate#cli-usage).

   - Run migrations to set up the database schema:

     ```bash
     migrate -path ./migrations -database $DATABASE_URL up
     ```

6. **Run the Application**:

   ```bash
   go run main.go
   ```

   If everything is configured correctly, you should see:
   ```bash
   Database connected successfully.
   Migrations applied successfully.
   ```

## Usage

### Create a Tenant

1. Log in to PostgreSQL and insert a tenant into the `tenants` table:

   ```sql
   INSERT INTO tenants (name, domain) VALUES ('Tenant1', 'tenant1.localhost:8080');
   ```

2. Insert a user associated with the tenant:

   ```sql
   INSERT INTO users (tenant_id, name, email, password_hash) 
   VALUES (1, 'John Doe', 'john.doe@example.com', 'somehashedpassword');
   ```

### Testing the API

You can test the API with tools like `curl` or Postman. Ensure you use the `Host` header to specify the tenant domain.

#### Get Users for a Tenant

Use the following `curl` command to retrieve the users for `Tenant1`:

```bash
curl -H "Host: tenant1.localhost:8080" http://localhost:8080/users
```

You should see a response like:

```json
[
    {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com"
    }
]
```

### Routes

- **GET /users**: Returns all users for the current tenant based on the request's domain.

## Troubleshooting

- **Tenant not found**: Ensure that the tenant exists in the `tenants` table, and the domain matches exactly (including port number).
- **Database connection issues**: Double-check your `.env` configuration and verify that the PostgreSQL service is running.
- **Migrations failed**: Ensure that the `migrations` folder contains both `.up.sql` and `.down.sql` files for each migration.


### Concurrency and Background Task Processing

This multi-tenant Go app efficiently handles concurrent operations and background tasks using Go’s concurrency features.

#### **Concurrent Data Fetching**
When users are requested for a tenant, additional user-related data is fetched concurrently using goroutines. This improves performance by allowing multiple operations to execute in parallel, reducing the overall response time.

#### **Background Task Processing**
After the API response is sent, background jobs (like sending emails) are delegated to a worker pool. These tasks are processed asynchronously, without blocking the main request, ensuring scalability under heavy load.

#### **Worker Pool**
The application implements a worker pool to manage background tasks efficiently. Each worker handles tasks such as sending emails for users in the background, logged in the console as they complete.




## License

This project is licensed under the MIT License.
