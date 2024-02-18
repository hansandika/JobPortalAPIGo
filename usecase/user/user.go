package user

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hansandika/go-job-portal-api/domain/general"
	dto "github.com/hansandika/go-job-portal-api/domain/user"
	"github.com/hansandika/go-job-portal-api/global/exception"
	utilsGeneral "github.com/hansandika/go-job-portal-api/global/utils/general"
	utilsJwt "github.com/hansandika/go-job-portal-api/global/utils/jwt"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/hansandika/go-job-portal-api/repository"
	r "github.com/hansandika/go-job-portal-api/repository/user"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserDataUsecaseItf interface {
	GetUserToken(req *dto.UserLoginRequest) (*dto.UserToken, *general.Error)
	Register(req *dto.UserAddRequest) (*dto.UserAddResponse, *general.Error)
}

type UserDataUsecase struct {
	Repo     r.UserDataRepoItf
	Database *infra.DbConnection
	Conf     *general.SectionService
	Log      *logrus.Logger
}

func newUserDataUsecase(r repository.Repo, conf *general.SectionService, logger *logrus.Logger, database *infra.DbConnection) UserDataUsecase {
	return UserDataUsecase{
		Repo:     r.User.User,
		Conf:     conf,
		Log:      logger,
		Database: database,
	}
}

func (ut UserDataUsecase) GetUserToken(req *dto.UserLoginRequest) (*dto.UserToken, *general.Error) {
	// Validate request
	if err := req.Validate(); err != nil {
		ut.Log.Error(err)
		return nil, err
	}

	// Get user by email
	user, err := ut.Repo.GetUserByEmail(req.Email)
	if err != nil {
		ut.Log.Error(err)
		return nil, exception.NewError(http.StatusBadRequest, "Failed to get user by email")
	}

	if user == nil {
		return nil, exception.NewError(http.StatusBadRequest, "User not found")
	}

	// Check password
	isValid := utilsGeneral.CheckPasswordHash(req.Password, user.Password)

	if !isValid {
		return nil, exception.NewError(http.StatusBadRequest, "Invalid password")
	}

	// Generate JWT token
	token, err := utilsJwt.GenerateNewJwtToken(jwt.MapClaims{
		"email":   req.Email,
		"role_id": user.RoleId,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	})

	if err != nil {
		ut.Log.Error(err)
		return nil, exception.NewError(http.StatusBadRequest, "Failed to generate token")
	}

	return &dto.UserToken{
		AccessToken: token,
	}, nil
}

func (ut UserDataUsecase) Register(req *dto.UserAddRequest) (*dto.UserAddResponse, *general.Error) {
	if err := req.Validate(); err != nil {
		ut.Log.Error(err)
		return nil, err
	}

	// Set role name
	req.RoleName = dto.GetUserRoleName(req.RoleId)

	// Check if user already exists
	user, err := ut.Repo.GetUserByEmail(req.Email)
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		ut.Log.Error(err)
		return nil, exception.NewError(http.StatusBadRequest, "Failed to get user by email")
	}

	if user != nil {
		return nil, exception.NewError(http.StatusBadRequest, "User already exists")
	}

	// Hash password
	hash, err := utilsGeneral.GeneratePassword(req.Password)
	if err != nil {
		ut.Log.Error(err)
		return nil, exception.NewError(http.StatusBadRequest, "Failed to hash password")
	}

	// Insert user
	user = &dto.User{
		Email:    req.Email,
		Name:     req.Name,
		Password: hash,
		RoleId:   req.RoleId,
		RoleName: req.RoleName,
	}

	user, err = ut.Repo.Insert(user)
	if err != nil {
		ut.Log.Error(err)
		return nil, exception.NewError(http.StatusInternalServerError, "Failed to insert user")
	}

	// Create response
	userResponse := &dto.UserAddResponse{
		UserId:   user.UserId,
		Email:    user.Email,
		Name:     user.Name,
		RoleId:   user.RoleId,
		RoleName: user.RoleName,
	}

	return userResponse, nil
}
