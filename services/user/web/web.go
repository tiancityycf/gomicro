package main

import (
	"fmt"
	"log"
	"net/http"

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

			name := r.Form.Get("name")
			if len(name) == 0 {
				name = "World"
			}

			cl := proto.NewUserService("user", client.DefaultClient)
			rsp, err := cl.Info(context.Background(), &proto.UserRequest{
				Name: name,
			})

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			w.Write([]byte(`<html><body><h1>` + rsp.Code + `</h1></body></html>`))
			return
		}

		fmt.Fprint(w, `<html><body><h1>Enter Name<h1><form method=post><input name=name type=text /></form></body></html>`)
	})

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
