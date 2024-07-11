package daos

import (
	"UseGoft/src/models"
	"github.com/shenyisyn/goft-gin/goft"
)

type UserDAO struct {
	Db *XOrmAdapter `inject:"-"`
	// Db *gorm.DB `inject:"-"`
}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// FindByUserName 通过ORM访问数据库  凡是findBy前缀的都是以ORM方式访问数据库
func (this *UserDAO) FindByUserName(username string) *models.UserModel {
	userModel := &models.UserModel{}
	has, err := this.Db.Table("users").Where("user_name=?", username).Get(userModel)
	if err != nil || !has {
		panic("user not found")
	}
	//goft.Error(this.Db.Table("users").Where("user_name=?", username).Find(userModel).Error)
	return userModel
}

// GetUserById 通过SQL语句访问数据库 凡是Get前缀的都是以SQL方式访问数据库
func (this *UserDAO) GetUserById(uid int) goft.Query {
	return goft.SimpleQuery("select * from users where user_id = ?").
		WithArgs(uid).WithFirst()
}

func (this *UserDAO) GetUserByName(uname string) goft.Query {
	return goft.SimpleQuery("select user_id, user_name from users where user_name = ?").
		WithArgs(uname).WithFirst()
}

func (this *UserDAO) GetUserDetail(uid string) goft.Query {
	return goft.SimpleQuery("select user_id, user_name from users where user_id = ?").
		WithArgs(uid).WithFirst()
}
