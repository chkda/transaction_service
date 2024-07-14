# transaction_service

## Project Structure

The project is organized into several directories, each with a specific purpose:

- **`/cmd`**: Contains the entry point to the service with a `main.go` file.
- **`/internal`**: Core application logic.
    - `/app`        # Application business rules (core logic).
    - `/db`      # Interactions with database.
    - `/interfaces` # Interfaces (HTTP handlers) for consumption.
- **`/pkg`**: Shared libraries and utilities used across services.
- **`/config`**: Configuration files and environment specifics.
- **`tables.sql`**: To build  database and  table for mysql.