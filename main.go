package main

import (
	"dasalgadoc.com/go-gprc/application"
	"dasalgadoc.com/go-gprc/studentpb"
	"dasalgadoc.com/go-gprc/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	var app = application.BuildApplication()
	list, err := net.Listen(app.Configuration.Network, app.Configuration.Port)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Initializing server...")

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, app.StudentServer)
	testpb.RegisterTestServiceServer(s, app.TestServer)

	// to metadata to facilitate consume
	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatalln(err)
	}
}
