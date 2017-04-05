package face

import (
	"fmt"
	"reflect"

	series "zeus/timeseries"
)

type Uid struct {
	uid uint64
}

func UID(u uint64) (*Uid) {
	return &Uid{u}
}

func (u*Uid) Key() (interface{}) {
	return u.uid
}
func (u*Uid) ParseKey(key interface{}) (t series.FaceKey, err error) {
	k, e := key.(uint64)
	if !e {
		info := fmt.Sprintf("Parse Uid fail,key: %d", key)
		fmt.Println(reflect.TypeOf(key))
		return nil, &Error{info}
	}
	uid := uint64(k)
	return &Uid{uid: uid}, nil
}

func (u Uid) String() string {
	return fmt.Sprintf("%d", u.uid)
}
