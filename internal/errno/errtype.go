/**
* @file   : errtype
* @descrip: 自定义错误类型: 返回给用户的 Errno, 开发环境的 Err (代替系统标准库的 error)
* @author : ch-yk
* @create : 2018-09-03 下午5:53
* @email  : commonheart.yk@gmail.com
**/

package errno

import "fmt"

//Error 的代号以及信息 --- Errno
type Errno struct {
	Code    int
	Message string
}

func (err Errno) Error() string {
	return err.Message
}

// 包含异常本身的 Err 对象 --- Err
type Err struct {
	Code    int
	Message string
	Err     error
}

func New(errno *Errno, err error) *Err {
	return &Err{Code: errno.Code, Message: errno.Message, Err: err}
}

/* Err 成员方法 ----- begin */
func (err *Err) Add(message string) error {
	//err.Message = fmt.Sprintf("%s %s", err.Message, message)
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.Message = fmt.Sprintf("%s %s", err.Message, fmt.Sprintf(format, args...))
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

/* Err 成员方法 ----- end*/

/*实例应用*/
func IsServiceNotFoundErr(err error) bool {
	code, _ := DecodeErr(err)
	return code == ServiceNotFoundErr.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerErr.Code, err.Error()
}
