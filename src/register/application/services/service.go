package services

import (
	"brujulavirtual-auth/src/register/domain/models"
	"brujulavirtual-auth/src/register/domain/ports"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Service struct {
	repository ports.Repository
}

func Register(repo ports.Repository) ports.Service {
	return &Service{repo}
}

func (service *Service) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (service *Service) Save(auth models.Register) (models.Register, error) {

	storedAuth, err := service.repository.Save(auth)

	if err != nil {
		return models.Register{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedAuth.Password), []byte(auth.Password))
	if err != nil {
		log.Default().Println("Error: ", err)
		return models.Register{}, err
	}

	return storedAuth, nil
}
