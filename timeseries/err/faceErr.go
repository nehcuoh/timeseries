package err

import (
	"zeus/timeseries/face"
)

var FaceError = map[string](*face.Error){
	"parse ip fail":               &face.NewError(1001, "parse ip fail"),
	"invalid uid":                 &face.NewError(1002, "invalid uid string"),
	"watcher receive a nil event": &face.NewError(1003, "nil event"),
}
