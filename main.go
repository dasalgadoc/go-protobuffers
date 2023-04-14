package main

import (
	"dasalgadoc.com/go-gprc/application"
	"dasalgadoc.com/go-gprc/domain"
	"dasalgadoc.com/go-gprc/infrastructure"
	"dasalgadoc.com/go-gprc/studentpb"
	"dasalgadoc.com/go-gprc/testpb"
	"fmt"
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

func testUnaryClient(config domain.Config) {
	client := infrastructure.NewGRPCClient(config)
	defer client.CloseClient()
	t := client.DoUnaryRequest("t1")
	fmt.Println("Result: ", t)
}

func testStreamClient(config domain.Config) {
	client := infrastructure.NewGRPCClient(config)
	defer client.CloseClient()
	question := []*testpb.Question{
		{
			Id:       "q8t1",
			Answer:   "blue",
			Question: "Go color?",
			TestId:   "t1",
		},
		{
			Id:       "q9t1",
			Answer:   "google",
			Question: "Go creator?",
			TestId:   "t1",
		},
		{
			Id:       "q9t1",
			Answer:   "backend",
			Question: "Go target?",
			TestId:   "t1",
		},
	}
	m := client.DoClientStreaming(question)
	fmt.Println("Result: ", m)
}

func testServerStreaming(config domain.Config) {
	client := infrastructure.NewGRPCClient(config)
	defer client.CloseClient()
	client.DoServerStreaming("t1")
}

func testBidirectionalStreaming(config domain.Config) {
	client := infrastructure.NewGRPCClient(config)
	defer client.CloseClient()
	client.DoBidirectionalStreaming()
}
