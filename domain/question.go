package domain

type Question struct {
	Id       string `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	TestId   string `json:"test_ id"`
}
