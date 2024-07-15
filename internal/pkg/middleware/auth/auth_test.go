package auth

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	tk := GenerateToken("face")
	spew.Dump(tk)
	panic("tk")
}

/*
明明都对，postman测试也通过了但是终端返回的结果却是
(interface {}) <nil>
打印spew.Dump(claims)后发现
在测试了之后发现，要在pkg/middleware/auth的位置重新go test，改变的代码才会生效。
*/
