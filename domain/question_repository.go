package domain

import "context"

type QuestionRepository interface {
	SetQuestion(ctx context.Context, question *Question) error
	GetQuestionPerText(ctx context.Context, testId string) ([]*Question, error)
}
