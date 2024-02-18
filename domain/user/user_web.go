package user

import (
	"net/http"

	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/global/constant"
	"github.com/hansandika/go-job-portal-api/global/exception"
)

type UserAddRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

func (req UserAddRequest) Validate() *general.Error {
	if req.Email == "" {
		return exception.NewError(http.StatusBadRequest, "Email is required")
	}
	if req.Name == "" {
		return exception.NewError(http.StatusBadRequest, "Name is required")
	}
	if req.Password == "" {
		return exception.NewError(http.StatusBadRequest, "Password is required")
	}

	if req.RoleId < 1 || req.RoleId > 2 {
		return exception.NewError(http.StatusBadRequest, "Role Id is invalid")
	}

	return nil
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req UserLoginRequest) Validate() *general.Error {
	if req.Email == "" {
		return exception.NewError(http.StatusBadRequest, "Email is required")
	}
	if req.Password == "" {
		return exception.NewError(http.StatusBadRequest, "Password is required")
	}

	return nil
}

type UserToken struct {
	AccessToken string `json:"access_token"`
}

type UserAddResponse struct {
	UserId   uint   `json:"user_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

func GetUserRoleName(roleId int) string {
	switch roleId {
	case constant.TalentRoleId:
		return constant.TalentRoleName
	case constant.EmployerRoleId:
		return constant.EmployerRoleName
	default:
		return ""
	}
}
