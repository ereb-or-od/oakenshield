package interfaces

import (
	"time"
)

type OakenshildID interface {
	Next() uint64
	WithMachineId(machineId byte) OakenshildID
	WithEpochStart(time time.Time) OakenshildID
}