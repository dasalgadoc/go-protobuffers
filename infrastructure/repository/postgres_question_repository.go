package repository

import (
	"context"
	"dasalgadoc.com/go-gprc/domain"
	"database/sql"
	"log"
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

func (p *PostgresQuestionRepository) GetQuestionPerText(ctx context.Context, testId string) ([]*domain.Question, error) {
	rows, err := p.db.QueryContext(ctx,
		"SELECT id, question FROM questions WHERE test_id = $1", testId)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return nil, err
	}
	var questions []*domain.Question
	for rows.Next() {
		var question = domain.Question{}
		if err = rows.Scan(&question.Id, &question.Question); err == nil {
			questions = append(questions, &question)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}
