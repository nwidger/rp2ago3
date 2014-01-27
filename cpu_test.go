package rp2ago3

import (
	"github.com/nwidger/m65go2"
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	mem := NewMappedMemory(m65go2.NewBasicMemory())
	clock := m65go2.NewClock(1 * time.Nanosecond)
	cpu := NewRP2A03(mem, clock, 12)
	cpu.Reset()
	clock.Start()

	cpu.APU.Registers.Pulse1[0] = 0xde
	cpu.Memory.Store(0x4000, 0xff)

	if cpu.APU.Registers.Pulse1[0] != 0xff {
		t.Error("Register is not 0xff")
	}

	cpu.Memory.Store(0x0800, 0xff)

	if cpu.Memory.Fetch(0x0000) != 0xff {
		t.Error("Memory is not 0xff")
	}

	cpu.Memory.Store(0x0800, 0x00)

	if cpu.Memory.Fetch(0x0000) != 0x00 {
		t.Error("Memory is not 0x00")
	}

	clock.Stop()
}
