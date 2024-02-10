package request

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
)

type (
	ByUserId struct {
		UserId string `json:"user_id"`
	}
	ByKey struct {
		Key string `json:"access_key"`
	}
	MessageParams struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	ById struct {
		Id string `json:"id"`
	}
	BasicList struct {
		Page    int `json:"page"`
		PerPage int `json:"per_page"`
	}
	ByEmail struct {
		Email string `json:"email"`
	}
	ByPhone struct {
		Phone string `json:"phone"`
	}
)

func (o ByUserId) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.UserId, validation.Required, validation.NotNil),
	)
}

func (o ByKey) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Key, validation.Required, validation.NotNil),
	)
}

func (o ById) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Id, validation.Required, validation.NotNil),
	)
}

func (o BasicList) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Page, validation.Required, validation.Min(0)),
		validation.Field(&o.PerPage, validation.Required, validation.Min(1), validation.Max(50)),
	)
}

func (o ByEmail) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.Email, validation.Required, validation.NotNil),
	)
	if e != nil {
		return e
	}
	i := utils.EmailRegex(o.Email)
	if !i {
		return errors.New("invalid email address")
	}
	return nil
}

func (o ByPhone) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.Phone, validation.Required, validation.NotNil),
	)
	if e != nil {
		return e
	}
	i := utils.PhoneRegex(o.Phone)
	if !i {
		return errors.New("invalid phone number")
	}
	return nil
}
