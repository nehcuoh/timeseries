package main

import (
	"fmt"
	//"time"

	"timeseries"
)

func main() {
	tl := timeseries.NewSeries(timeseries.Minute)
	(*tl).Append(timeseries.NewPoint(100))
	(*tl).Append(timeseries.NewMinutePoint(100))
	(*tl).Append(timeseries.NewMinutePoint(100))
	(*tl).Append(timeseries.NewPoint(100))
	//time.Sleep(61 * time.Second)
	(*tl).Append(timeseries.NewPoint(100))
	t, _ := (*tl).(*timeseries.MinuteSeries)
	fmt.Println(t)
	fmt.Printf("len: %d\n", t.Len())

	tl = timeseries.NewSeries(timeseries.Second)
	(*tl).Append(timeseries.NewMinutePoint(1000))
	(*tl).Append(timeseries.NewMinutePoint(1000))
	(*tl).Append(timeseries.NewMinutePoint(1000))
	(*tl).Append(timeseries.NewMinutePoint(1000))
	(*tl).Append(timeseries.NewMinutePoint(1000))
	t2, _ := (*tl).(*timeseries.Series)
	fmt.Println(t2)
	fmt.Printf("len: %d\n", t2.Len())

}
