package basic

import (
	"gomicro/basic/conf"
	"gomicro/basic/db"
	"gomicro/basic/redis"
)

func Init() {
	conf.Init()
	db.Init()
	redis.Init()
}
