package handler

import (
	"context"
	"gomicro/services/user/model"
	pb "gomicro/services/user/proto"
)

var user model.User

type Service struct{}

func (u *Service) GetProfileById(ctx context.Context, req *pb.UserRequest, rsp *pb.UserResponse) error {
	if req.GetId() == 0 {
		rsp.Error = &pb.Error{Code: 404, Message: "user not exits"}
		return nil
	}

	rsp.User, _ = user.GetProfileById(req.GetId())

	return nil
}
