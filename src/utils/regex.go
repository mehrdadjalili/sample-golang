package utils

import (
	"regexp"
	"unicode"
)

func UUIDRegex(t string) bool {
	var r string = `^[0-9A-Fa-f]{8}(?:-[0-9A-Fa-f]{4}){3}-[0-9A-Fa-f]{12}$`
	return checkRegex(t, r)
}

func PublicKeyRegex(t string) bool {
	var r string = `^-----BEGIN RSA PUBLIC KEY-----(.*)-----END RSA PUBLIC KEY-----$`
	return checkRegex(t, r)
}

func EmailRegex(t string) bool {
	var r string = `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
	return checkRegex(t, r)
}

func PhoneRegex(t string) bool {
	var r string = `^09\d{9}$`
	//r := `/^\+?[1-9][0-9]{7,14}$/`
	return checkRegex(t, r)
}

func MongoIDRegex(t string) bool {
	var r string = `/^(?=[a-f\d]{24}$)(\d+[a-f]|[a-f]+\d)/i`
	return checkRegex(t, r)
}

func McAddressRegex(t string) bool {
	var r string = `^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`
	return checkRegex(t, r)
}

func PasswordRegex(t string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(t) >= 8 && len(t) < 17 {
		hasMinLen = true
	}
	for _, char := range t {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func checkRegex(t string, r string) bool {
	matched, err := regexp.MatchString(r, t)
	if err != nil {
		return false
	}
	return matched
}
