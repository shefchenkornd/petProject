#!/bin/bash
export HTTP_ADDR=localhost:8080

### DB settings

# Postgres settings
#export PG_URL=postgres://postgres:postgres@localhost/test?sslmode=disable
#export PG_MIGRATIONS_PATH=file://../../store/pg/migrations

# MySQL settings
#export MYSQL_ADDR=127.0.0.1:3306
export MYSQL_ADDR=mysql:3306
export MYSQL_USER=root
export MYSQL_PASSWORD=root
export MYSQL_DB=pet_project
export MYSQL_MIGRATIONS_PATH=file://../../store/mysql/migrations

# Logger settings
export LOG_LEVEL=debug
