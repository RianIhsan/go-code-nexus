package handler

import (
	"github.com/RianIhsan/go-code-nexus/domain/users"
	"github.com/RianIhsan/go-code-nexus/domain/users/dto"
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/RianIhsan/go-code-nexus/utils/cloudinary"
	"github.com/RianIhsan/go-code-nexus/utils/response"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
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

// @Summary Create user detail
// @Description Create user detail for the currently authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Param request body dto.TCreateUserDetailRequest true "User detail information"
// @Success 200 {string} string "User detail created successfully"
// @Failure 400 {string} string "Invalid input format"
// @Failure 401 {string} string "User not found" or "Access denied"
// @Router /api/v1/user/me/detail [post]
// @Security Bearer
func (h *UserHandler) CreateUserDetail(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || currentUser == nil {
		return response.SendStatusUnauthorized(c, "User not found")
	}

	var request dto.TCreateUserDetailRequest
	if err := c.BodyParser(&request); err != nil {
		return response.SendStatusBadRequest(c, "Invalid input format")
	}

	_, err := h.userService.CreateUserDetail(currentUser.ID, &request)
	if err != nil {
		return response.SendStatusInternalServerError(c, "Failed to create user detail: "+err.Error())
	}

	return response.SendStatusOkResponse(c, "User detail created successfully")
}

// @Summary Update user avatar
// @Description Update the avatar of the currently authenticated user
// @Tags User
// @Accept mpfd
// @Produce json
// @Security Bearer
// @Param avatar formData file true "Image file to upload"
// @Success 200 {string} string "Avatar update success"
// @Failure 400 {string} string "Invalid request or error uploading image"
// @Failure 401 {string} string "User not found or unauthorized"
// @Failure 500 {string} string "Internal server error"
// @Router /api/v1/user/me/avatar [patch]
func (h *UserHandler) UpdateAvatar(c *fiber.Ctx) error {
	currentUser, ok := c.Locals("CurrentUser").(*entities.UserEntity)
	if !ok || currentUser == nil {
		return response.SendStatusUnauthorized(c, "User not found")
	}

	file, err := c.FormFile("avatar")
	var uploadedURL string
	if err == nil {
		fileToUpload, err := file.Open()
		if err != nil {
			return response.SendStatusInternalServerError(c, "Failed open to open file: "+err.Error())
		}
		defer func(fileToUpload multipart.File) {
			_ = fileToUpload.Close()
		}(fileToUpload)
		uploadedURL, err = cloudinary.Uploader(fileToUpload)
		if err != nil {
			return response.SendStatusInternalServerError(c, "Failed to upload image: "+err.Error())
		}
	}

	userUpdateAvatar := &dto.TUpdateAvatarRequest{
		Avatar: uploadedURL,
	}

	image, err := h.userService.UpdateAvatar(currentUser.ID, userUpdateAvatar)
	if err != nil {
		return response.SendStatusBadRequest(c, "Error upload image: "+err.Error())
	}

	return response.SendStatusOkWithDataResponse(c, "Success updating avatar", dto.UpdateAvatarResponse(image))
}
