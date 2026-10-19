package in

import (
	"root_academy_go/internal/user/adapter/in/web/dto"
)

type RegisterUser interface {
	Register(*dto.Register) error
}
