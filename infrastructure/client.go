package infrastructure

import (
	"context"
	"dasalgadoc.com/go-gprc/domain"
	"dasalgadoc.com/go-gprc/testpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

type Client struct {
	ClientConnection *grpc.ClientConn
}

func NewGRPCClient(config domain.Config) Client {
	target := config.Host + config.Port
	grpcConnection, err := grpc.Dial(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("could not connect: ", err)
	}

	return Client{
		ClientConnection: grpcConnection,
	}
}

func (c *Client) CloseClient() {
	c.ClientConnection.Close()
}

func (c *Client) DoUnaryRequest(id string) *testpb.Test {
	client := testpb.NewTestServiceClient(c.ClientConnection)
	req := &testpb.GetTestRequest{
		Id: id,
	}

	resp, err := client.GetTest(context.Background(), req)
	if err != nil {
		log.Fatalln("error while executed GetTest: ", err)
	}

	return resp
}

func (c *Client) DoClientStreaming(questions []*testpb.Question) *testpb.SetQuestionResponse {
	client := testpb.NewTestServiceClient(c.ClientConnection)
	stream, err := client.SetQuestions(context.Background())
	if err != nil {
		log.Fatalln("error while executed SetQuestions: ", err)
	}

	for _, q := range questions {
		stream.Send(q)
		time.Sleep(2 * time.Second)
	}

	msg, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error on stream: ", err)
	}

	return msg
}

func (c *Client) DoServerStreaming(id string) {
	client := testpb.NewTestServiceClient(c.ClientConnection)
	req := &testpb.GetStudentsPerTestRequest{
		TestId: id,
	}

	stream, err := client.GetStudentsPerTest(context.Background(), req)
	if err != nil {
		log.Fatalln("error: ", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("error reading from streaming")
			break
		}
		log.Println("server response: ", msg)
	}
}

func (c *Client) DoBidirectionalStreaming() {
	client := testpb.NewTestServiceClient(c.ClientConnection)
	answer := testpb.TakeTestRequest{
		Answer: "rta",
	}

	numberOfQuestions := 4

	waitChannel := make(chan struct{})

	stream, err := client.TakeTest(context.Background())
	if err != nil {
		log.Fatalln("error: ", err)
	}

	go func() {
		for i := 0; i < numberOfQuestions; i++ {
			stream.Send(&answer)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("error: ", err)
				break
			}
			log.Println("server response: ", resp)
		}
		close(waitChannel)
	}()
	<-waitChannel
}
