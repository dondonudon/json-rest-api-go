Simple Go API

Built with Golang, GORM, Gin Gonic, MySQL

# Migration

    go run migrate/migrate.go

# Run the server

    go run main.go

# Build the server

    go build main.go

## POST

`curl -X POST -H "Content-Type: application/json" -d '{"Title":"test","Content":"test"} http://localhost:{port}/posts`

## GET ALL

`curl -X GET http://localhost:{port}/posts`

## GET ONE

`curl -X GET http://localhost:{port}/posts/{id}`

## PATCH

`curl -X PATCH -H "Content-Type: application/json" -d '{"Title":"test","Content":"test"} http://localhost:{port}/posts/{id}`

## DELETE

`curl -X DELETE http://localhost:{port}/posts/{id}`
