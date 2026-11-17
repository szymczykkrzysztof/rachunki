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
