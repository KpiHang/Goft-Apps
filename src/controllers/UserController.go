package controllers

import (
	"UseGoft/src/service"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type UserController struct {
	// Db   *gorm.DB      `inject:"-"` // 注入gorm 这是什么意思？？？
	//Db   *xorm.Engine  `inject:"-"`
	// Db *configuration.XOrmAdapter `inject:"-"` // 单例多例？？？
	// user *daos.UserDAO              `inject:"-"` // user 信息相关的Dao层
	UserService *service.UserService `inject:"-"`
}

func NewUserController() *UserController {
	return &UserController{}
}

// 控制器方法
//func (this *UserController) UserDetail(ctx *gin.Context) goft.Json {
//	//req := models.NewUserDetailRequest()
//	//goft.Error(ctx.BindUri(req))
//	req, _ := ctx.Get("_req")                     // 在用户中间件中设置的。
//	uid := req.(*models.UserDetailRequest).UserId // 中间件验证过的一定不会出错的userId
//	user := &models.UserModel{}
//	goft.Error(this.Db.Table("users").
//		Where("user_id = ?", uid).
//		Find(user).Error)
//	return user
//}

// UserDetail 使用简单的sql查询
func (this *UserController) UserDetail(ctx *gin.Context) goft.Query { // goft.Json
	//user := &models.UserModel{}
	//_, err := this.Db.Table("users").Where("user_id = ?", ctx.Param("uid")).
	//	Get(user)
	//goft.Error(err)
	//return user

	//return goft.SimpleQuery(`
	//						select
	//							user_id, user_name
	//						from
	//						    users
	//						where user_id = ?`).
	//	WithArgs(ctx.Param("uid")).WithFirst()

	//return this.user.GetUserDetail(ctx.Param("uid"))

	// Dao + Service 写法
	return this.UserService.GetUserDetail(ctx.Param("uid"))
}

// UserAccessToken 获取登陆token /access_token?uname=xxx&upwd=xxx
func (this *UserController) UserAccessToken(ctx *gin.Context) goft.Json { // goft.Json
	if uname, upwd := ctx.Query("uname"), ctx.Query("upwd"); uname != "" && upwd != "" {
		return gin.H{"token": this.UserService.UserLogin(uname, upwd)}
	}
	panic("error user access params")
}

func (this *UserController) UserList(ctx *gin.Context) goft.Json { // goft.SimpleQuery
	ret := goft.SimpleQuery("select * from users").WithKey("result").Get() // 自定义响应的json map结构，eg {result : []}
	return ret
}

func (this *UserController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/users", this.UserList)
	goft.Handle("GET", "/users/:uid", this.UserDetail)
	goft.Handle("GET", "/access_token", this.UserAccessToken)
	//goft.HandleWithFairing("GET", "/users/:uid", this.UserDetail,
	//	middleware.NewUserMiddleware()) // 注册路由中间件
}

func (this *UserController) Name() string {
	return "UserController"
}

func (this *UserController) Init() error {
	return nil
}
