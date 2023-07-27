package main

import (
	"add/api"
	"add/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	fmt.Println("run......",config.Load().HTTPPort)

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	api.NewApi(r)
		err := r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
