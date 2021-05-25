package domain

import (
	"encoding/binary"
	"github.com/ereb-or-od/oakenshield/pkg/domain/interfaces"
	"github.com/ereb-or-od/oakenshield/pkg/infrastructure/utilities"
	"sync"
	"time"
)

// [interval(4byte)][sequence/random(3byte)][machine(1byte)]
const (
	intervalBits    = 32
	sequenceBits    = 23
	machineIdBits   = 8
	ignoredTimeBits = 30
	intervalMask    = (1 << intervalBits) - 1
	machineIdMask   = (1 << machineIdBits) - 1
)

type oakenshildID struct {
	mutex           *sync.Mutex
	machineId       byte
	epochStart      int64
	sequence        int32
	currentInterval int64
}


func (o oakenshildID) Next() uint64 {
	raw := o.next()

	// Shuffle bits
	uid := make([]byte, 8, 8)
	for i := int64(0); i < 8; i++ {
		for l := int64(0); l < 8; l++ {
			uid[l] |= byte((raw & (1 << (i*8 + l))) >> (i*7 + l))
		}
	}
	return binary.LittleEndian.Uint64(uid)
}

func (o oakenshildID) WithMachineId(machineId byte) interfaces.OakenshildID {
	o.machineId = machineId
	return o
}

func (o oakenshildID) WithEpochStart(time time.Time) interfaces.OakenshildID {
	o.epochStart = time.UnixNano()
	return o
}

func NewOakenshildID() interfaces.OakenshildID {
	return &oakenshildID{
		mutex:      &sync.Mutex{},
		machineId:  byte(utilities.GetLocalIPv4() & machineIdMask),
		epochStart: 1577833200000000000, // 1/1/2020
	}
}

func (o *oakenshildID) next() int64 {
	// 32 bit time interval with nano-time >> 20 (~1s) clock loops after reaching end of epoch each ~ 146 years
	interval := ((time.Now().UnixNano() - o.epochStart) >> ignoredTimeBits) & intervalMask
	// 23 bit sequence and random
	sequence := int32(0)
	o.mutex.Lock()
	loop := (o.sequence + 0x400000 - 0x2020) >> sequenceBits // 4194304 - 8224 = 4186080
	if interval-int64(loop) <= o.currentInterval {
		o.sequence++
		if o.sequence < 0x20 {
			// Small counter and 2 random bytes
			sequence = (o.sequence << 16) | (utilities.GenerateRandomBytes() << 8) | utilities.GenerateRandomBytes()
		} else if o.sequence < 0x2020 {
			// Enlarge the counter
			sequence = (0x200000 - 0x2000 + (o.sequence << 8)) | utilities.GenerateRandomBytes()
		} else {
			// Use all space for the counter
			sequence = 0x400000 - 0x2020 + o.sequence
		}
	} else {
		o.currentInterval = interval
		o.sequence = int32(0)
	}
	o.mutex.Unlock()

	raw := interval
	raw = (raw << sequenceBits) + int64(sequence) // + to increment the interval too on rollover
	raw = (raw << machineIdBits) | int64(o.machineId)

	return raw
}
