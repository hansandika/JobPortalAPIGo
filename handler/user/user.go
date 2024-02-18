package user

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hansandika/go-job-portal-api/domain/general"
	dto "github.com/hansandika/go-job-portal-api/domain/user"
	"github.com/hansandika/go-job-portal-api/global/constant"
	utils "github.com/hansandika/go-job-portal-api/global/utils/general"
	"github.com/hansandika/go-job-portal-api/usecase"
	uu "github.com/hansandika/go-job-portal-api/usecase/user"
	"github.com/sirupsen/logrus"
)

type UserDataHandler struct {
	Usecase uu.UserDataUsecaseItf
	conf    *general.SectionService
	log     *logrus.Logger
}

func newUserHandler(usecase usecase.Usecase, conf *general.SectionService, logger *logrus.Logger) UserDataHandler {
	return UserDataHandler{
		Usecase: usecase.User.User,
		conf:    conf,
		log:     logger,
	}
}

func (uh UserDataHandler) Login(c *gin.Context) {
	var payload dto.UserLoginRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Error Parsing Json", nil, nil)
		return
	}

	if err := utils.UserLoginValidator.Validate(payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Invalid Request", nil, err)
		return
	}

	// Usecase
	tokenString, err := uh.Usecase.GetUserToken(&payload)
	if err != nil {
		general.CreateResponse(c, err.Code, err.Message, nil, nil)
		return
	}

	// Set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(os.Getenv(constant.JWT_COOKIE_KEY), tokenString.AccessToken, 3600, "", "", false, true)

	general.CreateResponse(c, http.StatusOK, "Success To Login", tokenString, nil)
}

func (uh UserDataHandler) Register(c *gin.Context) {
	var payload dto.UserAddRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Error Parsing Json", nil, nil)
		return
	}

	if err := utils.UserRegisterValidator.Validate(payload); err != nil {
		general.CreateResponse(c, http.StatusBadRequest, "Invalid Request", nil, err)
		return
	}

	// Usecase
	userResponse, err := uh.Usecase.Register(&payload)
	if err != nil {
		general.CreateResponse(c, err.Code, err.Message, nil, nil)
		return
	}

	general.CreateResponse(c, http.StatusOK, "Success To Register", userResponse, nil)
}

func (uh UserDataHandler) Logout(c *gin.Context) {

	c.SetCookie(os.Getenv(constant.JWT_COOKIE_KEY), "", -1, "", "", false, true)

	general.CreateResponse(c, http.StatusOK, "Success To Logout", nil, nil)
}
