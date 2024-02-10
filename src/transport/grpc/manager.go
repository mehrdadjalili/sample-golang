package grpc

import (
	"context"
	"github.com/mehrdadjalili/facegram_auth_service/proto/pd_auth/pd_auth_manager"
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_auth_service/src/service"
)

type (
	managerHandler struct {
		pd_auth_manager.UnimplementedAuthManagerServiceServer
		userService    service.User
		sessionService service.Session
	}
)

func (m *managerHandler) UserList(ctx context.Context,
	req *pd_auth_manager.UserListRequest) (*pd_auth_manager.UserListResponse, error) {
	list, total, err := m.userService.List(req.Search, req.Sort, int(req.GetPage()), int(req.GetPerPage()))
	if err != nil {
		return nil, err
	}

	var users []*pd_auth_manager.UserResponse
	for _, item := range list {
		users = append(users, &pd_auth_manager.UserResponse{
			Id:          item.ID,
			Email:       item.Email,
			Phone:       item.Phone,
			EmailStatus: item.EmailStatus,
			PhoneStatus: item.PhoneStatus,
			FirstName:   item.FirstName,
			LastName:    item.LastName,
			Gender:      item.Gender,
			Age:         int32(item.Age),
			Status:      item.Status,
			IsAgent:     item.IsAgent,
			Avatar:      "",
			CreatedAt:   "",
			Username:    item.Username,
		})
	}

	return &pd_auth_manager.UserListResponse{
		Count: total,
		Users: users,
	}, nil
}

func (m *managerHandler) UserById(ctx context.Context,
	req *pd_auth_manager.ByIdRequest) (*pd_auth_manager.UserByIdResponse, error) {
	user, err := m.userService.ById(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pd_auth_manager.UserByIdResponse{
		User: &pd_auth_manager.UserResponse{
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

func (m *managerHandler) EditUser(ctx context.Context,
	req *pd_auth_manager.EditUserRequest) (*pd_auth_manager.RegularResponse, error) {
	edit, err := m.userService.Edit(&request.User{
		ID:          req.Id,
		Email:       req.User.Email,
		Phone:       req.User.Phone,
		EmailStatus: req.User.EmailStatus,
		PhoneStatus: req.User.PhoneStatus,
		FirstName:   req.User.FirstName,
		LastName:    req.User.LastName,
		Gender:      req.User.Gender,
		Age:         int(req.User.Age),
		Status:      req.User.Status,
		IsAgent:     req.User.IsAgent,
		Username:    req.User.Username,
	})
	if err != nil {
		return nil, err
	}
	return &pd_auth_manager.RegularResponse{
		Message: edit.Message,
	}, nil
}

func (m *managerHandler) CountByRole(ctx context.Context,
	req *pd_auth_manager.CountByRoleRequest) (*pd_auth_manager.CountByRoleResponse, error) {
	total, err := m.userService.CountByRole(req.IsAgent)
	if err != nil {
		return nil, err
	}
	return &pd_auth_manager.CountByRoleResponse{
		Count: total,
	}, nil
}

func (m *managerHandler) Count(ctx context.Context,
	req *pd_auth_manager.NullRequest) (*pd_auth_manager.CountResponse, error) {
	count, err := m.userService.Count()
	if err != nil {
		return nil, err
	}
	return &pd_auth_manager.CountResponse{
		Count: count,
	}, nil
}

func (m *managerHandler) Statistics(ctx context.Context,
	req *pd_auth_manager.NullRequest) (*pd_auth_manager.StatisticsResponse, error) {
	//stats, err := m.userService.Statistics()
	//if err != nil {
	//	return nil, err
	//}
	return &pd_auth_manager.StatisticsResponse{
		//Count:             stats.Count,
		//Active:            stats.Active,
		//Inactive:          stats.InActive,
		//RegisteredByPhone: stats.RegisteredByPhone,
		//RegisteredByEmail: stats.RegisteredByEmail,
		//Disabled:          stats.Disabled,
	}, nil
}
