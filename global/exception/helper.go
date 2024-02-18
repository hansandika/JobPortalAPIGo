package exception

import "github.com/hansandika/go-job-portal-api/domain/general"

func NewError(code int, message string) *general.Error {
	return &general.Error{
		Code:    code,
		Message: message,
	}
}
