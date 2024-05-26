package interfaces

import ("github.com/nomad-kzn/template/models")

type (
	Usecase interface {
		GetUser(id string) (models.User, error)
	}
)