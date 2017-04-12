package api

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
)

func List(ctx *routing.Context) error {
	fmt.Fprintf(ctx, "create new list %s",ctx.Param("list_name"))
	return nil
}

