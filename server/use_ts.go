package main

import (
	"fmt"
	"time"

	s "zeus/timeseries"
	f "zeus/timeseries/face"
)

var (
	default_config = s.FaceConfig{Resolution: s.Second}
)

func main() {
	tl := s.NewSeries(s.Minute)
	p := s.NewMinutePoint(100)
	p.SetTimeStamp(time.Now().Unix())
	p.SetValue(1000)
	(*tl).Append(s.NewPoint(100))
	(*tl).Append(s.NewMinutePoint(100))
	(*tl).Append(s.NewMinutePoint(100))
	(*tl).Append(s.NewPoint(100))
	//time.Sleep(61 * time.Second)
	(*tl).Append(s.NewPoint(100))
	t, _ := (*tl).(*s.MinuteSeries)
	fmt.Println(t)
	fmt.Printf("len: %d\n", t.Size())

	tl = s.NewSeries(s.Second)
	(*tl).Append(s.NewMinutePoint(1000))
	(*tl).Append(s.NewMinutePoint(1000))
	(*tl).Append(s.NewMinutePoint(1000))
	(*tl).Append(s.NewMinutePoint(1000))
	(*tl).Append(s.NewMinutePoint(1000))
	t2, _ := (*tl).(*s.Series)
	fmt.Println(t2)
	fmt.Printf("len: %d\n", t2.Size())

	uid_log := s.NewFace("test face", &default_config)
	//uid_log.AddNotifier()
	uid_log.Dump(&f.Uid{})
	for i := 0; i < 100000; i++ {
		for j := 0; j < 1440; j++ {
			uid_log.Append(f.UID(uint64(i)), s.NewPoint(1))
		}
	}
	line, _ := uid_log.Get(f.UID(uint64(0)))

	///l, ok := line.(s.Series)
	///if !ok {

	///}
	///l, ok := line.(s.MinuteSeries)
	///if !ok {
	///	//
	//}
	fmt.Println("line size", line.Size(), "last value: ", line.Last().Value())
	fmt.Println("append complete")
	time.Sleep(time.Minute)
}
