# Rachunki

Aplikacja do zarządzania rachunkami: backend w Go (`backend/`) i baza PostgreSQL 18 uruchamiana w Dockerze.

## Uruchomienie

```bash
cp .env.example .env          # potem ustaw własne hasło w .env
docker compose up -d          # startuje PostgreSQL na localhost:5432
```

Sprawdzenie połączenia z bazą:

```bash
docker compose exec postgres sh -c 'psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -c "SELECT version();"'
```

## Backend

Serwer HTTP nasłuchuje na porcie ze zmiennej `PORT` (domyślnie `8080`):

```bash
cd backend
go run ./cmd/api              # albo: PORT=3000 go run ./cmd/api
curl localhost:8080/health    # {"status":"ok"}
```

Backend nie wczytuje sam pliku `.env`, więc żeby użyć wartości `PORT` z `.env`, wyeksportuj ją przed startem, np. `set -a; source ../.env; set +a`.
