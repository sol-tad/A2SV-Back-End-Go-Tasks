# Library Management System (Console-Based)

## Objective

To build a simple Go-based library system that can:
- Add, remove, borrow, return books
- Track members and borrowed books
- Use interfaces, structs, maps, and slices

## Structure

- `Book`, `Member` models
- `LibraryManager` interface
- `Library` struct for logic
- Console interface via `controllers`
- Organized project folders

## How to Run

```sh
go mod tidy
go run main.go
