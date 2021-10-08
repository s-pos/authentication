package controllers

import (
	"net/http"

	"spos/auth/constant"
	"spos/auth/models"

	"github.com/labstack/echo"
	"github.com/s-pos/go-utils/utils/request"
	"github.com/s-pos/go-utils/utils/response"
)

func (c *controller) RegisterHandler(e echo.Context) error {
	var (
		req     = e.Request()
		ctx     = req.Context()
		payload models.RequestRegister
	)

	if err := request.BodyValidation(ctx, e, &payload, request.JSON); err != nil {
		return response.Errors(ctx, http.StatusBadRequest, string(constant.BodyRequired), constant.Message[constant.RegisterFailed], constant.Reason[constant.BodyRequired], err).Write(e)
	}

	return c.usecase.Register(ctx, payload).Write(e)
}
