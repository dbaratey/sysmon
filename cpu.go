package sysmon

import (
	"log"
	"os/exec"
	"time"
)

type CpuWatcher struct{}

func (cw *CpuWatcher) Get() ([]WatchInfo, error) {
	ts := time.Now().Unix()
	cmd := exec.Command(CPU_COMMAND[0], CPU_COMMAND[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	userV, sysV, idleV, err := ParceCPU(out)
	if err != nil {
		return []WatchInfo{}, err
	}
	return []WatchInfo{{Name: "cpu:user",
		Ts:    ts,
		Value: userV},
		{
			Name:  "cpu:system",
			Ts:    ts,
			Value: sysV,
		}, {
			Name:  "cpu:idle",
			Ts:    ts,
			Value: idleV,
		},
	}, nil

}
