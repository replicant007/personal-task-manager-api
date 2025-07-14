Personal Task Manager API
Build a simple REST API service for managing tasks.
Core Features:

CRUD operations for tasks (Create, Read, Update, Delete)
JSON API endpoints: GET /tasks, POST /tasks, PUT /tasks/{id}, DELETE /tasks/{id}
In-memory storage initially (slice of structs), then optionally upgrade to SQLite
Basic task fields: ID, title, description, completed status, created date

What You'll Learn:

HTTP servers with net/http and routing
JSON marshaling/unmarshaling with struct tags
Error handling patterns in Go
Package organization and project structure
Testing with the standard library
Middleware concepts (logging, CORS)

Progressive Complexity - Start Simple, Add Features:

Basic CRUD with hardcoded data
Proper error responses and status codes
Request validation
Add filtering (GET /tasks?completed=true)
Database integration
Authentication with JWT tokens