package dto

// TCreateUserDetailRequest represents the request payload for creating user details
type TCreateUserDetailRequest struct {
	// Address of the user (optional)
	// Example: 123 Main Street
	Address *string `json:"address"`

	// Phone number of the user (optional)
	// Example: +1234567890
	Phone *string `json:"phone"`

	// Job of the user (optional)
	// Example: Software Engineer
	Job *string `json:"job"`
}

type TUpdateAvatarRequest struct {
	Avatar string `form:"avatar"`
}
