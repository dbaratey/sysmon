package sysmon

import (
	"testing"
	"time"
)

//func TestCore_Add(t *testing.T) {
//	c := NewCore(10)
//	for i := 0; i < 100; i++ {
//		c.Add(WatchInfo{
//			Name:  "k",
//			Ts:    time.Now().Unix(),
//			Value: 0,
//		})
//		time.Sleep(time.Millisecond * 500)
//		if len(c.series["k"]) > 20 {
//			t.Fatal("TS crop not working",len(c.series["k"]))
//		}
//	}
//}

func TestCore_Avg(t *testing.T) {
	c := NewCore(10)
	for i := 0; i < 100; i++ {
		c.Add(WatchInfo{
			Name:  "k",
			Ts:    time.Now().Unix(),
			Value: float64(i + 1),
		})
	}
	if c.Avg("k") != 50.5 {
		t.Fatal("Avg not working correct")
	}
}
