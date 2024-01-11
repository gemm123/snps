# SNPS

## Prerequisites

This application use:

- Golang
- PostgreSQL

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/gemm123/snps.git
   ```
   
2. Change to the project directory 

    ```bash
   cd golang-project
    ```

3. Install dependencies
    
    ```bash
   go mod tidy
   ```
   or
   ```bash
   make install
   ```
   
## Using Docker Compose

1. Run docker compose

   ```bash
   docker-compose up -d
   ```

2. Delete docker compose
   
   ```bash
   docker-compose down
   ```

3. Happy run program

## Configuration

1. Copy the example env file:

    ```bash
   cp .env.example .env
    ```

2. Open `.env` file and configure environment variables

## Database Setup

1. Migrate file .sql inside database folder

## Run the Application

1. Run the application
   
   ```bash
   go run cmd/app/app.go
   ```
   or
   ```bash
   make run
   ```
