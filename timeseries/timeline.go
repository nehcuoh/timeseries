package timeseries

type TimeLine interface {
	Append(point IPoint) (event EventType)
	Last() (IPoint)
	TimeRange(start int64, stop int64) ([]IPoint)
	Size() int
}
