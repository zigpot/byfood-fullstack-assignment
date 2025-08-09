# Backend TODOs

- [ ] Use environment variables for all DB credentials and config values.  
      - Store them in `.env`.  
      - Load via `os.Getenv` in `config/config.go`.  
      - Keys: `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, `POSTGRES_HOST`, `POSTGRES_PORT`.

- [ ] Keep `cmd/main.go` minimal — delegate all init logic to packages.  
      - Example: `func main() { app.Run() }`.

- [ ] Use a migration tool (`golang-migrate` or `goose`).  
      - Migrate with: `make migrate-up`.  
      - Add seeds as a separate step (`make seed`).

- [ ] Make seed runner idempotent (`INSERT ... WHERE NOT EXISTS`).  
      - Ensure it exits gracefully if migrations haven’t been run.

- [ ] Add a `Makefile` for common commands:  
      - `migrate-up`: run migrations.  
      - `seed`: run seed runner.  
      - `run`: start the application.

- [ ] Add at least one integration test hitting the DB.  
      - Place `_test.go` files in relevant packages or create `tests/` dir.

- [ ] Add "Local Development with Docker" section to `README.md` for zero-setup review.