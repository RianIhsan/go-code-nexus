package handler

import (
	"github.com/RianIhsan/go-code-nexus/domain/auth"
	"github.com/RianIhsan/go-code-nexus/domain/auth/dto"
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

// @Summary Verify user email
// @Description Verify user email with the provided token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.TVerificationRequest true "Email verification details"
// @Success 200 {string} string "Your email has been successfully verified"
// @Failure 400 {string} string "invalid payload" or "error validating payload" or "email verification failed"
// @Router /api/v1/auth/verify [post]
func (h AuthHandler) Verification(c *fiber.Ctx) error {
	var payload dto.TVerificationRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}
	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}
	err := h.authService.Verification(&payload)
	if err != nil {
		return response.SendStatusBadRequest(c, err.Error())
	}
	return response.SendStatusOkResponse(c, "Your email has been successfully verified")
}

// @Summary Sign in a user
// @Description Sign in a user with the provided credentials
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.TLoginRequest true "User login details"
// @Success 200 {object} dto.TLoginResponse "Signin successfully"
// @Failure 400 {string} string "invalid payload" or "error validating payload"
// @Failure 401 {string} string "user not found" or "your account has not been verified" or "incorrect password"
// @Router /api/v1/auth/signin [post]
func (h AuthHandler) SignIn(c *fiber.Ctx) error {
	var payload dto.TLoginRequest
	if err := c.BodyParser(&payload); err != nil {
		return response.SendStatusBadRequest(c, "invalid payload:"+err.Error())
	}
	if err := validator.ValidateStruct(payload); err != nil {
		return response.SendStatusBadRequest(c, "error validating payload:"+err.Error())
	}

	userLogin, accessToken, err := h.authService.SignIn(&payload)
	if err != nil {
		if err.Error() == "user not found" {
			return response.SendStatusNotFound(c, "user not found")
		} else if err.Error() == "your account has not been verified" {
			return response.SendStatusUnauthorized(c, "your account has not been verified")
		}

		return response.SendStatusUnauthorized(c, "incorrect password")
	}

	return response.SendStatusOkWithDataResponse(c, "signin successfully", dto.LoginResponse(userLogin, accessToken))

}
