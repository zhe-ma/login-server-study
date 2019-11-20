package errno

const (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error ouccured while binding the request body to the struct."}

	// User errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
)
