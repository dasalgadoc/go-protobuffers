package repository

import (
	"context"
	"dasalgadoc.com/go-gprc/domain"
	"database/sql"
	"log"
)

type PostgresTestRepository struct {
	db *sql.DB
}

func NewPostgresTestRepository(url string) (*PostgresTestRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresTestRepository{db: db}, nil
}

func (p *PostgresTestRepository) GetTest(ctx context.Context, id string) (*domain.Test, error) {
	response, err := p.db.QueryContext(ctx,
		"SELECT id, name FROM test WHERE id = $1", id)
	defer func() {
		err = response.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return nil, err
	}

	var test = domain.Test{}
	for response.Next() {
		err = response.Scan(&test.Id, &test.Name)
		if err != nil {
			return nil, err
		}
	}

	return &test, nil
}

func (p *PostgresTestRepository) SetTest(ctx context.Context, test *domain.Test) error {
	_, err := p.db.ExecContext(ctx,
		"INSERT INTO test (id, name) VALUES ($1, $2)",
		test.Id, test.Name)

	return err
}
