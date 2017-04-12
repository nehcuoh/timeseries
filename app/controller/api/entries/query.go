package entries

import (
	"fmt"

	"github.com/qiangxue/fasthttp-routing"
)

func Check(context *routing.Context) error {
	list_name := context.Param("list_name")
	member := context.Param("object_id")
	fmt.Fprintf(context, "check  member: %s in list %s,", member, list_name)
	return nil
}
