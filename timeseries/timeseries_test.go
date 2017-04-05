package timeseries

import (
	"testing"
	"time"
)

//func TestNewPoint(t *testing.T) {
//	value := int32(1231)
//	point := NewPoint(value)
//	if point.Value != value {
//		t.Error("NewPoint assert failed")
//	} else {
//		t.Log("time series test ok")
//	}
//}

func TestNewMinutePoint(t *testing.T) {
	value := int32(1234)
	point := NewMinutePoint(value)
	now := time.Now().Unix()
	if point.Value() != value {
		t.Error("new Minute point value incorrect")
	}
	now = now - (now % Minute)
	if now != point.TimeStamp() {
		t.Errorf("nMinutes mismatch,now= %d, nMinutes=%d\n", now, point.TimeStamp())
	}
}

