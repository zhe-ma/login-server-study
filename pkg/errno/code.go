package errno

var (
	// System errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error ouccured while binding the request body to the struct."}

	// Common errors
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrSignToken  = &Errno{Code: 20003, Message: "Failed to sign a Json Web Token."}

	// User errors
	ErrEncrypt         = &Errno{Code: 20101, Message: "Failed to encrypt the user password."}
	ErrUserNotFound    = &Errno{Code: 20102, Message: "The user was not found."}
	ErrInvalidPassword = &Errno{Code: 20103, Message: "Invalid password."}
	ErrTokenInvalid    = &Errno{Code: 20104, Message: "Invalid token."}
)
