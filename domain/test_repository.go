package domain

import "context"

type TestRepository interface {
	GetTest(ctx context.Context, id string) (*Test, error)
	SetTest(ctx context.Context, test *Test) error
}
