package handler

import (
	"github.com/RianIhsan/go-code-nexus/module/auth"
	"github.com/RianIhsan/go-code-nexus/module/auth/dto"
	"github.com/RianIhsan/go-code-nexus/utils/response"
	"github.com/RianIhsan/go-code-nexus/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService auth.IServiceAuth
}

func NewAuthHandler(authService auth.IServiceAuth) auth.IHandlerAuth {
	return &AuthHandler{
		authService,
	}
}

func (h AuthHandler) SignUp(c *fiber.Ctx) error {
	var payload dto.TRegisterRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}
	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	_, err := h.authService.SignUp(&payload)
	if err != nil {
		return response.SendStatusBadRequest(c, err.Error())
	}
	return response.SendStatusOkResponse(c, "registration is successful, please check your email for email verification")
}
