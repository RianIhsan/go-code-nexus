package handler

import (
	"github.com/RianIhsan/go-code-nexus/domain/users"
	"github.com/RianIhsan/go-code-nexus/domain/users/dto"
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/RianIhsan/go-code-nexus/utils/response"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService users.IServiceUser
}

func NewUserHandler(userService users.IServiceUser) users.IHandlerUsers {
	return &UserHandler{
		userService,
	}
}

// @Summary Get current user
// @Description Get information about the currently authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} dto.TGetUserResponse "User information"
// @Failure 401 {string} string "user not found" or "Access denied"
// @Router /api/v1/user/me [get]
// @Security Bearer
func (h *UserHandler) GetCurrentUser(c *fiber.Ctx) error {
	user, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || user == nil {
		return response.SendStatusUnauthorized(c, "user not found")
	}

	return response.GetCurrentUser(c, dto.GetUserResponse(user))
}
