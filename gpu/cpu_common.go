package gpu

import (
	"golang.org/x/sys/cpu"
)

func GetCPUCapability() CPUCapability {
	return CPUCapabilityAVX
}
