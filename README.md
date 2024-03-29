# Educabot Challenge
Please refer to the newest repository instead, here:  
```
https://gitlab.com/federicomatias.aguero1987
```

## Prerequisites
you should create an .env file with the following variables
* DB_USER= postgres
* DB_HOST= the host of your database
* DB_NAME= the name of your database
* DB_PORT= the port of your database

## Libraries
* Golang 1.15.x
* jwt-go 
* gin v1.6.3
* godotenv v1.4.0
* smapping 
* v1.21.0
* crypto 
* postgres v1.3.8
* sqlite v1.3.6
* gorm v1.23.8

Structure project
```
.
├── internal
│   ├── api
│   │    ├──main.go
│   ├── server
│   │     ├──server.go
│   │     ├──wire.go
├── internal
│   ├── domain
│   │   ├── dto
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   │                 
│   │   ├── entity
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   ├── platform
│   │   ├── handler
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   ├── helper
│   │   │    ├──errors
│   │   │    │    ├──error.go
│   │   │    ├──response
│   │   │    │    ├──response.go
│   │   ├── middleware
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   ├── service
│   │   │    ├──driverSearch-dto.go
│   │   │    ├──driverSearch-dto.go
│   │   ├── storage
│   │   │    ├──repository
│   │   │    │    ├──error.go
│   │   │    ├──driverSearch-dto.go
│   │   ├── utils
│   │   │    ├──utils.go
├── go.mod
├── go.sum
```