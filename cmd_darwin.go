// +build darwin

package sysmon

import (
	"regexp"
	"strconv"
)



var (
	CPU_COMMAND = []string{"top", "-l1"}
)

func ParceCPU(out []byte) (float64, float64, float64, error) {
	t := regexp.MustCompile("CPU usage:.*")
	tVal := regexp.MustCompile("CPU usage: (.+)% user, (.+)% sys, (.+)% idle")
	s := t.FindSubmatch(out)
	data := tVal.FindSubmatch(s[0])
	uV, err := strconv.ParseFloat(string(data[1]), 64)
	if err != nil {
		return 0, 0, 0, err
	}
	sysV, err := strconv.ParseFloat(string(data[2]), 64)
	if err != nil {
		return 0, 0, 0, err
	}
	idleV, err := strconv.ParseFloat(string(data[3]), 64)
	if err != nil {
		return 0, 0, 0, err
	}
	return uV, sysV, idleV, nil
}
