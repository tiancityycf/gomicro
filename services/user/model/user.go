package model

import (
	"context"
	pb "gomicro/services/user/proto"
)

type User struct{}

func (u *User) GetProfileById(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	if req.GetId() == 0 {
		rsp.Error = &pb.Error{Code: 404, Message: "user not exits"}
		return nil
	}
	profile := pb.User{Id: req.GetId(), Name: "hello" + req.GetName()}
	rsp.User = &profile
	return nil
}
