package api

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
)

func Status(ctx *routing.Context) error {
	fmt.Fprintf(ctx, "query %s status",ctx.Param("list_name"))
	return nil
}
