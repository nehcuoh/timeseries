package timeseries

import "fmt"

type FaceKey interface {
	Key() interface{}
	ParseKey(interface{}) (FaceKey, error)
}

type Face struct {
	*FaceConfig
	m    map[interface{}]TimeLine
	name string
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
	f.sendNotification(&key, nil, EVENT_FACEKEY_DEL)
}

func (f*Face) Append(key FaceKey, point IPoint) {
	event := Event(0)
	tl, _ := f.Get(key)
	if tl == nil {
		line := NewSeries(f.Resolution)
		f.Register(key, *line)
		tl = *line
		event |= EVENT_FACEKEY_NEW
	}
	appendEvent := tl.Append(point)
	event |= appendEvent
	f.sendNotification(&key, point, event)
}

func (f*FaceConfig) AddNotifier(n Notifier) {
	f.observer = append(f.observer, n)
}

func (f*FaceConfig) RemoveNotifier() {
}

func (f*Face) sendNotification(key *FaceKey, point IPoint, event Event) (ok bool) {
	if f.observer == nil {
		return false
	}
	for _, receiver := range f.observer {
		ok, _ = receiver.Send(key, point, event)
	}
	return
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
