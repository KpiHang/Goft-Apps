package controllers

import (
	"UseGoft/src/middleware"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

// Index 控制器方法
func (this *IndexController) Index(ctx *gin.Context) goft.Json {
	// goft.Throw("test error", 500, ctx) // 测试异常。
	return gin.H{"result": "index"}
}

// 若实现IClass接口，就认为是控制器；
func (this *IndexController) Name() string {
	return "IndexController"
}
func (this *IndexController) Build(goft *goft.Goft) { // 注册路由
	goft.Handle("GET", "/", this.Index).
		HandleWithFairing("GET", "/test", this.IndexTest, middleware.NewIndexTest()) // 注册单路径生效的中间件
}

func (this *IndexController) IndexTest(ctx *gin.Context) goft.Json {
	// goft.Throw("test error", 500, ctx) // 测试异常。
	return gin.H{"result": "index test"}
}
