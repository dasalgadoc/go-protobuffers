<h1 align="center">
  🚀 🐹 Go ProtoBuffers & gRPC example 
</h1>

<p align="center">
    <a href="#"><img src="https://img.shields.io/badge/technology-go-blue.svg" alt="Go"/></a>
</p>

<p align="center">
  This repository serves as an example of ProtoBuffers and gRPC. 
  It was created to investigate some techniques using Go.
</p>

## 🧲 Environment Setup

### 🛠️ Needed tools

1. Go 1.20.2 or higher
2. Docker and Docker compose (I use Docker version 23.01.1 and docker-compose v2.17.0)

### 🏃🏻 Application execution

1. Make sure to download all Needed tools
2. Clone the repository
```bash
git clone https://github.com/dasalgadoc/go-protobuffers.git
```
3. Build up go project
```bash
go mod download
go get .
```
4. The project uses Docker to manage a Postgres database and localhost to compile and run the Go source code.
```bash
docker-compose up --build 
```
5. Run the API
```bash
go run main.go
```
6. Now, you can consume the gRPC

## 👶🏻 Notes: Start with ProtoBuffers

1. [Install](https://grpc.io/docs/protoc-installation/) the Protobuffers compiler (protoc)
2. Some plugins can facilitate Protobuffers and gRPC
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
3. Define the .proto file where you can define message (data structures) and services (signatures of methods)
    - [Protobuffer](https://protobuf.dev/getting-started/gotutorial/)
    - [gRPC](https://grpc.io/docs/languages/go/basics/)
4. Compile the .proto file using protoc to generate the Go source code
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <FOLDER>/<FILE>.proto
```
5. The Source code generated required this dependency in the project
```bash
go get google.golang.org/protobuf
go get google.golang.org/grpc
```
6. After this you can use the Services and Structs in your source code.