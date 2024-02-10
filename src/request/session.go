package request

import validation "github.com/go-ozzo/ozzo-validation"

type (
	SetFcmToken struct {
		Token string `json:"token"`
	}
)

func (o SetFcmToken) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Token, validation.Required, validation.NotNil),
	)
}
