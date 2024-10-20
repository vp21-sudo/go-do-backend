
# Go-Do Backend

Welcome to the **Go-Do** project! This is the backend API for a simple to-do app built using **GoLang** and **Fiber**. The purpose of this project is to explore GoLang, understand its unique features, and compare it to other languages like TypeScript.


## Features

- Create, read, update, and delete (CRUD) to-do items.
- Manage tasks with simple, efficient APIs.
- Explore GoLang’s concurrency, performance, and scalability.
- Learn the differences between GoLang and TypeScript.

## Tech Stack

- **GoLang**: The programming language used for building the backend.
- **Fiber**: A Go web framework for handling HTTP requests.
- **MongoDB**: Database for storing tasks (optional, adjust as needed).
- **Docker**: For containerization (optional).

## API Endpoints
###  Here are some of the core API endpoints:

- **GET /todos** - Fetch all to-do items.
- **GET /todos/:id** - Fetch a specific to-do item by ID.
- **POST /todos** - Create a new to-do item.
- **PUT /todos/:id** - Update an existing to-do item.
- **DELETE /todos/:id** - Delete a to-do item.

## Example
### To create a new to-do item:
```bash
curl -X POST http://localhost:9000/todos \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn GoLang", "description": "Explore the Fiber framework"}'
```

## Future Enhancements
- Add user authentication and authorization.
- Implement task prioritization.
- Add database options for persistence.
- Enhance the project to handle larger-scale operations.
