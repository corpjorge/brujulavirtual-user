package ports

import (
	"brujulavirtual-auth/src/register/domain/models"
)

type Repository interface {
	Save(auth models.Register) (models.Register, error)
}
