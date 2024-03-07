# Shrtify URL Shortener

## Requirements

### Frontend

1. [bun](https://bun.sh/)

1. [golang](https://go.dev/doc/install)

1. [volta](https://volta.sh/)

1. [docker](https://docs.docker.com/engine/install/)

1. [docker-compose](https://docs.docker.com/compose/install/) best use manually

## Installation

`mv frontend/.env.local frontend/.env`

`mv backend/.env.local backend/.env`

`docker-compose up`

`source ./backend/.env && ./scripts/create-database.sh`

`cd frontend && bun install`

`cd backend && go build`

## Start Development

`cd backend && source .env && go run src`

`cd frontend && bun run dev`
