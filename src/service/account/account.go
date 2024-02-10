package account

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/entity/models"
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_auth_service/src/response"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

var logSection = "service->account"

func (s *srv) Profile(req *request.ById) (*response.Profile, error) {
	user, err := s.userRepo.ById(req.Id)
	if err != nil {
		return nil, err
	}

	email := ""

	if user.Email != user.ID.Hex() {
		email = user.Email
	}

	return &response.Profile{
		ID:          user.ID.Hex(),
		Username:    user.Username,
		PhoneStatus: user.PhoneStatus,
		EmailStatus: user.EmailStatus,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Gender:      user.Gender,
		Age:         user.Age,
		Phone:       user.Phone,
		Email:       email,
		IsAgent:     user.IsAgent,
		ConnectionStatus: response.ConnectionStatus{
			IsOnline:  user.ConnectionStatus.IsOnline,
			LastVisit: user.ConnectionStatus.LastVisit,
		},
		Statistics: response.UserStatistics{
			TotalLives:           user.Statistics.TotalLives,
			LiveDuration:         user.Statistics.LiveDuration,
			TotalReceiveMessages: user.Statistics.TotalReceiveMessages,
			TotalSendMessages:    user.Statistics.TotalSendMessages,
			Score:                user.Statistics.Score,
			TotalLikes:           user.Statistics.TotalLikes,
			TotalFavorites:       user.Statistics.TotalFavorites,
			TotalProfileVisit:    user.Statistics.TotalProfileVisit,
		},
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *srv) Edit(userId string, req *request.EditAccount) (*response.Regular, error) {
	user, err := s.userRepo.ById(userId)
	if err != nil {
		return nil, err
	}

	var u = user

	u.FirstName = req.FirstName
	u.LastName = req.LastName
	u.Gender = req.Gender
	u.Age = req.Age

	err = s.userRepo.Edit(u)
	if err != nil {
		return nil, err
	}

	return &response.Regular{
		Message: "ok",
	}, nil
}

func (s *srv) SetEmail(userId string, email string) (*response.Regular, error) {
	user, err := s.userRepo.ById(userId)
	if err != nil {
		return nil, err
	}

	if user.Email == email && user.EmailStatus {
		return &response.Regular{
			Message: "AlreadyVerifiedEmail",
		}, nil
	}

	if user.Email != email {
		existsEmail, err := s.userRepo.ExistsEmail(email)
		if err != nil {
			return nil, err
		}
		if existsEmail {
			return &response.Regular{
				Message: "AlreadyExistsEmail",
			}, nil
		}
	}

	code := utils.Create6digitsRandom()
	id := primitive.NewObjectID()
	err = s.verifyCodeRepo.Create(&models.VerifyCode{
		Id:         id,
		UserID:     user.ID,
		TypeCode:   "setEmail",
		CheckCount: 0,
		Code:       strconv.Itoa(code),
		Sender:     "email",
		Receiver:   email,
		Timestamp:  time.Now().Unix() + 300,
	})
	if err != nil {
		return nil, err
	}

	_ = s.sendVerifyCodeByEmail(email, strconv.Itoa(code))

	return &response.Regular{
		Message: id.Hex(),
	}, nil
}

func (s *srv) SetPhone(userId string, phone string) (*response.Regular, error) {
	user, err := s.userRepo.ById(userId)
	if err != nil {
		return nil, err
	}
	if user.Phone == phone && user.PhoneStatus {
		return &response.Regular{
			Message: "AlreadyVerifiedPhone",
		}, nil
	}
	if user.Phone != phone {
		existsPhone, err := s.userRepo.ExistsPhone(phone)
		if err != nil {
			return nil, err
		}
		if existsPhone {
			return &response.Regular{
				Message: "AlreadyExistsPhone",
			}, nil
		}
	}

	code := utils.Create6digitsRandom()
	id := primitive.NewObjectID()
	err = s.verifyCodeRepo.Create(&models.VerifyCode{
		Id:         id,
		UserID:     user.ID,
		TypeCode:   "setPhone",
		CheckCount: 0,
		Code:       strconv.Itoa(code),
		Sender:     "sms",
		Receiver:   phone,
		Timestamp:  time.Now().Unix() + 180,
	})
	if err != nil {
		return nil, err
	}

	_ = s.sendVerifyCodeByPhone(phone, strconv.Itoa(code))

	return &response.Regular{
		Message: id.Hex(),
	}, nil
}

func (s *srv) VerifyEmail(userId string, req *request.VerifyChange) (*response.Regular, error) {
	verifyCode, err := s.verifyCodeRepo.ById(req.VerifyId)
	if err != nil {
		return nil, err
	}
	if verifyCode.UserID.Hex() != userId || verifyCode.TypeCode != "setEmail" {
		return &response.Regular{
			Message: "NotFoundVerifyCode",
		}, nil
	}
	tm := time.Now().Unix()
	if verifyCode.Timestamp < tm {
		err = s.verifyCodeRepo.DeleteById(req.VerifyId)
		if err != nil {
			return nil, err
		}
		return &response.Regular{
			Message: "ExpiredVerifyCode",
		}, nil
	}
	if verifyCode.Code != req.Code {
		if verifyCode.CheckCount > 3 {
			err = s.verifyCodeRepo.DeleteById(req.VerifyId)
			if err != nil {
				return nil, err
			}
		} else {
			v := *verifyCode
			v.CheckCount += 1
			err = s.verifyCodeRepo.Edit(&v)
			if err != nil {
				return nil, err
			}
		}
		return &response.Regular{
			Message: "WrongVerifyCode",
		}, nil
	}

	user, err := s.userRepo.ById(verifyCode.UserID.Hex())
	if err != nil {
		return nil, err
	}

	u := user
	u.Email = verifyCode.Receiver
	u.EmailStatus = true

	existsEmail, err := s.userRepo.ExistsEmail(verifyCode.Receiver)
	if err != nil {
		return nil, err
	}
	if existsEmail {
		return &response.Regular{
			Message: "AlreadyExistsEmail",
		}, nil
	}

	err = s.userRepo.Edit(u)
	if err != nil {
		return nil, err
	}

	err = s.verifyCodeRepo.DeleteById(req.VerifyId)
	if err != nil {
		return nil, err
	}

	return &response.Regular{
		Message: "ok",
	}, nil
}

func (s *srv) VerifyPhone(userId string, req *request.VerifyChange) (*response.Regular, error) {
	verifyCode, err := s.verifyCodeRepo.ById(req.VerifyId)
	if err != nil {
		return nil, err
	}
	if verifyCode.UserID.Hex() != userId || verifyCode.TypeCode != "setPhone" {
		return &response.Regular{
			Message: "NotFoundVerifyCode",
		}, nil
	}
	tm := time.Now().Unix()
	if verifyCode.Timestamp < tm {
		err = s.verifyCodeRepo.DeleteById(req.VerifyId)
		if err != nil {
			return nil, err
		}
		return &response.Regular{
			Message: "ExpiredVerifyCode",
		}, nil
	}
	if verifyCode.Code != req.Code {
		if verifyCode.CheckCount > 3 {
			err = s.verifyCodeRepo.DeleteById(req.VerifyId)
			if err != nil {
				return nil, err
			}
		} else {
			v := *verifyCode
			v.CheckCount += 1
			err = s.verifyCodeRepo.Edit(&v)
			if err != nil {
				return nil, err
			}
		}
		return &response.Regular{
			Message: "WrongVerifyCode",
		}, nil
	}

	user, err := s.userRepo.ById(verifyCode.UserID.Hex())
	if err != nil {
		return nil, err
	}

	u := user
	u.Phone = verifyCode.Receiver
	u.PhoneStatus = true

	existsPhone, err := s.userRepo.ExistsPhone(verifyCode.Receiver)
	if err != nil {
		return nil, err
	}
	if existsPhone {
		return &response.Regular{
			Message: "AlreadyExistsPhone",
		}, nil
	}

	err = s.userRepo.Edit(u)
	if err != nil {
		return nil, err
	}

	err = s.verifyCodeRepo.DeleteById(req.VerifyId)
	if err != nil {
		return nil, err
	}

	return &response.Regular{
		Message: "ok",
	}, nil
}

func (s *srv) ReSendCode(userId string, req *request.ReSendCode) (*response.Regular, error) {
	verifyCode, err := s.verifyCodeRepo.ById(req.VerifyId)
	if err != nil {
		return nil, err
	}
	if verifyCode.UserID.Hex() != userId {
		return &response.Regular{
			Message: "NotFoundVerifyCode",
		}, nil
	}
	if (req.SendType != "email" && verifyCode.Sender == "email") ||
		(verifyCode.Sender != "email" && req.SendType != "voice" && req.SendType != "sms") {
		return &response.Regular{
			Message: "NotFoundVerifyCode",
		}, nil
	}
	v := verifyCode
	v.Timestamp = time.Now().Unix() + 180
	v.Sender = req.SendType

	if req.SendType != "email" {
		_ = s.sendVerifyCodeByEmail(verifyCode.Receiver, verifyCode.Code)
	}

	if req.SendType != "sms" {
		_ = s.sendVerifyCodeByPhone(verifyCode.Receiver, verifyCode.Code)
	}

	if req.SendType != "voice" {
		err = s.sendVerifyCodeByCall(verifyCode.Receiver, verifyCode.Code)
	}

	return &response.Regular{
		Message: "ok",
	}, nil
}

func (s *srv) ChangePassword(userId string, old, new string) (*response.Regular, error) {
	user, err := s.userRepo.ById(userId)
	if err != nil {
		return nil, err
	}
	compare := utils.ComparePassword(old, user.Password)
	if !compare {
		return &response.Regular{
			Message: "IncorrectPassword",
		}, nil
	}
	pass, err := utils.HashPassword(new)
	if err != nil {
		return nil, derrors.InternalError()
	}

	u := user
	u.Password = pass

	err = s.userRepo.Edit(u)
	if err != nil {
		return nil, err
	}

	return &response.Regular{
		Message: "ok",
	}, nil
}
