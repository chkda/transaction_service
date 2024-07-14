# transaction_service

## Project Structure

The project is organized into several directories, each with a specific purpose:

- **`/cmd`**: Contains the entry point to the service with a `main.go` file.
- **`/internal`**: Core application logic.
    - `/app`        Application business rules (core logic).
    - `/db`      Interactions with database.
    - `/interfaces`  Interfaces (HTTP handlers) for consumption.
- **`/pkg`**: Shared libraries and utilities.
- **`/config`**: Configuration files and environment specifics.
- **`tables.sql`**: To build  database and  table for mysql.

## Getting Started

### Prerequisites

- Go 1.17 or higher
- Docker (optional - for redis and mysql)

### Setting Up Your Local Development Environment

1. Clone the repository:

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up datastores:
    ```bash
    docker volume create loco_db
    docker run --rm -d --name loco-mysql -v loco_db:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root1234 -p 3306:3306 mysql
    docker run --rm -p 6379:6379 redis
    ```
    
4. Create the necessary tables

5. Run `go run cmd/main.go` 

### cURL for the respective APIs

1. Add transaction
    ```bash
    curl --location --request PUT 'localhost:8000/transactions/16' \
    --header 'Content-Type: application/json' \
    --data '{
            "amount": 25000,
            "type": "bike"
        }'
    ```

2. Get transaction
    ```bash
    curl --location 'localhost:8000/transactions/14'
    ```

3. Get transaction id based on types
    ```bash
    curl --location 'localhost:8000/types/bike'
    ```

4. Get sum for a transaction id
    ```bash
    curl --location 'localhost:8000/sum/10'
    ```