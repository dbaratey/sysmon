package sysmon

import (
	"testing"
)

func BenchmarkCpuWatcher_Get(b *testing.B) {
	w := CpuWatcher{}
	for n := 0; n < b.N; n++ {
		w.Get()
	}
}
