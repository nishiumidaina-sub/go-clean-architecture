# go-clean-architecture
Go/Echoでクリーンアーキテクチャを頑張るリポジトリ
## 構成
```
observation-system-backend
 ├──docker
 |   └──go
 |      └──Dockerfile
 ├──docs
 |   ├─migration
 |   |  ├──000001_create_example.up.sql
 |   |  └──000001_create_example.down.sql
 |   └──swagger
 |      ├──docs.go
 |      ├──swagger.json
 |      └──swagger.yaml
 ├──domain(Entities)
 |   └──exampleDomain.go
 ├──infrastructure
 |   └──sqlhandler.go
 ├──lib
 ├──log
 |   ├──api_debug.log
 |   └──log.go
 ├──presentation
 |   ├──controller
 |   |   ├──context.go
 |   |   └──exampleControlle.go
 |   ├──database
 |   |   ├──sqlhandler.go
 |   |   └──exampleRepository.go
 |   ├──response
 |   |   ├──errorResponse.go
 |   |   └──exampleResponse.go
 |   └──router
 |       └──outer.go
 ├──usecase
 |   ├──exampleInteractor.go
 |   └──exampleRepository.go
 ├──.air.toml
 ├──.env.example
 ├──.gitignore
 ├──docker-compose.yaml
 ├──go.mod
 ├──go.sum
 └──README.md
```
