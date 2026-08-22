# Go Web Server

A basic Go web server for me to practice on.

## Commands

Build & run dev

```bash
docker compose build
docker compose up
```

Build & run prod
```bash
docker build --target prod . -t go-app:prod
docker run --rm -p 80:8080 --env-file .env.prod go-app:prod
```

Bash into dev DB using container
```bash
docker compose exec -it db sh -c 'psql -U "$POSTGRES_USER" -d "$POSTGRES_DB"'
```

Update init sql file (Old)
```bash
docker compose exec -it db pg_dump -s > ./database/init.sql
``
