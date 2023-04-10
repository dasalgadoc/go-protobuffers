package application

import (
	"dasalgadoc.com/go-gprc/domain"
	"dasalgadoc.com/go-gprc/infrastructure/repository"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type (
	Application struct {
		Configuration domain.Config
		StudentServer *StudentServer
		TestServer    *TestServer
	}

	applicationRepositories struct {
		studentRepository  domain.StudentRepository
		testRepository     domain.TestRepository
		questionRepository domain.QuestionRepository
	}
)

func BuildApplication() *Application {
	appConfig, err := getConfiguration()
	if err != nil {
		log.Fatalln("error loading configurations: ", err)
		panic(err)
	}

	repositories, err := buildRepositories(appConfig)
	if err != nil {
		log.Fatalln("error building repos: ", err)
		panic(err)
	}

	return &Application{
		Configuration: appConfig,
		StudentServer: NewStudentServer(repositories.studentRepository),
		TestServer: NewTestServer(repositories.testRepository,
			repositories.questionRepository,
			repositories.studentRepository),
	}
}

func getConfiguration() (domain.Config, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return domain.Config{}, err
	}
	config := domain.Config{
		Database: os.Getenv("DATABASE_URL"),
		Port:     os.Getenv("PORT"),
		Network:  os.Getenv("NETWORK"),
	}
	if err = config.ConfigErrors(); err != nil {
		return domain.Config{}, err
	}
	return config, nil
}

func buildRepositories(config domain.Config) (applicationRepositories, error) {
	studentRepo, err := repository.NewPostgresStudentRepository(config.Database)
	if err != nil {
		return applicationRepositories{}, err
	}

	testRepo, err := repository.NewPostgresTestRepository(config.Database)
	if err != nil {
		return applicationRepositories{}, err
	}

	questionRepo, err := repository.NewPostgresQuestionRepository(config.Database)
	if err != nil {
		return applicationRepositories{}, err
	}

	return applicationRepositories{
		studentRepository:  studentRepo,
		testRepository:     testRepo,
		questionRepository: questionRepo,
	}, nil
}
