package models

type Register struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (a *Register) IsValid() bool {
	return a.User != "" && a.Password != ""
}
