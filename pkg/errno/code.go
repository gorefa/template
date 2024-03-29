package errno

var (
	// Common errors
	OK                  = &Errno{Code: 200, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrHeaderNull        = &Errno{Code: 20105, Message: "The header was not found."}

	// k8s
	ErrClusterIsExists     = &Errno{Code: 30001, Message: "Cluster Is Exists."}
	ErrClusterNotSpecified = &Errno{Code: 30002, Message: "Please specify the cluster."}

	// alert
	ErrAlert = &Errno{Code: 40001, Message: "Alert message error."}
)
