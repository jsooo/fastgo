package v1

import (
	"fastgo/api"

	"github.com/valyala/fasthttp"
)

type Controller struct {
	api.Controller
}

type HelloWorldInputData struct {
	Say string `json:"say"`
}

func (c *Controller) HelloWorld(ctx *fasthttp.RequestCtx) {
	var helloWorldData HelloWorldInputData
	c.GetData(&helloWorldData, ctx)

	retStr := helloWorldData.Say + "，welcome to fastgo :)"
	c.SendText(ctx, retStr)
}
