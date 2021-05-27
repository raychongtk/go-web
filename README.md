# Go Web

a demo for building web application in Go

## Tech Stack

1. Dependency Injection - Wire
2. HTTP Web Framework - Gin
3. Session Management - Gin-Session, Redis
4. Database ORM - Gorm
5. Database - Postgres
6. Load Configuration - Viper

# Spin up a Postgres instance with Docker

go to `.postgres` directory and run:

```
docker-compose up -d
```

# Configuration

configure your database host, username and password at `dev.env`
