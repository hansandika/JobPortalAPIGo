package general

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	utils "github.com/hansandika/go-job-portal-api/global/utils/general"
)

type ResponseData struct {
	Data      interface{}        `json:"data,omitempty"`
	Error     interface{}        `json:"error,omitempty"`
	Status    StatusResponseData `json:"status"`
	TimeStamp time.Time          `json:"timeStamp"`
}

type StatusResponseData struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

func CreateResponse(ctx *gin.Context, code int, message string, data interface{}, err interface{}) {
	RespData := ResponseData{
		Data:  data,
		Error: err,
		Status: StatusResponseData{
			Code:    code,
			Message: message,
		},
		TimeStamp: utils.GenerateTimeNowJakarta(),
	}

	switch code {
	case http.StatusOK:
		RespData.Status.Type = "OK"
	case http.StatusBadRequest:
		RespData.Status.Type = "BAD REQUEST"
	case http.StatusUnauthorized:
		RespData.Status.Type = "UNAUTHORIZED"
	case http.StatusForbidden:
		RespData.Status.Type = "FORBIDDEN"
	case http.StatusNotFound:
		RespData.Status.Type = "NOT FOUND"
	case http.StatusInternalServerError:
		RespData.Status.Type = "INTERNAL SERVER ERROR"
	}

	ctx.JSON(code, RespData)
}
