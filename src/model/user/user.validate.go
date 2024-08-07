package user

import (
	"context"

	"github.com/kowiste/boilerplate/pkg/validator"

	"github.com/google/uuid"
)

func (u *User) Validate(c context.Context) (err error) {
	u.ID = uuid.NewString()
	validate, err := validator.Get()
	if err != nil {
		return
	}
	return validate.Struct(u)
}
