package api

import (
	"fastgo/module/util"
	"fmt"
	"reflect"

	"github.com/gosexy/to"
	"github.com/valyala/fasthttp"
)

type Controller struct {
	UserId string
}

type RetData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type InputData map[int]string

type DataNil struct {
}

func (c *Controller) GetData(inputData interface{}, ctx *fasthttp.RequestCtx) {
	if ctx.IsPost() {
		c.collectPostData(inputData, ctx)
	} else if ctx.IsGet() {
		c.collectGetData(inputData, ctx)
	}
}

func (c *Controller) collectGetData(inputData interface{}, ctx *fasthttp.RequestCtx) {
	// log.Info("request data: " + string(ctx.QueryArgs().String()) + ", api: " + string(ctx.Path()))

	s := reflect.ValueOf(inputData).Elem()
	typeOf := s.Type()

	for i := 0; i < s.NumField(); i++ {
		field := typeOf.Field(i)
		data := string(ctx.FormValue(field.Tag.Get("json")))
		switch fmt.Sprintf("%s", field.Type) {
		case "string":
			s.Field(i).SetString(to.String(reflect.ValueOf(data)))
		case "int":
			fallthrough
		case "int8":
			fallthrough
		case "int16":
			fallthrough
		case "int32":
			fallthrough
		case "int64":
			s.Field(i).SetInt(to.Int64(reflect.ValueOf(data)))
		case "uint":
			fallthrough
		case "uint8":
			fallthrough
		case "uint16":
			fallthrough
		case "uint32":
			fallthrough
		case "uint64":
			s.Field(i).SetUint(to.Uint64(reflect.ValueOf(data)))
		case "[]string":
			tmpSlice := make([]string, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		case "[]uint8":
			tmpSlice := make([]uint8, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		case "[]int":
			tmpSlice := make([]int, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		case "[]int32":
			tmpSlice := make([]int32, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		case "[]int64":
			tmpSlice := make([]int64, 0)
			util.JsonDecode(data, &tmpSlice)
			s.Field(i).Set(reflect.ValueOf(tmpSlice))
		}
	}
}

func (c *Controller) collectPostData(inputData interface{}, ctx *fasthttp.RequestCtx) {
	formData := string(ctx.PostBody())
	// log.Infof("request data: %s, api: %v", formData, string(ctx.Path()))

	util.JsonDecode(formData, &inputData)
	return
}

func (c *Controller) SendText(ctx *fasthttp.RequestCtx, retData string) {
	fmt.Fprintf(ctx, "%s", retData)
}

func (c *Controller) SendJson(ctx *fasthttp.RequestCtx, retData RetData) {
	ctx.SetContentType("application/json")
	retJson, _ := util.JsonEncode(retData)
	// log.Info("response data: " + retJson + ", api: " + string(ctx.Path()))
	fmt.Fprintf(ctx, "%s", retJson)
}
