package main

import (
	"dasalgadoc.com/go-gprc/application"
	"dasalgadoc.com/go-gprc/studentpb"
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

	s := grpc.NewServer()
	studentpb.RegisterStudentServiceServer(s, app.Server)

	// to metadata to facilitate consume
	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatalln(err)
	}
}
