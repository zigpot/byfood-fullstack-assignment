# Book API Backend

A Go + PostgreSQL backend for managing books.  
Designed with migrations, seeding, and clean architecture for maintainability.  
Part of a full-stack project for a software engineering job application.

---

## Project Structure

```

.
├── cmd
│   ├── main.go                 # App entrypoint
│   └── seed
│       └── main.go              # Seed runner
├── config
│   └── config.go                # Env config loader
├── db
│   ├── migrations               # SQL migrations
│   │   ├── 000001\_create\_books\_table.down.sql
│   │   └── 000001\_create\_books\_table.up.sql
│   ├── postgres.go              # DB connection
│   └── seed
│       └── seed.sql             # Idempotent seeds
├── handler
│   └── book\_handler.go          # HTTP handlers
├── models
│   └── book.go                   # Domain models
├── repository
│   └── go\_repository.go         # DB queries
├── router
│   └── router.go                 # Route definitions
├── go.mod
├── go.sum
└── README.md

````

---

## Requirements

- Go 1.21+
- PostgreSQL 15+
- `golang-migrate` or `goose` for migrations
- `make` (optional, for convenience commands)

---

## Setup

### 1. Clone the repo
```bash
git clone https://github.com/yourusername/book-api.git
cd book-api
````

### 2. Set up environment variables

Create `.env` in the root directory:

```env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=yourpassword
POSTGRES_DB=bookdb
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
```

### 3. Create the database

```bash
createdb -U postgres bookdb
```

### 4. Run migrations

```bash
make migrate-up
```

### 5. Seed initial data

```bash
make seed
```

### 6. Run the application

```bash
make run
```

---

## API Endpoints

| Method | Endpoint     | Description       |
| ------ | ------------ | ----------------- |
| GET    | `/books`     | List all books    |
| GET    | `/books/:id` | Get book by ID    |
| POST   | `/books`     | Create a new book |
| PUT    | `/books/:id` | Update a book     |
| DELETE | `/books/:id` | Delete a book     |

---

## Development Notes

* All migrations are stored in `db/migrations/`.
* Seed logic is idempotent — running it multiple times won’t duplicate data.
* Use the Makefile for consistent workflows:

```bash
make migrate-up   # Apply migrations
make migrate-down # Rollback migrations
make seed         # Run seed runner
make run          # Start the app
```

