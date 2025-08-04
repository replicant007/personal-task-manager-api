# Personal Task Manager API
Build a simple REST API service for managing tasks.


## Roadmap

:white_check_mark: Create a slice of type Task (ID, title, description, completed status, created date)

:white_check_mark: Swap slice of Tasks with a SQLite DB

:white_check_mark: Create handlers and routes for all type of operations (CRUD)

:white_check_mark: Implement validation

:white_check_mark: Extract database interaction into an implementation of DB interface

:point_right: Add tests

Swagger or Redoc integration

Add users & Authentication with JWT tokens

Implement filtering

Add error handling (Handle errors gracefully)

Enhance logging, add CORS

Implement struct tags for validation and DB

Gracefully shut down server (handling of an error from http.ListenAndServe)

Deploy it locally and with Docker


## Core features

CRUD operations for tasks (Create, Read, Update, Delete)

JSON API endpoints: GET /tasks, POST /tasks, PUT /tasks/{id}, DELETE /tasks/{id}

In-memory storage initially (slice of structs), then optionally upgrade to SQLite

Basic task fields: ID, title, description, completed status, created date

## What will I learn

HTTP servers with net/http and routing

JSON marshaling/unmarshaling with struct tags

Error handling patterns in Go

Package organization and project structure

Testing with the standard library

Middleware concepts (logging, CORS)

## Progressive Complexity - Start Simple, Add Features

Basic CRUD with hardcoded data
Proper error responses and status codes
Request validation
Add filtering (GET /tasks?completed=true)
Database integration
Authentication with JWT tokens

## Additional features

Write a .gitlab-ci.yml file to automate builds and tests for the project

How to manage shared variables (slice of tasks / DB) betweeen multiple goroutines?

Implement pagination or filtering logic in the service layer
