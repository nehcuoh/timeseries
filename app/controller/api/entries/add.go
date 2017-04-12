package entries

import (
	"fmt"

	"github.com/qiangxue/fasthttp-routing"
)

func Add(context *routing.Context) error {
	list_name := context.Param("list_name")
	member := context.Param("object_id")
	fmt.Fprintf(context, "add new entry, list %s, member: %s", list_name, member)
	return nil
}
