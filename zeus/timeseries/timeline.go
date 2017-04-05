package timeseries

type TimeLine interface {
	Append(point IPoint) (event Event)
	Last() (IPoint)
	TimeRange(start int64, stop int64) ([]IPoint)
	Size() int
}
