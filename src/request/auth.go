package request

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
)

type (
	UserInfo struct {
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Password   string `json:"password"`
		MacAddress string `json:"mac_address"`
		DeviceName string `json:"device_name"`
		DeviceId   string `json:"device_id"`
	}
	RegisterByPhone struct {
		Phone string `json:"phone"`
		UserInfo
	}
	RegisterByEmail struct {
		Email string `json:"email"`
		UserInfo
	}
	ReSendCode struct {
		VerifyId string `json:"verify_id"` //verify code id
		SendType string `json:"send_type"` // sms call email
	}
	ChangePassword struct {
		VerifyId string `json:"verify_id"` //verify code id
		Code     string `json:"code"`
		User     string `json:"user"`
		Password string `json:"password"`
	}
	ForgotPassword struct {
		User string `json:"user"` //email or phone
	}
	Login struct {
		User       string `json:"user"`
		Password   string `json:"password"`
		MacAddress string `json:"mac_address"`
		DeviceName string `json:"device_name"`
		DeviceId   string `json:"device_id"`
	}
	Verify struct {
		VerifyId   string `json:"verify_id"` //verify code id
		Code       string `json:"code"`
		MacAddress string `json:"mac_address"`
		DeviceName string `json:"device_name"`
		DeviceId   string `json:"device_id"`
	}
)

func (o UserInfo) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.FirstName, validation.Required, validation.NotNil, validation.Length(3, 50)),
		validation.Field(&o.LastName, validation.Required, validation.NotNil, validation.Length(3, 50)),
		validation.Field(&o.Password, validation.Required, validation.NotNil),
		validation.Field(&o.MacAddress, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceName, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceId, validation.Required, validation.NotNil),
	)
	if e != nil {
		return e
	}
	if ok := utils.PasswordRegex(o.Password); !ok {
		return errors.New("invalid password")
	}
	if ok := utils.McAddressRegex(o.MacAddress); !ok {
		return errors.New("invalid mac address")
	}
	return nil
}

func (o RegisterByPhone) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.FirstName, validation.Required, validation.NotNil, validation.Length(3, 50)),
		validation.Field(&o.LastName, validation.Required, validation.NotNil, validation.Length(3, 50)),
		validation.Field(&o.Password, validation.Required, validation.NotNil),
		validation.Field(&o.MacAddress, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceName, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceId, validation.Required, validation.NotNil),
		validation.Field(&o.Phone, validation.Required, validation.NotNil),
	)
	if e != nil {
		return e
	}
	if ok := utils.PasswordRegex(o.Password); !ok {
		return errors.New("invalid password")
	}
	if ok := utils.McAddressRegex(o.MacAddress); !ok {
		return errors.New("invalid mac address")
	}
	if ok := utils.PhoneRegex(o.Phone); !ok {
		return errors.New("invalid phone number")
	}
	return nil
}

func (o RegisterByEmail) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.FirstName, validation.Required, validation.NotNil, validation.Length(3, 50)),
		validation.Field(&o.LastName, validation.Required, validation.NotNil, validation.Length(3, 50)),
		validation.Field(&o.Password, validation.Required, validation.NotNil),
		validation.Field(&o.MacAddress, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceName, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceId, validation.Required, validation.NotNil),
		validation.Field(&o.Email, validation.Required, validation.NotNil),
	)
	if e != nil {
		return e
	}
	if ok := utils.PasswordRegex(o.Password); !ok {
		return errors.New("invalid password")
	}
	if ok := utils.McAddressRegex(o.MacAddress); !ok {
		return errors.New("invalid mac address")
	}
	if ok := utils.EmailRegex(o.Email); !ok {
		return errors.New("invalid email address")
	}
	return nil
}

func (o ReSendCode) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.VerifyId, validation.Required, validation.NotNil),
		validation.Field(&o.SendType, validation.Required, validation.In("sms", "call", "email")),
	)
	if e != nil {
		return e
	}
	//if ok := utils.MongoIDRegex(o.VerifyId); !ok {
	//	return errors.New("invalid verify id")
	//}
	return nil
}

func (o ChangePassword) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.VerifyId, validation.Required, validation.NotNil),
		validation.Field(&o.Code, validation.Required, validation.Length(6, 6)),
		validation.Field(&o.User, validation.Required, validation.NotNil),
		validation.Field(&o.Password, validation.Required, validation.NotNil),
	)
	if e != nil {
		return e
	}
	//if ok := utils.MongoIDRegex(o.VerifyId); !ok {
	//	return errors.New("invalid verify id")
	//}
	if !utils.EmailRegex(o.User) && !utils.PhoneRegex(o.User) {
		return errors.New("invalid user")
	}
	if ok := utils.PasswordRegex(o.Password); !ok {
		return errors.New("invalid password")
	}
	return nil
}

func (o ForgotPassword) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.User, validation.Required, validation.NotNil),
	)
	if e != nil {
		return e
	}
	if !utils.EmailRegex(o.User) && !utils.PhoneRegex(o.User) {
		return errors.New("invalid user")
	}
	return nil
}

func (o Login) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.Password, validation.Required, validation.NotNil),
		validation.Field(&o.MacAddress, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceName, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceId, validation.Required, validation.NotNil),
		validation.Field(&o.User, validation.Required, validation.NotNil),
	)
	if e != nil {
		return e
	}
	if ok := utils.PasswordRegex(o.Password); !ok {
		return errors.New("invalid password")
	}
	if ok := utils.McAddressRegex(o.MacAddress); !ok {
		return errors.New("invalid mac address")
	}
	if !utils.EmailRegex(o.User) && !utils.PhoneRegex(o.User) {
		return errors.New("invalid user")
	}
	return nil
}

func (o Verify) Validate() error {
	e := validation.ValidateStruct(&o,
		validation.Field(&o.MacAddress, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceName, validation.Required, validation.NotNil),
		validation.Field(&o.DeviceId, validation.Required, validation.NotNil),
		validation.Field(&o.VerifyId, validation.Required, validation.NotNil),
		validation.Field(&o.Code, validation.Required, validation.NotNil, validation.Length(6, 6)),
	)
	if e != nil {
		return e
	}
	//if ok := utils.MongoIDRegex(o.VerifyId); !ok {
	//	return errors.New("invalid verify id")
	//}
	if ok := utils.McAddressRegex(o.MacAddress); !ok {
		return errors.New("invalid mac address")
	}
	return nil
}
