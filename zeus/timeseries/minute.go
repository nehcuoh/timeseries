package timeseries

import (
	"time"
)

type MinutePoint struct {
	value        int32
	minuteDeltas uint32
}

func NewMinutePoint(value int32) *MinutePoint {
	now := time.Now().Unix()
	now = now - (now % Minute)
	d := now - EdenTime
	deltas := uint32(d / Minute)

	return &MinutePoint{value: value, minuteDeltas: deltas}
}
func (m MinutePoint) TimeStamp() int64 {
	real_time := int64(m.minuteDeltas*Minute + EdenTime)
	return real_time
}
func (m MinutePoint) Value() int32 {
	return m.value
}

func (m*MinutePoint) Incr(value int32) (new_value int32) {
	m.value += value
	return m.value
}

//仅限调试设置假时间戳
func (m*MinutePoint) SetTimeStamp(time int64) {
	t := time - (time % Minute)
	d := t - EdenTime
	deltas := uint32(d / Minute)
	m.minuteDeltas = deltas
}

func (m*MinutePoint) SetValue(v int32) {
	m.value = v
}

type MinuteSeries struct {
	Series
}

func NewMinuteSeries() (*MinuteSeries) {
	s := Series{}
	return &MinuteSeries{Series: s}
}

func (m*MinuteSeries) Append(point IPoint) (event Event) {
	if m.Series == nil {
		m.Series = make([]IPoint, 1)
	}
	p, ok := point.(*MinutePoint)
	if !ok {
		p = NewMinutePoint(point.Value())
	}

	old_size := len(m.Series)
	if old_size < 1 {
		m.Series = append(m.Series, p)
		return EVENT_TIMELINE_POINT_NEW
	}

	lastPoint := m.Series[old_size-1]
	lastStamp := lastPoint.TimeStamp()
	if p.TimeStamp() == lastStamp {
		lastPoint.Incr(point.Value())
		return EVENT_TIMELINE_POINT_INCR
	}

	m.Series = append(m.Series, p)
	return EVENT_TIMELINE_POINT_NEW
}

func (m *MinuteSeries) Last() (point IPoint) {
	return m.Series[len(m.Series)-1]
}

func (m*MinuteSeries) TimeRange(start int64, stop int64) ([]IPoint) {
	return nil
}

func (m*MinuteSeries) Size() (int) {
	return len(m.Series)
}
