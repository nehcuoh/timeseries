package test

import (
	"fmt"
	"testing"

	series "zeus/timeseries"
	"zeus/timeseries/face"
)

var (
	faceConfig = series.FaceConfig{Resolution: series.Minute}
)

func TestFace_AppendTo(t *testing.T) {
	uid_face := series.NewFace("test_face", &faceConfig)
	uid_face.Append(face.UID(100000), series.NewPoint(1))
	uid_face.Append(face.UID(100000), series.NewPoint(1))
	line, err := uid_face.Get(face.UID(100000))
	if err != nil {
		t.Errorf("Get uid %d %s", face.UID(100000), err)
	}
	fmt.Println(line)
}

func TestFace_Remove(t *testing.T) {
	uid_face := series.NewFace("test face", &faceConfig)
	uid_face.Register(face.UID(100001), *series.NewSeries(series.Second))
	line, err := uid_face.Get(face.UID(100001))
	if err != nil {
		t.Error("Get timeline err:", err)
	}
	line.Append(series.NewMinutePoint(100))
	val := line.Last().Value()
	if val != 100 {
		t.Error("timeline get err:", err)
	} else {
		t.Logf("Timeline get last value ok,%v", line)
	}
	uid_face.Remove(face.UID(100001))
	line, err = uid_face.Get(face.UID(100001))
	if err == nil {
		t.Error("Remove key but got")
	} else {
		t.Logf("ok, old key is removed succ,%v", err)
	}
	if line != nil {
		t.Errorf("old line is not nil: %s", line)
	}
}

func benchmark_AppendTo(b *testing.B, f *series.Face, len int32) {
	for i := 0; i < b.N; i++ {
		for n := uint64(0); n < uint64(len); n++ {
			f.Append(face.UID(n), series.NewPoint(1))
		}
	}
}

func Benchmark_AppendTo_TwoDay(b *testing.B) {
	uid_face := series.NewFace("test face", &faceConfig)
	b.ReportAllocs()
	b.SetBytes(2)
	benchmark_AppendTo(b, uid_face, 2880)
}

func Benchmark_AppendTo_FullDay(b *testing.B) {
	uid_face := series.NewFace("test face", &faceConfig)
	b.ReportAllocs()
	b.SetBytes(2)
	benchmark_AppendTo(b, uid_face, 1440)
}

func Benchmark_AppendTo_12Hour(b *testing.B) {
	uid_face := series.NewFace("test face", &faceConfig)
	b.ReportAllocs()
	b.SetBytes(2)
	benchmark_AppendTo(b, uid_face, 720)
}

func Benchmark_AppendTo_1Hour(b *testing.B) {
	uid_face := series.NewFace("test face", &faceConfig)
	b.ReportAllocs()
	b.SetBytes(2)
	benchmark_AppendTo(b, uid_face, 60)
}
func Benchmark_AppendTo_5Minutes(b *testing.B) {
	uid_face := series.NewFace("test face", &faceConfig)
	b.ReportAllocs()
	b.SetBytes(2)
	benchmark_AppendTo(b, uid_face, 5)
}
