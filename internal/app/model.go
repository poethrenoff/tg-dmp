package app

type Profile struct {
	Id        string `json:"id"`
	Interests string `json:"interests"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}
