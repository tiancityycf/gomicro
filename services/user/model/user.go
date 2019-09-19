package model

import (
	"context"
	"database/sql"
	"github.com/micro/go-micro/util/log"
	"gomicro/basic/db"
	pb "gomicro/services/user/proto"
)

type User struct{}

func (u *User) GetProfileById(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	if req.GetId() == 0 {
		rsp.Error = &pb.Error{Code: 404, Message: "user not exits"}
		return nil
	}

	queryString := `SELECT id, username as name  FROM platv4_user WHERE id = ?`

	// 获取数据库
	o := db.GetDB()

	profile := &pb.User{}

	// 查询
	row := o.QueryRow(queryString, req.GetId())
	if err := row.Scan(&profile.Id, &profile.Name); err == nil {
		rsp.User = profile
	} else {
		if err != sql.ErrNoRows {
			log.Logf("[GetProfileById] 查询数据失败，err：%v", err)
		}
	}

	return nil
}
