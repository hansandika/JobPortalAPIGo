package user

import (
	do "github.com/hansandika/go-job-portal-api/domain/user"
	"github.com/hansandika/go-job-portal-api/infra"
	"github.com/sirupsen/logrus"
)

type UserDataRepo struct {
	DB  *infra.DbConnection
	Log *logrus.Logger
}

func newUserDataRepo(db *infra.DbConnection, logger *logrus.Logger) UserDataRepo {
	return UserDataRepo{
		DB:  db,
		Log: logger,
	}
}

type UserDataRepoItf interface {
	GetUserByEmail(email string) (*do.User, error)
	GetUserByUserId(userId uint) (*do.User, error)
	Insert(User *do.User) (*do.User, error)
}

func (ur UserDataRepo) GetUserByEmail(email string) (*do.User, error) {
	var user do.User
	if err := ur.DB.Database.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		ur.Log.Error(err)
		return nil, err
	}

	return &user, nil
}

func (ur UserDataRepo) GetUserByUserId(userId uint) (*do.User, error) {
	var user do.User
	if err := ur.DB.Database.Table("users").Where("user_id = ?", userId).First(&user).Error; err != nil {
		ur.Log.Error(err)
		return nil, err
	}

	return &user, nil
}

func (ur UserDataRepo) Insert(user *do.User) (*do.User, error) {
	if err := ur.DB.Database.Create(&user).Error; err != nil {
		ur.Log.Error(err)
		return nil, err
	}

	return user, nil
}
