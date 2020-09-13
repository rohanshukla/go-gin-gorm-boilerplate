# go-gin-gorm-boilerplate

#### CRUD API for User and Todo Table using GORM and Gin Web Framework 

<div align="center">
    <br />
    <img alt="Go" src="https://golang.org/lib/godoc/images/go-logo-blue.svg" width="100" />
</div>

## Installation & Run
```bash
# Clone this project
$ git clone https://github.com/rohanshukla/go-gin-gorm-boilerplate.git

# Download Gin Framework
$ go get -u github.com/gin-gonic/gin

# Download GORM
$ go get -u gorm.io/gorm

# Download UUID Repo
$ go get github.com/google/uuid

# Run Project
$ go run main.go   or   make dev
```

## Project Structure
```
1. Models   ->  Database Table Schema, and Connection setup in db.go
2. Routers  ->  API end point controllers
3. Utils    ->  Local Storage file and Send Response over HTTP Generic Function
4. main.go  ->  Gin Server Setup and Routes Managements
```