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
		Server        *Server
	}

	applicationRepositories struct {
		studentRepository domain.StudentRepository
	}
)

func BuildApplication() *Application {
	appConfig, err := getConfiguration()
	if err != nil {
		log.Fatalln("error loading configurations: ", err)
		panic(err)
	}

	repositories := buildRepositories(appConfig)

	return &Application{
		Configuration: appConfig,
		Server:        NewServer(repositories.studentRepository),
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

func buildRepositories(config domain.Config) applicationRepositories {
	studentRepo, err := buildStudentRepository(config)
	if err != nil {
		log.Fatalln("error building repo : ", err)
		panic(err)
	}

	return applicationRepositories{
		studentRepository: studentRepo,
	}
}

func buildStudentRepository(config domain.Config) (domain.StudentRepository, error) {
	repo, err := repository.NewPostgresStudentRepository(config.Database)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return repo, err
}
