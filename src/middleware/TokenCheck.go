package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

// TokenCheck 实现了OnRequest中间件 （执行控制器方法前，修改如头信息、判断参数等等。）
// 实现了中间件接口。
type TokenCheck struct{}

func NewTokenCheck() *TokenCheck {
	return &TokenCheck{}
}

func (this *TokenCheck) OnRequest(ctx *gin.Context) error {
	if ctx.Query("token") == "" {
		goft.Throw("token requered", 503, ctx)
	}
	return nil
}
func (this *TokenCheck) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}
