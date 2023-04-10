package repository

import (
	"context"
	"dasalgadoc.com/go-gprc/domain"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type PostgresStudentRepository struct {
	db *sql.DB
}

func NewPostgresStudentRepository(url string) (*PostgresStudentRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresStudentRepository{db: db}, nil
}

func (p *PostgresStudentRepository) GetStudent(ctx context.Context, id string) (*domain.Student, error) {
	response, err := p.db.QueryContext(ctx, "SELECT id, name, age FROM students WHERE id = $1", id)
	defer func() {
		err = response.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return nil, err
	}

	var student = domain.Student{}
	for response.Next() {
		err = response.Scan(&student.Id, &student.Name, &student.Age)
		if err != nil {
			return nil, err
		}
	}

	return &student, nil
}

func (p *PostgresStudentRepository) SetStudent(ctx context.Context, student *domain.Student) error {
	_, err := p.db.ExecContext(ctx,
		"INSERT INTO students (id, name, age) VALUES ($1, $2, $3)",
		student.Id, student.Name, student.Age)

	return err
}

func (p *PostgresStudentRepository) SetEnrollment(ctx context.Context, enrollment *domain.Enrollment) error {
	_, err := p.db.ExecContext(ctx,
		"INSERT INTO enrollments (test_id, student_id) VALUES ($1, $2)",
		enrollment.TestId, enrollment.StudentId)

	return err
}

func (p *PostgresStudentRepository) GetStudentsPerTest(ctx context.Context, testId string) ([]*domain.Student, error) {
	rows, err := p.db.QueryContext(ctx,
		"SELECT id, name, age FROM students WHERE id IN (SELECT student_id FROM enrollments WHERE test_id = $1)",
		testId)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		return nil, err
	}
	var students []*domain.Student
	for rows.Next() {
		var student = domain.Student{}
		if err = rows.Scan(&student.Id, &student.Name, &student.Age); err == nil {
			students = append(students, &student)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}
