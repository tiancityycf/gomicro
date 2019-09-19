package model

import (
	"database/sql"
	"github.com/micro/go-micro/util/log"
	"gomicro/basic/db"
	pb "gomicro/services/user/proto"
)

type User struct{}

func (u *User) GetProfileById(id int32) (*pb.User, error) {

	var profile pb.User
	queryString := `SELECT id, username as name  FROM platv4_user WHERE id = ?`
	// 获取数据库
	o := db.GetDB()
	// 查询
	row := o.QueryRow(queryString, id)
	if err := row.Scan(&profile.Id, &profile.Name); err == nil {
		return &profile, nil
	} else {
		if err != sql.ErrNoRows {
			log.Logf("[GetProfileById] 查询数据失败，err：%v", err)
			return nil, nil
		}
	}
	return nil, nil

}
