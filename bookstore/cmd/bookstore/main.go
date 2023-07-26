package main

import (
	_ "bookstore/internal/store" //自动注册
	"bookstore/server"
	"bookstore/store/factory"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	s, err := factory.New("redis")
	if err != nil {
		panic(err)
	}

	srv := server.NewBookStoreServer(":9999", s)

	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln("web server start failed:", err)
		return
	}

	log.Println("web server start success")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errChan:
		log.Println("web server run failed:", err)
		return
	case <-c:
		log.Println("bookstore server exit...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("web server shutdown failed:", err)
		return
	}
	log.Println("web server shutdown success")
}
