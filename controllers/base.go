package controllers

import (
	"spos/auth/usecase"

	"github.com/labstack/echo"
)

type controller struct {
	usecase usecase.Usecase
}

type Controller interface {
	// LoginHandler handler for user login
	LoginHandler(e echo.Context) error

	// RegisterHandler handler for user register
	RegisterHandler(e echo.Context) error

	// VerificationRegisterHandler handler for verification register
	// using OTP
	VerificationRegisterHandler(e echo.Context) error
}

func New(uc usecase.Usecase) Controller {
	return &controller{
		usecase: uc,
	}
}
