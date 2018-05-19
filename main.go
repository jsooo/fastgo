package main

import (
	_ "fastgo/model"
	"fastgo/router"
	"log"

	"github.com/astaxie/beego/config"

	"github.com/valyala/fasthttp"
)

func main() {
	//You can set port by yourself
	iniconf, _ := config.NewConfig("ini", "config/app.conf")
	port := iniconf.String("port")

	log.Println("Fastgo Running On Port: " + port + " .....")

	//Initialize router
	fastRouter := router.InitRouter()

	log.Fatal(fasthttp.ListenAndServe(":"+port, fastRouter.Handler))
}
