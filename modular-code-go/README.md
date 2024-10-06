# Structuring Go Applications with Packages

This project demonstrates how to structure a large Go application using packages and Go modules. It shows how to break down code into manageable pieces and avoid circular dependencies while keeping the code modular and maintainable.

## Key Features
- **Models**: Define your data structures (e.g., `User`).
- **Services**: Contains business logic (e.g., user operations like add and fetch).
- **Handlers**: HTTP request handlers (e.g., API endpoints for users).
- **Utils**: Helper functions (e.g., logging).

## Project Structure

 ```
modular-code-go/
├── services/   
    └──  user_service.go          
├── models/  
    └──  user.go            
├── handlers/   
    └──  user_handler.go            
├── utils/
    └──  logger.go           
├── go.mod               
└── main.go 
└── README.md              
```

## How Run

- **in cmd**:

 ```
go run main.go
 ```

- **add user with curl**:
 ```
curl -X POST -d '{"id": 1, "name": "John", "email": "john@example.com"}' http://localhost:8080/add-user
 ```

- **get users with curl**:
 ```
curl http://localhost:8080/users

 ```