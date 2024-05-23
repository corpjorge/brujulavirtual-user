package ports

import (
	"brujulavirtual-auth/src/register/domain/models"
)

type Service interface {
	Save(auth models.Register) (models.Register, error)
}
