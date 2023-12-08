package dto

import "github.com/RianIhsan/go-code-nexus/entities"

type TGetUserResponse struct {
	id         int    `json:"id"`
	Name       string `json:"username"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	IsVerified bool   `json:"is_verified"`
}

func GetUserResponse(user *entities.UserEntity) *TGetUserResponse {
	userFormatter := &TGetUserResponse{}
	userFormatter.id = user.ID
	userFormatter.Name = user.Name
	userFormatter.Avatar = user.Avatar
	userFormatter.Email = user.Email
	userFormatter.Role = user.Role
	userFormatter.IsVerified = user.IsVerified

	return userFormatter
}
