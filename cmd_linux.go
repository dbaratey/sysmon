// +build linux

package sysmon

var (
	CPU_COMMAND = []string{"top", "-bn1"}
)

func ParceCPU(out []byte) (float64, float64, float64, error) {
	return 0, 0, 0, nil
}
