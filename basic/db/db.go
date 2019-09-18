package db

import (
	"database/sql"
	"github.com/micro/go-micro/util/log"
	"gomicro/basic/conf"
	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

// Init 初始化数据库
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Log("[Init] db 已经初始化过")
		return
	}

	// 如果配置声明使用mysql
	if conf.GetMysqlConfig().GetEnabled() {
		initMysql()
	}

	inited = true
}

// GetDB 获取db
func GetDB() *sql.DB {
	return mysqlDB
}

func initMysql() {

	var err error

	// 创建连接
	mysqlDB, err = sql.Open("mysql", conf.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(conf.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(conf.GetMysqlConfig().GetMaxIdleConnection())

	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
