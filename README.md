# Example-CRUD-Vue-3-Go
- The example shows how to Building a Vue CRUD App with a Go API and using MySQL as a database.
- The article of this repository https://blog.stackpuz.com/building-a-vue-crud-app-with-a-go-api
- To find more resources, please visit https://stackpuz.com

## Prerequisites
- Node.js
- Go 1.21
- MySQL

## Installation
- Clone this repository `git clone https://github.com/stackpuz/Example-CRUD-Vue-3-Go .`
- Change directory to Vue project. `cd view`
- Install the Vue dependencies. `npm install`
- Change directory to Go project. `cd ../api`
- Install the Go dependencies. `go mod tidy`
- Create a new database and run [/database.sql](/database.sql) script to create tables and import data.
- Edit the database configuration in [/api/.env](/api/.env) file.

## Run project

- Run Vue project. `npm run dev`
- Run Go API project `go run main.go`
- Navigate to http://localhost:5173