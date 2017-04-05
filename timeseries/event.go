package timeseries

const (
	EVENT_FACEKEY_NEW         = 0x01
	EVENT_FACEKEY_DEL         = 0x02
	EVENT_TIMELINE_POINT_INCR = 0x04
	EVENT_TIMELINE_POINT_NEW  = 0x08
)

type Event uint16

type Notifier interface {
	Send(key *FaceKey, point IPoint, event Event) (ok bool, err Error)
}
