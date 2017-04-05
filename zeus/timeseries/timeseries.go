package timeseries

import (
	"fmt"
	"time"
)

type Error struct {
	error string
}

func (e*Error) Error() string {
	return e.error
}

type IPoint interface {
	Value() int32
	Incr(int32) int32
	TimeStamp() int64
}

type Series []IPoint

type Point struct {
	value     int32
	timestamp int64
}

type TimeInterval int32

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

//仅限调试设置假时间戳
func (p*Point) SetTimeStamp(time int64) {
	p.timestamp = int64(time)
}

func NewSeries(resolution TimeInterval) (*TimeLine) {
	var ret *TimeLine
	ret = new(TimeLine)
	if resolution == Second {
		s := new(Series)
		*ret = s
	}
	if resolution == Minute {
		s := NewMinuteSeries()
		*ret = s
	}

	return ret
}

func (s *Series) Append(point IPoint) (e Event) {

	old_size := len(*s)
	if old_size < 1 {
		*s = append(*s, point)
		return EVENT_TIMELINE_POINT_NEW
	}
	//todo:
	lastPoint := (*s)[old_size-1]
	lastStamp := lastPoint.TimeStamp()
	if (point).TimeStamp() == lastStamp {
		lastPoint.Incr((point).Value())
		return EVENT_TIMELINE_POINT_INCR
	}
	*s = append(*s, point)
	return EVENT_TIMELINE_POINT_NEW
}

func (s*Series) Last() (IPoint) {
	return (*s)[len(*s)-1]
}

func (s*Series) TimeRange(start int64, stop int64) ([]IPoint) {
	return nil
}

func (s*Series) Size() (int) {
	return len(*s)
}

func (s *Series) String() string {
	ret := string("")
	time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	for _, p := range *s {
		ret = ret + fmt.Sprintf("timestamp: %s, value: %d\n", time.Unix(p.TimeStamp(), 0).Format("2006-01-02 03:04:05"), p.Value())
	}
	return ret
}
