package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/global/constant"
	utilsJwt "github.com/hansandika/go-job-portal-api/global/utils/jwt"
	"github.com/sirupsen/logrus"
)

func JWTMiddleware(conf *general.SectionService, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate jwt token
		tokenString, err := c.Cookie(os.Getenv(constant.JWT_COOKIE_KEY))

		if err != nil {
			log.Error(err)
			general.CreateResponse(c, http.StatusUnauthorized, "Error Getting Token From Cookie", nil, nil)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if tokenString == "" {
			log.Error("Token not found")
			general.CreateResponse(c, http.StatusUnauthorized, "Token not found", nil, nil)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := utilsJwt.ParseToken(tokenString)
		if err != nil {
			log.Error(err)
			general.CreateResponse(c, http.StatusUnauthorized, err.Error(), nil, nil)
			c.AbortWithStatus(401)
			return
		}

		// set claims to context
		c.Set("email", claims["email"])
		c.Set("role_id", claims["role_id"])

		c.Next()
	}
}

func EmployerMiddleware(conf *general.SectionService, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, _ := c.Get("role_id")

		roleId := roleID.(float64)
		roleIDInt := int(roleId)

		if roleIDInt != constant.EmployerRoleId {
			general.CreateResponse(c, http.StatusForbidden, "Forbidden Access", nil, nil)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}
