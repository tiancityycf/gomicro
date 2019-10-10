package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	//ctxHttp()
	ctxWithCancal()
	//ctxWithTimeout()
}

func ctxHttp() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// monitor
		go func() {
			for range time.Tick(time.Second) {
				fmt.Println("req is processing")
			}
		}()

		// assume req processing takes 3s
		time.Sleep(3 * time.Second)
		w.Write([]byte("hello"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func ctxWithCancal() {
	ctx, cancel := context.WithCancel(context.Background())

	// monitor
	go func() {
		for range time.Tick(time.Second) {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				fmt.Println("ctx Done")
				return
			default:
				fmt.Println("monitor woring", ctx)
			}
		}
	}()

	time.Sleep(3 * time.Second)

	cancel()

	fmt.Println("call cancel Done")

	time.Sleep(3 * time.Second)

}

func ctxWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// monitor
	go func() {
		for range time.Tick(time.Second) {
			select {
			case <-ctx.Done():
				fmt.Println(ctx.Err())
				fmt.Println("ctx Done")
				return
			default:
				fmt.Println("monitor woring", ctx)
			}
		}
	}()
	time.Sleep(4 * time.Second)

	fmt.Println("end", ctx)

}
