package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhe-ma/login-server-study/router"
)

func pingServer() error {
	for i := 0; i < 2; i++ {
		rsp, err := http.Get("http://localhost:8080/version")
		if err == nil && rsp.StatusCode == 200 {
			return nil
		}

		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the server")
}

func main() {
	engine := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(engine, middlewares...)

	go func() {
		err := pingServer()
		if err != nil {
			log.Fatal("Failed to ping server. The server has not response. Error: ", err)
			return
		}
		log.Println("Ping server successfully!")
	}()

	engine.Run()
}
