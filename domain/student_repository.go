package domain

import "context"

type StudentRepository interface {
	GetStudent(ctx context.Context, id string) (*Student, error)
	SetStudent(ctx context.Context, student *Student) error
	SetEnrollment(ctx context.Context, enrollment *Enrollment) error
	GetStudentsPerTest(ctx context.Context, testId string) ([]*Student, error)
}
