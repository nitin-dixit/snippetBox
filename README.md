# Snippetbox

A web application for creating and sharing code snippets, built in Go.

## Features

- Session-based authentication
- Authorization middleware
- Modern CORF protection (`http.CrossOriginProtection`)
- Secure HTTP headers
- Structured logging
- Embedded templates and static assets
- Database migrations
- Unit, handler, middleware, end-to-end, and integration tests

## Tech Stack

- Go 1.26
- PostgreSQL
- pgx
- Goose
- Docker
- HTML templates
- CSS
- JavaScript

## Project Structure

cmd/
internal/
pkg/
ui/
migrations/

## Getting Started

Prerequisites

Docker
Go

Clone repository

...

Start PostgreSQL

docker compose up -d

Run migrations

...

Run application

go run ./cmd/web

Open

<https://localhost:4000>

## Running Tests

go test ./...
