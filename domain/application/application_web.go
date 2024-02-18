package application

import (
	"net/http"
	"net/url"
	"time"

	"github.com/hansandika/go-job-portal-api/domain/general"
	"github.com/hansandika/go-job-portal-api/global/constant"
	"github.com/hansandika/go-job-portal-api/global/exception"
)

type ApplicationAddRequest struct {
	JobId            uint   `json:"job_id"`
	Email            string `json:"email"`
	Status           int    `json:"status"`
	CvLink           string `json:"cv_link"`
	CoverLetterLink  string `json:"cover_letter_link"`
	YearOfExperience int    `json:"year_of_experience"`
	IsActive         int    `json:"is_active"`
}

func (req ApplicationAddRequest) Validate() *general.Error {
	if req.JobId == 0 {
		return exception.NewError(http.StatusBadRequest, "JobId is required")
	}
	if req.Email == "" {
		return exception.NewError(http.StatusBadRequest, "Email is required")
	}
	if req.Status == 0 {
		return exception.NewError(http.StatusBadRequest, "Status is required")
	}
	if req.IsActive != 0 && req.IsActive != 1 {
		return exception.NewError(http.StatusBadRequest, "IsActive is required and must be 0 or 1")
	}
	if req.CvLink == "" {
		return exception.NewError(http.StatusBadRequest, "CvLink is required")
	}
	if req.CoverLetterLink == "" {
		return exception.NewError(http.StatusBadRequest, "CoverLetterLink is required")
	}
	if req.YearOfExperience == 0 {
		return exception.NewError(http.StatusBadRequest, "YearOfExperience is required")
	}

	_, err := url.ParseRequestURI(req.CvLink)
	if err != nil {
		return exception.NewError(http.StatusBadRequest, "Invalid CvLink")
	}

	_, err = url.ParseRequestURI(req.CoverLetterLink)
	if err != nil {
		return exception.NewError(http.StatusBadRequest, "Invalid CoverLetterLink")
	}

	return nil
}

type ApplicationUpdateRequest struct {
	JobId    uint   `json:"job_id"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	IsActive int    `json:"is_active"`
}

func (req ApplicationUpdateRequest) Validate() *general.Error {
	if req.JobId == 0 {
		return exception.NewError(http.StatusBadRequest, "JobId is required")
	}
	if req.Email == "" {
		return exception.NewError(http.StatusBadRequest, "Email is required")
	}
	if req.Status == 0 {
		return exception.NewError(http.StatusBadRequest, "Status is required")
	}
	if req.IsActive != 0 && req.IsActive != 1 {
		return exception.NewError(http.StatusBadRequest, "IsActive is required and must be 0 or 1")
	}

	return nil
}

type ApplicationResponse struct {
	ApplicationId    uint      `json:"application_id"`
	JobId            uint      `json:"job_id"`
	JobTitle         string    `json:"job_name"`
	UserId           uint      `json:"user_id"`
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	Status           int       `json:"status"`
	StatusName       string    `json:"status_name"`
	IsActive         int       `json:"is_active"`
	CvLink           string    `json:"cv_link"`
	CoverLetterLink  string    `json:"cover_letter_link"`
	YearOfExperience int       `json:"year_of_experience"`
	IssuedAt         time.Time `json:"issued_at"`
}

type GetApplicationRequest struct {
	JobId uint   `json:"job_id"`
	Email string `json:"email"`
}

func (req GetApplicationRequest) Validate() *general.Error {
	if req.JobId == 0 {
		return exception.NewError(http.StatusBadRequest, "JobId is required")
	}
	if req.Email == "" {
		return exception.NewError(http.StatusBadRequest, "Email is required")
	}

	return nil
}

func GetApplicationStatusName(status int) string {
	switch status {
	case constant.OnHoldStatusId:
		return constant.OnHoldStatusName
	case constant.InterviewStatusId:
		return constant.InterviewStatusName
	case constant.AcceptedStatusId:
		return constant.AcceptedStatusName
	case constant.RejectedStatusId:
		return constant.RejectedStatusName
	default:
		return ""
	}
}

func GetApplicationActiveBaseOnStatus(status int) int {
	if status == constant.AcceptedStatusId || status == constant.RejectedStatusId {
		return 0
	}
	return 1
}
