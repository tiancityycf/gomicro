package basic

import (
	"gomicro/basic/conf"
	"gomicro/basic/db"
)

func Init() {
	conf.Init()
	db.Init()
}
