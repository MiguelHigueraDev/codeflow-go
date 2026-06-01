# codeflow-go

Rewrite of a past live coding interview web app I made in Go to learn the language.

Not finished at all.

## Prerequisites

- [Go](https://go.dev/dl/) 1.25+
- [Docker](https://www.docker.com/) (for Postgres)

## Setup

1. Clone the repo and install dependencies:

   ```bash
   git clone https://github.com/miguelhigueradev/codeflow-go.git
   cd codeflow-go
   go mod download
   ```

2. Start Postgres:

   ```bash
   docker compose up -d
   ```

3. Configure the auth service environment:

   ```bash
   cp services/auth-service/.env.example services/auth-service/.env
   ```

   Set `DATABASE_URL` and your GitHub OAuth credentials in `services/auth-service/.env`:

   ```
   PORT=8080
   DATABASE_URL=postgres://postgres:postgres@localhost:5432/codeflow?sslmode=disable
   GITHUB_CLIENT_ID=your_client_id
   GITHUB_CLIENT_SECRET=your_client_secret
   ```

4. Run the auth service from the repo root:

   ```bash
   go run ./services/auth-service/cmd/api
   ```

   The server listens on port 8080.

## Services

| Service            | Path                              | Description              |
| ------------------ | --------------------------------- | ------------------------ |
| auth-service       | `services/auth-service/cmd/api`   | Authentication API (WIP)     |
| room-coordinator   | `services/room-coordinator/cmd/api` | Room coordination (WIP) |
