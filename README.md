# go-protobuffers
Example project to explore Protocol buffers in go


## 👶🏻 Start with ProtoBuffers

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