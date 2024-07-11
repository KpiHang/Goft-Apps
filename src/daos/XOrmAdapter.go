package daos

import (
	"database/sql"
	"xorm.io/xorm"
)

// XOrmAdapter 下面代码是使用xorm框架操作数据库，和gorm同类产品
type XOrmAdapter struct {
	*xorm.Engine
}

func (this *XOrmAdapter) DB() *sql.DB {
	return this.Engine.DB().DB
}
