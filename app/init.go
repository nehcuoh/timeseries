package app

import (
	"fmt"

	"github.com/qiangxue/fasthttp-routing"
	"zeus/app/controller"
	"zeus/app/controller/api"
	"zeus/app/controller/api/entries"
)

//全局路由
var Route *routing.Router

func init() {
	fmt.Println("init controller !")
	Route = routing.New()
	Route.Get("/", controller.Index)

	entry_curd_uri := "/<list_name>/entries/<object_id>"

	v1apiRouter := Route.Group("/list")
	v1apiRouter.Get("/", api.List)

	v1apiRouter.Get(entry_curd_uri, entries.Check)
	v1apiRouter.Post(entry_curd_uri, entries.Add)
	v1apiRouter.Delete(entry_curd_uri, entries.Del)

	v1apiRouter.Post("/<list_name>", api.List)
	v1apiRouter.Get("/<list_name>/status", api.Status)
}
