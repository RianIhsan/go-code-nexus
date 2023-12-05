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

// @Summary Register a new user
// @Description Register a new user with the provided details
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.TRegisterRequest true "User registration details"
// @Success 200 {string} string "registration is successful, please check your email for email verification"
// @Failure 400 {string} string "invalid payload" or "error validating payload" or "registration failed"
// @Router /api/v1/auth/signup [post]
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
