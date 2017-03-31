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
