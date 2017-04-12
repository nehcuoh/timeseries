package timeseries

const (
	EVENT_FACEKEY_NEW         = 0x01
	EVENT_FACEKEY_DEL         = 0x02
	EVENT_TIMELINE_POINT_INCR = 0x04
	EVENT_TIMELINE_POINT_NEW  = 0x08
)

type EventType uint16

type Event struct {
	Type  EventType
	key   *FaceKey
	point *IPoint
}

type Notifier interface {
	Send(e *Event) ( err Error)
}
