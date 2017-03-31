package main

import (
	"fmt"
	"time"

	s "timeseries"
	f "timeseries/Face"
)

func main() {
	tl := s.NewSeries(s.Minute)
	(*tl).Append(s.NewPoint(100))
	(*tl).Append(s.NewMinutePoint(100))
	(*tl).Append(s.NewMinutePoint(100))
	(*tl).Append(s.NewPoint(100))
	//time.Sleep(61 * time.Second)
	(*tl).Append(s.NewPoint(100))
	t, _ := (*tl).(*s.MinuteSeries)
	fmt.Println(t)
	fmt.Printf("len: %d\n", t.Len())

	tl = s.NewSeries(s.Second)
	(*tl).Append(s.NewMinutePoint(1000))
	(*tl).Append(s.NewMinutePoint(1000))
	(*tl).Append(s.NewMinutePoint(1000))
	(*tl).Append(s.NewMinutePoint(1000))
	(*tl).Append(s.NewMinutePoint(1000))
	t2, _ := (*tl).(*s.Series)
	fmt.Println(t2)
	fmt.Printf("len: %d\n", t2.Len())

	login_ip := s.NewFace()
	login_ip.AppendTo(f.ParseIP("192.1.1.1"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.1"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.1"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.2"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.2"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.2"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.2"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.3"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.3"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.4"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.4"), s.NewMinutePoint(1), s.Minute)
	login_ip.Dump()
	time.Sleep(time.Minute + (2 * time.Second))
	login_ip.AppendTo(f.ParseIP("192.1.1.1"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.1"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.1"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.2"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.2"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.2"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.2"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.3"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.3"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.4"), s.NewMinutePoint(1), s.Minute)
	login_ip.AppendTo(f.ParseIP("192.1.1.4"), s.NewMinutePoint(1), s.Minute)
	login_ip.Dump()
	//login_ip.Get(f.ParseIP("192.1.1.1"))

}
