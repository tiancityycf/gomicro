package main

import (
	"encoding/json"
	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	proto "gomicro/services/user/proto"
	"log"

	"context"
)

type User struct {
	Client proto.UserService
}

func (s *User) GetProfileById(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received User API request")

	id, ok := req.Get["id"]
	if !ok || len(id.Values) == 0 {
		return errors.BadRequest("gomicro.user", "id cannot be blank")
	}

	response, err := s.Client.GetProfileById(ctx, &proto.UserRequest{
		//Id: strings.Join(id.Values, " "),
		Id: 1,
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.User.Name,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("gomicro.api.user"),
	)

	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: proto.NewUserService("user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
