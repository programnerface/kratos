package errors

import (
	//"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
)

// 辅助方法，通过这个方法来传一些东西
func NewHTTPError(code int, field string, detail string) *HTTPError {
	return &HTTPError{
		Code: code,
		Errors: map[string][]string{
			field: {detail},
		},
	}
}

type HTTPError struct {
	Errors map[string][]string `json:"errors"`
	Code   int                 `json:"-"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError: %d", e.Code)
}

// FromError try to convert an error to *HTTPError.
func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	//之前测试的 0719
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}
	//判定一下这个东西是kratos error
	/*这一步你在正常写kratos的业务逻辑的时候或者其他项目的时候是不需要的，这个东西的转换是因为我们要改他的格式，
	所以在这里面需要对kratos内部的错误， kratos里面定义的错误和这个realworld的错误格式转换，*/
	if se := new(errors.Error); errors.As(err, &se) {
		//return出去的就是我们自己定义的格式
		//然后我们直接从kratoserror里面把code拿出来，因为这个东西也是一个httpcode,直接用就可以了
		//然后把Reason拿出来，Message放第三个就好了
		//这样的话就可以把kratos的错误转换到这边(realworld)了
		return NewHTTPError(int(se.Code), se.Reason, se.Message)
	}
	//return &HTTPError{} -0719
	return NewHTTPError(500, "internal", "error")
}
