


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
├── Dockerfile            ## docker image build할떄
├── build.sh              ## 서버 실행 파일
├── docker-compose.yml    ## docker image 실행할떄
├── go.mod                ## go module 목록
├── go.sum                ## go module checksum
├── handler         
│   └── handler.go        ## handler 함수들
├── main.go               ## 서버 메인 파일
├── middleware
│   └── tokenAuth.go      ## JWT token exp 검증할때
├── persistence
│   ├── db.sql      
│   └── initDB.go         ## DB initialize 
├── routes
│   └── router.go         ## API router
├── types
│   ├── data.go
│   └── user.go
└── utils
    └── common.go


