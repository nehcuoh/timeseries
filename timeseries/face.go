package timeseries

import (
	"fmt"
	facePkg "zeus/timeseries/face"
)

type FaceKey interface {
	Key() interface{}
	ParseKey(interface{}) (FaceKey, error)
}

type Face struct {
	*FaceConfig
	m            map[interface{}]TimeLine
	name         string
	notification chan *Event
}

type FaceConfig struct {
	observer   []Notifier
	Resolution TimeInterval
}

func NewFace(name string, config *FaceConfig) (*Face) {
	return &Face{
		m:          make(map[interface{}]TimeLine),
		name:       name,
		FaceConfig: config,
	}
}

func (f*Face) Register(key FaceKey, line TimeLine) {
	f.m[key.Key()] = line
}

func (f*Face) Get(key FaceKey) (t TimeLine, err *Error) {
	e, ok := f.m[key.Key()]
	if !ok {
		return nil, &Error{error: fmt.Sprintf("face key : %v not exist", key.Key())}
	}
	return e, nil
}

func (f*Face) Remove(key FaceKey) {
	delete(f.m, key.Key())
	f.notification <- &Event{key: &key, point: nil, Type: EVENT_FACEKEY_DEL}
}

func (f*Face) Append(key FaceKey, point IPoint) {
	event := EventType(0)
	tl, _ := f.Get(key)
	if tl == nil {
		line := NewSeries(f.Resolution)
		f.Register(key, *line)
		tl = *line
		event |= EVENT_FACEKEY_NEW
	}
	appendEvent := tl.Append(point)
	event |= appendEvent
	f.notification <- &Event{key: &key, point: &point, Type: event}
}

func (f*FaceConfig) AddNotifier(n Notifier) {
	f.observer = append(f.observer, n)
}

func (f*FaceConfig) RemoveNotifier() {

}

//read from f.notification and call observer
func (f*Face) initWatcher() {
	watcher := func(face *Face) {
		for ; ; {
			e := <-face.notification
			if e == nil {
				err := facePkg.NewFatalError(1, "s")
				panic(err)
			}
			for _, observer := range (*face).observer {
				observer.Send(e)
			}
		}
	}
	go watcher(f)
}

func (f*Face) Dump(keyType FaceKey) {
	for i, l := range f.m {
		key, err := keyType.ParseKey(i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("face[%s] =\n%s\n", key, l)
	}
}
