package controller

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
)

func Index(ctx *routing.Context) error {
	fmt.Fprintf(ctx, "index controller")
	return nil
}

