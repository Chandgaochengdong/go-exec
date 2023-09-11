package mysql

import (
	"com.shopline.richard/exec/dal/dao/query"
	"gopkg.inshopline.com/commons/bizerror/try"
	"gopkg.inshopline.com/commons/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
)

var (
	Db *gorm.DB
)

func Init() {
	// 连接数据库
	dsn, found := config.Find("rdb.default.dsn")
	if !found {
		panic("missing RDB DSN config")
	}
	Db = try.Must(gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		PrepareStmt: true,
	}))
	sqlDB, err := Db.DB()
	if err == nil {
		conns, _ := config.Find("rdb.default.max-open-conns")
		cons, _ := strconv.Atoi(conns)
		sqlDB.SetMaxIdleConns(cons)
	}
	query.SetDefault(Db)
	//gen.SetDefault(db)
}
