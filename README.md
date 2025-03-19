# Go Practice - ToDo API

A simple ToDo API built with Go to learn Go web development.

📝 **演習問題**: 各演習ファイルの説明を参照してください

## Features
- RESTful API with CRUD operations for ToDo items
- Data validation
- Simple in-memory database (for learning purposes)
- Request handling with standard library and Gorilla Mux

## Requirements
- Go 1.16+
- Gorilla Mux (will be installed via go modules)

## Setup
```bash
# Navigate to project directory
cd go-practice

# Initialize Go modules
go mod init go-practice

# Install dependencies
go get -u github.com/gorilla/mux

# Run development server
go run app/main.go
```

## Testing
Once the server is running, you can:

1. Test the API with curl or any API client
2. Run the tests:
   ```bash
   go test ./tests
   ```

## API Endpoints
- `GET /todos` - List all ToDo items
- `GET /todos/{id}` - Get a single ToDo item
- `POST /todos` - Create a new ToDo item
- `PUT /todos/{id}` - Update a ToDo item
- `DELETE /todos/{id}` - Delete a ToDo item