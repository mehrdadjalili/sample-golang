package request

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
)

type (
	EditAccount struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Gender    string `json:"gender"`
		Age       int    `json:"age"`
	}
	SetTwoAuthentication struct {
		Status bool `json:"status"`
	}
	VerifyChange struct {
		VerifyId string `json:"verify_id"`
		Code     string `json:"code"`
	}
	ChangeAccountPassword struct {
		Old string `json:"old"`
		New string `json:"new"`
	}
)

func (o EditAccount) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.FirstName, validation.Required, validation.NotNil, validation.Length(3, 50)),
		validation.Field(&o.LastName, validation.Required, validation.NotNil, validation.Length(3, 50)),
		validation.Field(&o.Gender, validation.Required, validation.NotNil, validation.In("f", "m")),
		validation.Field(&o.Age, validation.Required, validation.Min(7), validation.Max(150)),
	)
}

func (o SetTwoAuthentication) Validate() error {
	return validation.ValidateStruct(&o) //validation.Field(&o.Status, validation.Required, validation.In(false, true)),

}

func (o VerifyChange) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.VerifyId, validation.Required, validation.NotNil),
		validation.Field(&o.Code, validation.Required, validation.NotNil, validation.Length(6, 6)),
	)
	if e != nil {
		return e
	}
	i := utils.MongoIDRegex(o.VerifyId)
	if !i {
		return errors.New("invalid verify id")
	}
	return nil
}

func (o ChangeAccountPassword) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.Old, validation.Required, validation.NotNil),
		validation.Field(&o.New, validation.Required, validation.NotNil),
	)
	if e != nil {
		return e
	}
	i := utils.PasswordRegex(o.New)
	if !i {
		return errors.New("invalid new password")
	}
	return nil
}
