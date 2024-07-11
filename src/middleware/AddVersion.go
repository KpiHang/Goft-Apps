package middleware

import (
	"github.com/gin-gonic/gin"
)

// AddVersion 实现了OnResponse中间件 （执行控制器方法后，可以修改返回值内容。）
// 对控制方法返回的数据进行操作。
type AddVersion struct{}

func NewAddVersion() *AddVersion {
	return &AddVersion{}
}

func (this *AddVersion) OnRequest(ctx *gin.Context) error {
	return nil
}
func (this *AddVersion) OnResponse(result interface{}) (interface{}, error) {
	if m, ok := result.(gin.H); ok {
		m["version"] = "1.0.0"
	}
	return result, nil
}
