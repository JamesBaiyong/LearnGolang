package errno

var (
	//Common Error
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	//User Error
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user not found."}
)
