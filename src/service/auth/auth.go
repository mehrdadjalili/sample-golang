package auth

import (
	"encoding/json"
	"github.com/mehrdadjalili/facegram_auth_service/resources/messages"
	"github.com/mehrdadjalili/facegram_auth_service/src/entity/models"
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_auth_service/src/response"
	"github.com/mehrdadjalili/facegram_auth_service/src/utils"
	"github.com/mehrdadjalili/facegram_common/pkg/derrors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

var logSection = "service->auth"

func (s *srv) ExistsEmail(req *request.ByEmail) (*response.Status, error) {
	hashEmail := utils.NewSHA256([]byte(req.Email))
	exists, err := s.userRepo.ExistsEmail(hashEmail)
	if err != nil {
		return nil, err
	}
	return &response.Status{
		Status: exists,
	}, nil
}

func (s *srv) ExistsPhone(req *request.ByPhone) (*response.Status, error) {
	hashPhone := utils.NewSHA256([]byte(req.Phone))
	exists, err := s.userRepo.ExistsPhone(hashPhone)
	if err != nil {
		return nil, err
	}
	return &response.Status{
		Status: exists,
	}, nil
}

func (s *srv) RegisterByPhone(req *request.RegisterByPhone) (*response.ResultData, error) {
	exists, err := s.ExistsPhone(&request.ByPhone{
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}
	if exists.Status {
		return nil, derrors.New(derrors.StatusBadRequest, messages.AlreadyExistsPhone)
	}

	pass, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, derrors.InternalError()
	}

	userId := primitive.NewObjectID()

	username := utils.CreateRandomString(10)

	err = s.userRepo.Create(&models.User{
		ID:          userId,
		Username:    username,
		Status:      true,
		PhoneStatus: false,
		EmailStatus: false,
		Password:    pass,
		Phone:       req.Phone,
		Email:       userId.Hex(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Gender:      "",
		IsAgent:     false,
		ConnectionStatus: models.ConnectionStatus{
			IsOnline:  false,
			LastVisit: time.Now(),
		},
	})
	if err != nil {
		return nil, err
	}

	code := utils.Create6digitsRandom()

	id := primitive.NewObjectID()

	err = s.verifyCodeRepo.Create(&models.VerifyCode{
		Id:         id,
		UserID:     userId,
		TypeCode:   "registerByPhone",
		CheckCount: 0,
		Code:       strconv.Itoa(code),
		Sender:     "sms",
		Receiver:   req.Phone,
		Timestamp:  time.Now().Unix() + 180,
	})
	if err != nil {
		return nil, err
	}

	_ = s.sendVerifyCodeByPhone(req.Phone, strconv.Itoa(code))

	return &response.ResultData{
		Result: "ok",
		Data:   id.Hex(),
	}, nil
}

func (s *srv) VerifyRegister(req *request.Verify) (*response.ResultData, error) {
	verifyCode, err := s.verifyCodeRepo.ById(req.VerifyId)
	if err != nil {
		if err.Error() == messages.NotFound {
			return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
		}
		return nil, err
	}

	v := *verifyCode

	if verifyCode.TypeCode != "registerByPhone" {
		return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
	}

	tm := time.Now().Unix()

	if verifyCode.Timestamp < tm {
		err = s.verifyCodeRepo.DeleteById(req.VerifyId)
		if err != nil {
			return nil, err
		}
		return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
	}

	if verifyCode.Code != req.Code {
		if verifyCode.CheckCount > 3 {
			err = s.verifyCodeRepo.DeleteById(req.VerifyId)
			if err != nil {
				return nil, err
			}
		} else {
			v.CheckCount += 1
			err = s.verifyCodeRepo.Edit(&v)
			if err != nil {
				return nil, err
			}
		}
		return nil, derrors.New(derrors.StatusBadRequest, messages.WrongVerifyCode)
	}

	user, err := s.userRepo.ById(verifyCode.UserID.Hex())
	if err != nil {
		return nil, err
	}

	u := *user
	if verifyCode.TypeCode == "registerByPhone" {
		u.PhoneStatus = true
	}

	err = s.userRepo.Edit(&u)
	if err != nil {
		return nil, err
	}

	err = s.verifyCodeRepo.DeleteById(req.VerifyId)
	if err != nil {
		return nil, err
	}

	return &response.ResultData{
		Result: "ok",
	}, nil
}

func (s *srv) ReSendCode(req *request.ReSendCode) (*response.ResultData, error) {
	verifyCode, err := s.verifyCodeRepo.ById(req.VerifyId)
	if err != nil {
		if err.Error() == messages.NotFound {
			return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
		}
		return nil, err
	}

	if (verifyCode.Sender == "email" && req.SendType != "email") ||
		(verifyCode.Sender != "email" && req.SendType != "voice" && req.SendType != "sms") {
		return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
	}

	v := *verifyCode
	if req.SendType == "email" || req.SendType == "voice" {
		v.Timestamp = time.Now().Unix() + 180
	} else {
		v.Timestamp = time.Now().Unix() + 120
	}
	v.Sender = req.SendType

	if req.SendType == "email" {
		_ = s.sendVerifyCodeByEmail(verifyCode.Receiver, verifyCode.Code)
	}

	if req.SendType == "sms" {
		_ = s.sendVerifyCodeByPhone(verifyCode.Receiver, verifyCode.Code)
	}

	if req.SendType == "voice" {
		_ = s.sendVerifyCodeByCall(verifyCode.Receiver, verifyCode.Code)
	}

	err = s.verifyCodeRepo.Edit(&v)
	if err != nil {
		return nil, err
	}

	return &response.ResultData{
		Result: "ok",
	}, nil
}

func (s *srv) ForgotPassword(req *request.ForgotPassword) (*response.ResultData, error) {
	user, err := s.userRepo.ByEmailOrPhone(req.User)
	if err != nil {
		return nil, err
	}

	code := utils.Create6digitsRandom()

	id := primitive.NewObjectID()

	model := models.VerifyCode{
		Id:         id,
		TypeCode:   "forgotPassword",
		UserID:     user.ID,
		CheckCount: 0,
		Code:       strconv.Itoa(code),
		Receiver:   req.User,
	}

	if user.Phone == req.User {
		model.Sender = "sms"
		model.Timestamp = time.Now().Add(time.Minute * 2).Unix()
	}

	if user.Email == req.User {
		model.Sender = "email"
		model.Timestamp = time.Now().Add(time.Minute * 5).Unix()
	}

	err = s.verifyCodeRepo.Create(&model)
	if err != nil {
		return nil, err
	}

	if model.Sender == "sms" {
		_ = s.sendVerifyCodeByPhone(user.Phone, strconv.Itoa(code))
	}

	if model.Sender == "email" {
		_ = s.sendVerifyCodeByPhone(user.Email, strconv.Itoa(code))
	}

	return &response.ResultData{
		Result: "ok",
		Data:   id.Hex(),
	}, nil
}

func (s *srv) ChangePassword(req *request.ChangePassword) (*response.ResultData, error) {
	verifyCode, err := s.verifyCodeRepo.ById(req.VerifyId)
	if err != nil {
		if err.Error() == messages.NotFound {
			return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
		}
		return nil, err
	}
	if verifyCode.Receiver != req.User || verifyCode.TypeCode != "forgotPassword" {
		return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
	}
	tm := time.Now().Unix()
	if verifyCode.Timestamp < tm {
		err = s.verifyCodeRepo.DeleteById(req.VerifyId)
		if err != nil {
			return nil, err
		}
		return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
	}
	if verifyCode.Code != req.Code {
		if verifyCode.CheckCount > 3 {
			err = s.verifyCodeRepo.DeleteById(req.VerifyId)
			if err != nil {
				return nil, err
			}
		}
		v := *verifyCode
		v.CheckCount += 1
		err = s.verifyCodeRepo.Edit(&v)
		if err != nil {
			return nil, err
		}
		return nil, derrors.New(derrors.StatusBadRequest, messages.WrongVerifyCode)
	}

	user, err := s.userRepo.ById(verifyCode.UserID.Hex())
	if err != nil {
		return nil, err
	}

	u := *user

	if (verifyCode.Sender == "sms" || verifyCode.Sender == "voice") && !user.PhoneStatus {
		u.PhoneStatus = true
	}

	if verifyCode.Sender == "email" && !user.EmailStatus {
		u.EmailStatus = true
	}

	pass, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, derrors.InternalError()
	}

	u.Password = pass

	err = s.userRepo.Edit(&u)
	if err != nil {
		return nil, err
	}

	err = s.verifyCodeRepo.DeleteById(req.VerifyId)
	if err != nil {
		return nil, err
	}

	return &response.ResultData{
		Result: "ok",
	}, nil
}

func (s *srv) Login(req *request.Login, ip string) (*response.ResultData, error) {
	user, err := s.userRepo.ByEmailOrPhone(req.User)
	if err != nil {
		if err.Error() == messages.NotFound {
			return nil, derrors.New(derrors.StatusBadRequest, messages.IncorrectUserOrPassword)
		}
		return nil, err
	}
	isEmail := false
	if req.User == user.Email {
		isEmail = true
	}

	if !utils.ComparePassword(req.Password, user.Password) {
		return nil, derrors.New(derrors.StatusBadRequest, messages.IncorrectUserOrPassword)
	}

	if !user.Status {
		return nil, derrors.New(derrors.StatusBadRequest, messages.DisabledAccount)
	}

	res := response.ResultData{}

	if (isEmail && !user.EmailStatus) || (!isEmail && !user.PhoneStatus) {
		code := utils.Create6digitsRandom()
		id := primitive.NewObjectID()
		var info = ""
		var typeCode = ""
		var sender = ""
		var tm int64 = 0
		if isEmail {
			info = user.Email
			typeCode = "loginByEmail"
			sender = "email"
			tm = time.Now().Add(time.Minute * 5).Unix()
		} else {
			info = user.Phone
			typeCode = "loginByPhone"
			sender = "sms"
			tm = time.Now().Add(time.Minute * 2).Unix()
		}
		err = s.verifyCodeRepo.Create(&models.VerifyCode{
			Id:         id,
			UserID:     user.ID,
			TypeCode:   typeCode,
			CheckCount: 0,
			Code:       strconv.Itoa(code),
			Sender:     sender,
			Receiver:   info,
			Timestamp:  tm,
		})
		if err != nil {
			return nil, err
		}
		if isEmail {
			err = s.sendVerifyCodeByEmail(
				info,
				strconv.Itoa(code),
			)
			if err != nil {
				utils.SubmitSentryLog(logSection, "Login", err)
				return nil, err
			}
		} else {
			err = s.sendVerifyCodeByPhone(info, strconv.Itoa(code))
			if err != nil {
				utils.SubmitSentryLog(logSection, "Login", err)
				return nil, err
			}
		}
		res.Result = "sentVerifyCode"
		t := id.Hex()
		res.Data = &t
	} else {
		sid := primitive.NewObjectID()
		err = s.sessionRepo.Create(&models.Session{
			Id:         sid,
			DeviceName: req.DeviceName,
			DeviceId:   req.DeviceId,
			MacAddress: req.MacAddress,
			UserID:     user.ID,
			CreatedAt:  time.Now(),
			TimeStamp:  time.Now().Unix(),
			IP:         ip,
		})
		if err != nil {
			return nil, err
		}
		data := map[string]string{
			"user_id":    user.ID.Hex(),
			"session_id": sid.Hex(),
		}
		js, err := json.Marshal(data)
		if err != nil {
			utils.SubmitSentryLog(logSection, "Login", err)
			return nil, derrors.InternalError()
		}
		enc, err := s.encryption.AesEncrypt(string(js))
		if err != nil {
			utils.SubmitSentryLog(logSection, "Login", err)
			return nil, derrors.InternalError()
		}

		token, err := utils.GenerateJwtToken(s.cfg.Encryption.Key, enc, time.Now().Add(time.Hour*24*30).Unix())
		if err != nil {
			utils.SubmitSentryLog(logSection, "Login", err)
			return nil, derrors.InternalError()
		}
		res.Result = "authorized"
		res.Data = &token
	}

	return &res, nil
}

func (s *srv) VerifyLogin(req *request.Verify, ip string) (*response.ResultData, error) {
	verifyCode, err := s.verifyCodeRepo.ById(req.VerifyId)
	if err != nil {
		if err.Error() == messages.NotFound {
			return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
		}
		return nil, err
	}

	if verifyCode.TypeCode != "loginByEmail" && verifyCode.TypeCode != "loginByPhone" {
		return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
	}

	v := *verifyCode

	tm := time.Now().Unix()

	if verifyCode.Timestamp < tm {
		err = s.verifyCodeRepo.DeleteById(req.VerifyId)
		if err != nil {
			return nil, err
		}
		return nil, derrors.New(derrors.StatusBadRequest, messages.ExpiredVerifyCode)
	}

	if verifyCode.Code != req.Code {
		if verifyCode.CheckCount > 3 {
			err = s.verifyCodeRepo.DeleteById(req.VerifyId)
			if err != nil {
				return nil, err
			}
		} else {
			v.CheckCount += 1
			err = s.verifyCodeRepo.Edit(&v)
			if err != nil {
				return nil, err
			}
		}
		return nil, derrors.New(derrors.StatusBadRequest, messages.WrongVerifyCode)
	}

	user, err := s.userRepo.ById(verifyCode.UserID.Hex())
	if err != nil {
		return nil, err
	}

	u := *user
	if verifyCode.TypeCode == "loginByEmail" {
		u.EmailStatus = true
	}
	if verifyCode.TypeCode == "loginByPhone" {
		u.PhoneStatus = true
	}

	err = s.userRepo.Edit(&u)
	if err != nil {
		return nil, err
	}

	sid := primitive.NewObjectID()

	err = s.sessionRepo.Create(&models.Session{
		Id:         sid,
		DeviceName: req.DeviceName,
		DeviceId:   req.DeviceId,
		MacAddress: req.MacAddress,
		UserID:     user.ID,
		CreatedAt:  time.Now(),
		TimeStamp:  time.Now().Unix(),
		IP:         ip,
	})
	if err != nil {
		return nil, err
	}

	data := map[string]string{
		"user_id":    user.ID.Hex(),
		"session_id": sid.Hex(),
	}

	js, err := json.Marshal(data)
	if err != nil {
		utils.SubmitSentryLog(logSection, "VerifyLogin", err)
		return nil, derrors.InternalError()
	}

	enc, err := s.encryption.AesEncrypt(string(js))
	if err != nil {
		utils.SubmitSentryLog(logSection, "VerifyLogin", err)
		return nil, derrors.InternalError()
	}

	token, err := utils.GenerateJwtToken(s.cfg.Encryption.Key, enc, time.Now().Add(time.Hour*24*30).Unix())
	if err != nil {
		utils.SubmitSentryLog(logSection, "VerifyLogin", err)
		return nil, derrors.InternalError()
	}

	err = s.verifyCodeRepo.DeleteById(req.VerifyId)
	if err != nil {
		return nil, err
	}

	return &response.ResultData{
		Result: "authorized",
		Data:   &token,
	}, nil
}

func (s *srv) CheckToken(token string) (*response.CheckToken, error) {
	data, err := utils.ParseJwtToken(s.cfg.Encryption.Key, token)
	if err != nil {
		utils.SubmitSentryLog(logSection, "CheckToken", err)
		return nil, err
	}
	dec, err := s.encryption.AesDecrypt(data)

	var userIfo map[string]interface{}

	err = json.Unmarshal([]byte(dec), &userIfo)
	if err != nil {
		utils.SubmitSentryLog(logSection, "CheckToken", err)
		return nil, derrors.InternalError()
	}

	userId, existsUserId := userIfo["user_id"]
	sessionId, existsSessionId := userIfo["session_id"]

	if !existsUserId || !existsSessionId {
		return nil, derrors.New(derrors.StatusUnauthorized, messages.Unauthorized)
	}

	hashSessionId := utils.NewSHA256([]byte(sessionId.(string)))

	redisToken, err := s.redis.GetString("tokens:" + hashSessionId)
	if err != nil {
		return nil, derrors.InternalError()
	}

	if redisToken != "" {
		decode, err := s.encryption.AesDecrypt(redisToken)
		if err == nil {
			var out response.CheckToken
			err = json.Unmarshal([]byte(decode), &out)
			if err == nil && out.User.ID == userId && out.Session.Id == sessionId {
				return &out, nil
			}
		}
	}

	session, err := s.sessionRepo.ById(sessionId.(string))
	if err != nil {
		if err.Error() == messages.NotFound {
			return nil, derrors.New(derrors.StatusUnauthorized, messages.Unauthorized)
		}
		return nil, derrors.InternalError()
	}
	user, err := s.userRepo.ById(userId.(string))
	if err != nil {
		if err.Error() == messages.NotFound {
			return nil, derrors.New(derrors.StatusUnauthorized, messages.Unauthorized)
		}
		return nil, derrors.InternalError()
	}

	phone := user.Phone
	email := ""

	if user.Email != userId {
		email = user.Email
	}

	userInfo := response.UserInfo{
		ID:          user.ID.Hex(),
		Username:    user.Username,
		PhoneStatus: user.PhoneStatus,
		EmailStatus: user.EmailStatus,
		Phone:       phone,
		Email:       email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		CreatedAt:   user.CreatedAt,
		IsAgent:     user.IsAgent,
		Gender:      user.Gender,
		Age:         user.Age,
	}

	sessionInfo := response.SessionInfo{
		Id:         session.Id.Hex(),
		DeviceName: session.DeviceName,
		DeviceId:   session.DeviceId,
		MacAddress: session.MacAddress,
		CreatedAt:  session.CreatedAt,
		TimeStamp:  session.TimeStamp,
		IP:         session.IP,
	}

	inf := response.CheckToken{User: userInfo, Session: sessionInfo}

	js, err := json.Marshal(inf)
	if err == nil {
		str, err := s.encryption.AesEncrypt(string(js))
		if err == nil {
			hash := utils.NewSHA256([]byte(sessionId.(string)))
			_ = s.redis.SetString("tokens:"+hash, str, time.Duration(time.Hour*24*30))
		} else {
			utils.SubmitSentryLog(logSection, "CheckToken", err)
		}
	}

	return &inf, nil
}
