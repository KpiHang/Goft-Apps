package middleware

import (
	"UseGoft/src/models"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

// 用户id校验中间件
type UserMiddleware struct{}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (this UserMiddleware) OnRequest(ctx *gin.Context) error {
	req := &models.UserDetailRequest{}
	goft.Error(ctx.BindUri(req))
	ctx.Set("_req", req)
	return nil
}

func (this *UserMiddleware) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}
