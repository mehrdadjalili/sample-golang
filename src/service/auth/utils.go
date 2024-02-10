package auth

import "github.com/mehrdadjalili/facegram_auth_service/src/utils"

func (s *srv) sendVerifyCodeByEmail(email, code string) error {
	return utils.SendEmail(email, code)
}

func (s *srv) sendVerifyCodeByPhone(phone, code string) error {
	return utils.SendSms(phone, code)
}

func (s *srv) sendVerifyCodeByCall(phone, code string) error {
	return nil
}
