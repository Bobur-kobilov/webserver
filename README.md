


### Getting Started
A simple Webserver written in golang

### Prerequisites
* Golang -> 1.15.1
* Docker -> 19.03.12


### 실행
```bash
  $ ./build.sh
```

### Directory Structure
```
├── Dockerfile
├── build.sh
├── docker-compose.yml
├── go.mod
├── go.sum
├── handler
│   └── handler.go
├── main
├── main.go
├── middleware
│   └── tokenAuth.go
├── persistence
│   ├── db.sql
│   └── initDB.go
├── routes
│   └── router.go
├── types
│   ├── data.go
│   └── user.go
└── utils
    └── common.go


