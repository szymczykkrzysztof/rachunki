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

Serwer HTTP nasłuchuje na porcie ze zmiennej `PORT` (domyślnie `8080`). Wymaga też `DATABASE_URL` i przy starcie sprawdza połączenie z bazą (najpierw `docker compose up -d`):

```bash
cd backend
go run ./cmd/api              # albo: PORT=3000 go run ./cmd/api
curl localhost:8080/health    # {"status":"ok"}
```

Backend sam wczytuje `.env` z bieżącego katalogu albo z katalogu nadrzędnego (więc działa i z `backend/`, i z głównego katalogu repo). Zmienne ustawione w środowisku mają pierwszeństwo przed `.env`.
