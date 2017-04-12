package entries

import (
	"fmt"

	"github.com/qiangxue/fasthttp-routing"
)

func Del(context *routing.Context) error {
	list_name := context.Param("list_name")
	member := context.Param("object_id")
	fmt.Fprintf(context, "delete , list %s, member: %s", list_name, member)
	return nil
}
