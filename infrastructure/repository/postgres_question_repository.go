package repository

import (
	"context"
	"dasalgadoc.com/go-gprc/domain"
	"database/sql"
)

type PostgresQuestionRepository struct {
	db *sql.DB
}

func NewPostgresQuestionRepository(url string) (*PostgresQuestionRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresQuestionRepository{db: db}, nil
}

func (p *PostgresQuestionRepository) SetQuestion(ctx context.Context, question *domain.Question) error {
	_, err := p.db.ExecContext(ctx,
		"INSERT INTO questions (id, test_id, question, answer) VALUES ($1, $2, $3, $4)",
		question.Id, question.TestId, question.Question, question.Answer)

	return err
}
