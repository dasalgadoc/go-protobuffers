package application

import (
	"context"
	"dasalgadoc.com/go-gprc/domain"
	"dasalgadoc.com/go-gprc/studentpb"
)

type Server struct {
	repo domain.StudentRepository
	studentpb.UnimplementedStudentServiceServer
}

func NewServer(repo domain.StudentRepository) *Server {
	return &Server{
		repo: repo,
	}
}

func (s *Server) GetStudent(ctx context.Context, req *studentpb.GetStudentRequest) (*studentpb.Student, error) {
	student, err := s.repo.GetStudent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &studentpb.Student{
		Id:   student.Id,
		Name: student.Name,
		Age:  student.Age,
	}, nil
}

func (s *Server) SetStudent(ctx context.Context, req *studentpb.Student) (*studentpb.SetStudentResponse, error) {
	student := &domain.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}
	err := s.repo.SetStudent(ctx, student)
	if err != nil {
		return nil, err
	}
	return &studentpb.SetStudentResponse{
		Id: student.Id,
	}, nil
}
