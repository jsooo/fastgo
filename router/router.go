package router

import (
	v1_hello "fastgo/api/v1"

	"github.com/buaazp/fasthttprouter"
)

func InitRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	//eg. test api
	v1 := v1_hello.Controller{}
	router.POST("/v1/hello_world", v1.HelloWorld)
	router.GET("/v1/hello_world", v1.HelloWorld)

	return router
}
