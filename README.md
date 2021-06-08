# Go Web
[![License MIT](https://img.shields.io/badge/License-MIT-blue.svg)](http://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/raychongtk/go-web)](https://goreportcard.com/report/github.com/raychongtk/go-web)

a demo for building web application in Go

## Tech Stack

1. Dependency Injection - Wire
2. HTTP Web Framework - Gin
3. Session Management - Gin-Session, Redis
4. Database ORM - Gorm
5. Database - Postgres
6. Load Configuration - Viper

# Secure Cookies

- SameSite (LAX Mode): cookies are sent when the request is navigating to the original site
- Secure: cookies only sent when using https connection
- HttpOnly: client side cannot touch the cookies

# Spin up a Postgres instance with Docker

go to `.postgres` directory and run:

```
export POSTGRES_PASSWORD=your password
docker-compose up -d
```

# Start up a Redis instance for session

please use Docker to install the Redis image

# Configuration

configure your database host, username and password at `dev.env`
