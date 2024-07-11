package models

// UserDetailRequest 用户请求发的数据的信息 gin 验证用户请求是否有效
type UserDetailRequest struct {
	UserId int `binding:"required,gt=0,lt=100000" uri:"uid"`
	//required:必填,gt:大于0,lt:小于100000, uri:"uid 从uri中提取uid
}

func NewUserDetailRequest() *UserDetailRequest {
	return &UserDetailRequest{}
}

// UserModel 用户信息
type UserModel struct {
	UserId   int    `gorm:"column:user_id" json:"user_id" xorm:"user_id"`
	UserName string `gorm:"column:user_name" json:"user_name" xorm:"user_name"`
	UserPwd  string `gorm:"column:user_pwd" json:"user_pwd" xorm:"user_pwd"`
}

func NewUserModel(userId int, userName string) *UserModel {
	return &UserModel{UserId: userId, UserName: userName}
}
