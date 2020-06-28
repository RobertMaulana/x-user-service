package services

import (
	"github.com/RobertMaulana/x-comment-service/utils/errors"
	"github.com/RobertMaulana/x-user-service/domain/users"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	GetMembers(int64) (*users.ApiGeneralResponse, *errors.RestErr)
}

func (s *userService) GetMembers(organizationId int64) (*users.ApiGeneralResponse, *errors.RestErr) {
	dao := &users.Organization{
		Id: organizationId,
	}
	res, err := dao.GetOrganizationMembers()
	if err != nil {
		return nil, err
	}
	return res, nil
}
