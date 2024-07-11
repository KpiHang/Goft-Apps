package service

import (
	"UseGoft/src/daos"
	"github.com/shenyisyn/goft-gin/goft"
	"strconv"
)

type UserService struct {
	UserDao *daos.UserDAO `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}
func (this *UserService) GetUserDetail(param string) goft.Query {
	if uid, err := strconv.Atoi(param); err == nil {
		return this.UserDao.GetUserById(uid)
	} else {
		return this.UserDao.GetUserByName(param)
	}
}

// jwt todo
func (this *UserService) UserLogin(uname string, upwd string) string {
	if this.UserDao.FindByUserName(uname).UserPwd == upwd {
		return "token " + uname
	}
	panic("error user access")
}
