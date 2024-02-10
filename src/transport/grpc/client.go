package grpc

import (
	"context"
	"github.com/mehrdadjalili/facegram_auth_service/src/service"
	"github.com/mehrdadjalili/facegram_auth_service/src/transport/grpc/pd_auth_client"
)

type (
	clientHandler struct {
		pd_auth_client.UnimplementedAuthClientServiceServer
		authService service.Auth
		userService service.User
	}
)

func (c *clientHandler) UserById(ctx context.Context,
	req *pd_auth_client.ByUserIdRequest) (*pd_auth_client.OneUserResponse, error) {
	user, err := c.userService.ById(req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &pd_auth_client.OneUserResponse{
		User: &pd_auth_client.User{
			Id:          user.ID,
			Email:       user.Email,
			Phone:       user.Phone,
			EmailStatus: user.EmailStatus,
			PhoneStatus: user.PhoneStatus,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Gender:      user.Gender,
			Age:         int32(user.Age),
			Status:      user.Status,
			IsAgent:     user.IsAgent,
			Avatar:      "",
			CreatedAt:   "",
			Username:    user.Username,
		},
	}, nil
}

func (c *clientHandler) CheckToken(ctx context.Context,
	req *pd_auth_client.CheckTokenRequest) (*pd_auth_client.CheckTokenResponse, error) {
	info, err := c.authService.CheckToken(req.GetToken())
	if err != nil {
		return nil, err
	}

	res := pd_auth_client.CheckTokenResponse{
		User: &pd_auth_client.User{
			Id:          info.User.ID,
			Email:       info.User.Email,
			Phone:       info.User.Phone,
			EmailStatus: info.User.EmailStatus,
			PhoneStatus: info.User.PhoneStatus,
			FirstName:   info.User.FirstName,
			LastName:    info.User.LastName,
			Gender:      info.User.Gender,
			Age:         int32(info.User.Age),
			Status:      info.User.Status,
			IsAgent:     info.User.IsAgent,
			Avatar:      "",
			CreatedAt:   "",
			Username:    info.User.Username,
		},
		Session: &pd_auth_client.Session{
			Id:         info.Session.Id,
			DeviceName: info.Session.DeviceName,
			DeviceId:   info.Session.DeviceId,
			MacAddress: info.Session.MacAddress,
			CreatedAt:  info.Session.CreatedAt.Unix(),
			Timestamp:  info.Session.TimeStamp,
			Ip:         info.Session.IP,
		},
	}

	return &res, nil
}
