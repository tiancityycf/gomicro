package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/web"
	proto "gomicro/services/user/proto"

	"context"
)

func main() {
	service := web.NewService(
		web.Name("gomicro.web.user"),
	)

	service.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			id, _ := strconv.Atoi(r.Form.Get("id"))
			if id == 0 {
				w.Write([]byte(`<html><body><h1>参数错误</h1></body></html>`))
				return
			}

			cl := proto.NewUserService("user", client.DefaultClient)
			rsp, err := cl.GetProfileById(context.Background(), &proto.UserRequest{
				Id: int32(id),
			})

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			if rsp.User == nil {
				http.Error(w, "user not exists", 404)
				return
			}

			w.Write([]byte(`<html><body><h1>` + rsp.User.Name + `</h1></body></html>`))
			return
		}

		fmt.Fprint(w, `<html><body><h1>Enter Name<h1><form method=post><input name=id type=text /></form></body></html>`)
	})

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
