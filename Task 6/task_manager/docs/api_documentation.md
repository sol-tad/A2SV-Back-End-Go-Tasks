# 📝 Task Management REST API (Go + Gin)

A RESTful API for managing tasks, built with Go and the Gin web framework. This project allows users to create, view, update, and delete tasks — perfect for building a simple to-do list or project tracker.

---



## 🚀 POST Man Api Documentation Link 

[text](https://documenter.getpostman.com/view/33523844/2sB34ikzaK)



## 🚀 Features

- Create, Read, Update, and Delete (CRUD) tasks
- RESTful API architecture
- JSON-based request and response
- SQLite database (easy to switch to PostgreSQL/MySQL)
- Organized folder structure (MVC-like)
- Built with Go modules

---
## 📌 API Endpoints

| Method | Endpoint     | Description            |
|--------|--------------|------------------------|
| GET    | `/tasks`     | Get all tasks          |
| GET    | `/tasks/:id` | Get task by ID         |
| POST   | `/tasks`     | Create a new task      |
| PUT    | `/tasks/:id` | Update an existing task|
| DELETE | `/tasks/:id` | Delete a task          |

