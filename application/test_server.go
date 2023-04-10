package application

import (
	"context"
	"dasalgadoc.com/go-gprc/domain"
	"dasalgadoc.com/go-gprc/studentpb"
	"dasalgadoc.com/go-gprc/testpb"
	"io"
	"time"
)

type TestServer struct {
	repo         domain.TestRepository
	questionRepo domain.QuestionRepository
	studentRepo  domain.StudentRepository
	testpb.UnimplementedTestServiceServer
}

func NewTestServer(
	repo domain.TestRepository, questionRepo domain.QuestionRepository, studentRepo domain.StudentRepository,
) *TestServer {
	return &TestServer{
		repo:         repo,
		questionRepo: questionRepo,
		studentRepo:  studentRepo,
	}
}

func (s *TestServer) GetTest(ctx context.Context, req *testpb.GetTestRequest) (*testpb.Test, error) {
	test, err := s.repo.GetTest(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &testpb.Test{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetTest(ctx context.Context, req *testpb.Test) (*testpb.SetTestResponse, error) {
	test := &domain.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	err := s.repo.SetTest(ctx, test)
	if err != nil {
		return nil, err
	}
	return &testpb.SetTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) SetQuestions(stream testpb.TestService_SetQuestionsServer) error {
	// Stream processing
	for {
		msg, err := stream.Recv()
		// End of stream
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}
		question := &domain.Question{
			Id:       msg.GetId(),
			Answer:   msg.GetAnswer(),
			Question: msg.GetQuestion(),
			TestId:   msg.GetTestId(),
		}
		err = s.questionRepo.SetQuestion(context.Background(), question)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) EnrollStudents(stream testpb.TestService_EnrollStudentsServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: true,
			})
		}
		if err != nil {
			return err
		}
		enrollment := &domain.Enrollment{
			StudentId: msg.GetStudentId(),
			TestId:    msg.GetTestId(),
		}
		err = s.studentRepo.SetEnrollment(context.Background(), enrollment)
		if err != nil {
			return stream.SendAndClose(&testpb.SetQuestionResponse{
				Ok: false,
			})
		}
	}
}

func (s *TestServer) GetStudentsPerTest(
	req *testpb.GetStudentsPerTestRequest, stream testpb.TestService_GetStudentsPerTestServer) error {
	students, err := s.studentRepo.GetStudentsPerTest(context.Background(), req.GetTestId())
	if err != nil {
		return err
	}

	for _, student := range students {
		s := &studentpb.Student{
			Id:   student.Id,
			Name: student.Name,
			Age:  student.Age,
		}
		err := stream.Send(s)
		time.Sleep(2 * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
