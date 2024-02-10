package user

import (
	"github.com/mehrdadjalili/facegram_auth_service/src/request"
	"github.com/mehrdadjalili/facegram_auth_service/src/response"
)

func (s *srv) ById(id string) (*response.User, error) {
	user, err := s.userRepo.ById(id)
	if err != nil {
		return nil, err
	}

	email := ""
	if user.Email != id {
		email = user.Email
	}

	return &response.User{
		ID:          user.ID.Hex(),
		Username:    user.Username,
		Email:       email,
		Phone:       user.Phone,
		EmailStatus: user.EmailStatus,
		PhoneStatus: user.PhoneStatus,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		IsAgent:     user.IsAgent,
		Gender:      user.Gender,
		Age:         user.Age,
		Status:      user.Status,
	}, nil
}

func (s *srv) Edit(data *request.User) (*response.Regular, error) {
	user, err := s.userRepo.ById(data.ID)
	if err != nil {
		return nil, err
	}

	u := *user

	if data.Phone != "" {
		u.Phone = data.Phone
	}

	if data.Email != "" {
		u.Email = data.Email
	}

	u.Status = user.Status
	u.IsAgent = user.IsAgent
	u.PhoneStatus = user.PhoneStatus
	u.EmailStatus = user.EmailStatus
	u.Password = user.Password
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Gender = user.Gender
	u.Age = user.Age

	err = s.userRepo.Edit(&u)
	if err != nil {
		return nil, err
	}

	return &response.Regular{Message: "ok"}, nil
}

func (s *srv) CountByRole(isAgent bool) (int64, error) {
	return s.userRepo.CountByRole(isAgent)
}

func (s *srv) Count() (int64, error) {
	return s.userRepo.Count()
}

func (s *srv) List(search, sort string, page, perPage int) ([]response.User, int64, error) {
	lst, total, err := s.userRepo.List(search, sort, page, perPage)
	if err != nil {
		return nil, 0, err
	}
	var list []response.User
	for _, user := range lst {
		email := ""
		if user.Email != user.ID.Hex() {
			email = user.Email
		}

		list = append(list, response.User{
			ID:          user.ID.Hex(),
			Username:    user.Username,
			Email:       email,
			Phone:       user.Phone,
			EmailStatus: user.EmailStatus,
			PhoneStatus: user.PhoneStatus,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Gender:      user.Gender,
			Age:         user.Age,
			Status:      user.Status,
			IsAgent:     user.IsAgent,
		})
	}

	return list, total, nil
}
