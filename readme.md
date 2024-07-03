# Go Clean Arch

## Installation

1. Clone repository:

   ```bash
   git clone https://github.com/shironxn/go-clean-arch
   ```

2. Navigate to the project directory:

   ```bash
   cd go-clean-arch
   ```

3. Setup environment variables:
   ```bash
   cp .env.example .env
   ```

## Usage

1. Check all available commands:

   ```bash
   make help
   ```

2. Run the application:

   ```bash
   make run
   ```

3. Run in development mode using Air:
   ```bash
   make dev
   ```

## API Reference

### Authors

| Method | URL             | Description            |
| ------ | --------------- | ---------------------- |
| POST   | `/authors`      | Create a new author    |
| GET    | `/authors`      | Get all authors        |
| GET    | `/authors/{id}` | Get an author by ID    |
| PUT    | `/authors/{id}` | Update an author by ID |
| DELETE | `/authors/{id}` | Delete an author by ID |

### Books

| Method | URL           | Description         |
| ------ | ------------- | ------------------- |
| POST   | `/books`      | Create a new book   |
| GET    | `/books`      | Get all books       |
| GET    | `/books/{id}` | Get a book by ID    |
| PUT    | `/books/{id}` | Update a book by ID |
| DELETE | `/books/{id}` | Delete a book by ID |
