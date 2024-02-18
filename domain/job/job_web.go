package job

import (
	"net/http"

	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/global/exception"
)

type JobAddRequest struct {
	JobTitle    string   `json:"job_title"`
	Description string   `json:"description"`
	Requirement []string `json:"requirement"`
}

func (req JobAddRequest) Validate() *general.Error {
	if req.JobTitle == "" {
		return exception.NewError(http.StatusBadRequest, "JobTitle is required")
	}
	if req.Description == "" {
		return exception.NewError(http.StatusBadRequest, "Description is required")
	}
	if len(req.Requirement) == 0 {
		return exception.NewError(http.StatusBadRequest, "Requirement is required")
	}

	return nil
}

type JobResponse struct {
	JobID       uint     `json:"job_id"`
	JobTitle    string   `json:"job_title"`
	Description string   `json:"description"`
	Requirement []string `json:"requirement"`
}
