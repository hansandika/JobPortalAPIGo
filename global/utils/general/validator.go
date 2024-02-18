package utils

import (
	"github.com/golodash/galidator"
)

var (
	// Validator is a global validator
	GlobalValidator = galidator.New()
)

var UserRegisterValidator = GlobalValidator.ComplexValidator(galidator.Rules{
	"Email":    GlobalValidator.R("Email").Required().Min(5).Max(100).Email(),
	"Name":     GlobalValidator.R("Name").Required().Min(3).Max(100),
	"Password": GlobalValidator.R("Password").Required().Min(5).Max(100).Password(),
	"RoleId":   GlobalValidator.R("RoleId").Required().Min(1).Max(2),
})

var UserLoginValidator = GlobalValidator.ComplexValidator(galidator.Rules{
	"Email":    GlobalValidator.R("Email").Required().Email(),
	"Password": GlobalValidator.R("Password").Required(),
})

var JobAddValidator = GlobalValidator.ComplexValidator(galidator.Rules{
	"JobTitle":    GlobalValidator.R("JobTitle").Required().Min(3).Max(100),
	"Description": GlobalValidator.R("Description").Required().Min(3).Max(1000),
	"Requirement": GlobalValidator.R("Requirement").Required(),
})

var ApplicationAddValidator = GlobalValidator.ComplexValidator(galidator.Rules{
	"JobId":            GlobalValidator.R("JobId").Required(),
	"Email":            GlobalValidator.R("Email").Required(),
	"Status":           GlobalValidator.R("Status").Required().Min(1).Max(4),
	"CvLink":           GlobalValidator.R("CvLink").Required().Min(5).Max(255),
	"CoverLetterLink":  GlobalValidator.R("CoverLetterLink").Required().Min(5).Max(255),
	"YearOfExperience": GlobalValidator.R("YearOfExperience").Required().Min(1).Max(100),
})

var ApplicationUpdateValidator = GlobalValidator.ComplexValidator(galidator.Rules{
	"JobId":  GlobalValidator.R("JobId").Required(),
	"Email":  GlobalValidator.R("Email").Required(),
	"Status": GlobalValidator.R("Status").Required().Min(1).Max(4),
})
