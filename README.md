# Effy-Gravatar

Service used for companies to manage the Job Openings and Applications (JOB PORTAL)

## Technologies

* Go - 1.23
* Gin
* GORM
* Postgresql
* Docker
* AWS EC2 and RDS

## Setup
```
git clone
cd effy-backend
 ```

One time db setup

``Create Database effy-gravatar with user effy-gravatar and password effy-gravatar``

Env File

``create .env file if not present and add below content``

```

DB_HOST = 127.0.0.1
DB_PORT = 5432
DB_USER =  effy-gravatar
DB_PASSWORD = effy-gravatar
DB_NAME = effy-gravatar
DB_SSLMODE = require
AUTH_KEY = "#123&456VR"
```

## Build
 ``make build``

## Run locally

* Database migration is handled using GORM.

`go run /main.go` 
