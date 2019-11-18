package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zhe-ma/login-server-study/config"
	"github.com/zhe-ma/login-server-study/router"
)

var configFilePath = pflag.StringP("config_file", "c", "", "Config file path")

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		url := "http://localhost" + viper.GetString("port") + "/version"
		log.Println("Ping URL: ", url)

		rsp, err := http.Get(url)
		if err == nil && rsp.StatusCode == http.StatusOK {
			return nil
		}

		time.Sleep(time.Second)
	}

	return errors.New("Cannot connect to the server")
}

func main() {
	pflag.Parse()

	if err := config.Init(*configFilePath); err != nil {
		panic(err)
	}

	gin.SetMode(viper.GetString("run_mode"))
	log.Println("Running mode from config file", viper.GetString("run_mode"))

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

	engine.Run(viper.GetString("port"))
}
