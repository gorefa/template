/**
* @file   : errcode
* @descrip: 错误类型对象的标识，即全局变量
* @author : ch-yk
* @create : 2018-09-03 下午5:52
* @email  : commonheart.yk@gmail.com
**/

package errno


var (
	OK                  = &Errno{Code: 0, Message: "OK"}

	/************ Common errors ***********/
	InternalServerErr   = &Errno{Code: 10001, Message: "Internal server error."}
	BindErr             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct.(一般是没传参数或者传入的参数错误)"}

	// Service errors
	ServiceNotFoundErr = &Errno{Code: 10003, Message: "The service was not found."}


	/*下面是外部类型错误，非 common errors*/

	//User errors
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
)