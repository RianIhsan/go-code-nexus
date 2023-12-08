package dto

import "github.com/RianIhsan/go-code-nexus/entities"

type TLoginResponse struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"access_token"`
}

func LoginResponse(user *entities.UserEntity, token string) *TLoginResponse {
	userFormatter := &TLoginResponse{}
	userFormatter.Email = user.Email
	userFormatter.Role = user.Role
	userFormatter.Token = token

	return userFormatter
}
