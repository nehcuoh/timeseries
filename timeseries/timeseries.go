package timeseries

import (
	"fmt"
	"time"
)

type IPoint interface {
	Value() int32
	Incr(int32) int32
	TimeStamp() int64
}

type TimeLine interface {
	Append(point IPoint)
	Last() (IPoint)
	TimeRange(start int64, stop int64) ([]IPoint)
}

type Series []IPoint

type internalSeries struct {
	Series
	Resolution int32
}

type Point struct {
	value     int32
	timestamp int64
}

type FaceKey interface {
	Key() interface{}
}

type Face struct {
	m map[interface{}]TimeLine
}

const (
	EdenTime = 1490345760
	Second   = 1
	Minute   = 60
	Hour     = 3600
	Day      = 86400
)

func NewPoint(v int32) (*Point) {
	now := time.Now().Unix()
	return &Point{value: v, timestamp: now}
}

func (p*Point) Value() int32 {
	return p.value
}

func (p*Point) Incr(deltas int32) int32 {
	p.value += deltas
	return p.value
}

func (p*Point) TimeStamp() int64 {
	return p.timestamp
}

func NewSeries(resolution int32) (*TimeLine) {
	var ret *TimeLine
	ret = new(TimeLine)
	if resolution == Second {
		//s := Series{}
		s := new(Series)
		*ret = s
	}
	if resolution == Minute {
		s := NewMinuteSeries()
		*ret = s
	}

	return ret
}

func (s *Series) Append(point IPoint) {

	old_size := len(*s)
	if old_size < 1 {
		*s = append(*s, point)
		return
	}

	lastPoint := (*s)[old_size-1]
	lastStamp := lastPoint.TimeStamp()
	if point.TimeStamp() == lastStamp {
		lastPoint.Incr(point.Value())
		return
	}
	*s = append(*s, point)
}

func (s*Series) Last() (IPoint) {
	return (*s)[len(*s)-1]
}

func (s*Series) TimeRange(start int64, stop int64) ([]IPoint) {
	return nil
}

func (s*Series) Len() (int) {
	return len(*s)
}

func (s *Series) String() string {
	ret := string("")
	for _, p := range *s {
		ret = ret + fmt.Sprintf("timestamp: %d, value: %d\n", p.TimeStamp(), p.Value())
	}
	return ret
}

func NewFace() (*Face) {
	return &Face{m: make(map[interface{}]TimeLine)}
}

func (f*Face) Register(key FaceKey, line TimeLine) {
	f.m[key.Key()] = line
}

func (f*Face) Get(key FaceKey) (t TimeLine, ok bool) {
	e, ok := f.m[key.Key()]
	if !ok {
		return nil, false
	}
	return e, true
}

func (f*Face) Remove(key FaceKey) {
	delete(f.m, key.Key())
}

func (f*Face) AppendTo(key FaceKey, point IPoint, resolution int32) {
	tl, ok := f.Get(key)
	if !ok {
		line := NewSeries(resolution)
		f.Register(key, *line)
		tl = *line
	}
	tl.Append(point)
}

func (f*Face) Dump() {
	for i, l := range f.m {
		fmt.Printf("Face[%s] =\n%s\n", i, l)
	}
}
