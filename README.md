# Go Web Server
A basic Go web server for me to practice on.

## Links
- http://localhost - API
- http://localhost:8081/?pgsql=db&username=root&db=db&ns=public - Adminer

## Commands

Update init sql file
```
docker compose exec -it db pg_dump -s > ./database/init.sql
```
