package configuration

import (
	"UseGoft/src/daos"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"xorm.io/xorm"
)

type DBConfig struct{}

func NewDBConfig() *DBConfig {
	return &DBConfig{}
}

func (this *DBConfig) GormDB() *gorm.DB {
	db, err := gorm.Open("mysql",
		"root:rootroot@tcp(localhost:3307)/gotest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)
	return db
}

func (this *DBConfig) XOrm() *daos.XOrmAdapter {
	engine, err := xorm.NewEngine("mysql", "root:rootroot@tcp(localhost:3307)/gotest?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	engine.DB().SetMaxIdleConns(5)
	engine.DB().SetMaxOpenConns(10)
	return &daos.XOrmAdapter{engine}
}
