package config

import (
	"os"

	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/global/constant"
	"github.com/joho/godotenv"
)

func GetCoreConfig() (*general.SectionService, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	data := &general.SectionService{
		App: general.AppAccount{
			Name:         os.Getenv(constant.APP_NAME),
			Environtment: os.Getenv(constant.ENVIRONMENT),
			URL:          os.Getenv(constant.BASE_URL),
			Port:         os.Getenv(constant.APP_PORT),
			Endpoint:     os.Getenv(constant.ENDPOINT_URL),
		},
		Database: general.DatabaseAccount{
			Username: os.Getenv(constant.DB_USERNAME),
			Password: os.Getenv(constant.DB_PASSWORD),
			DBHost:   os.Getenv(constant.DB_HOST),
			Port:     os.Getenv(constant.DB_PORT),
			DBName:   os.Getenv(constant.DB_NAME),
		},
		Authorization: general.AuthAccount{
			Public: general.PublicCredential{
				SecretKey: os.Getenv(constant.JWT_SECRET),
			},
		},
	}
	return data, nil
}
