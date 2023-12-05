package dto

// TRegisterRequest represents the request payload for user registration
type TRegisterRequest struct {
	// Name of the user (required)
	// Example: John Doe
	// Required: true
	Name string `json:"name" validate:"required"`

	// Email of the user (required, should be a valid email address)
	// Example: john.doe@example.com
	// Required: true
	// Format: email
	Email string `json:"email" validate:"required,email"`

	// Password of the user (required, minimum length 8)
	// Example: mySecurePassword
	// Required: true
	// MinLength: 8
	Password string `json:"password" validate:"required,min=8"`
}
