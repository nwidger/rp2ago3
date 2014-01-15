package rp2ago3

import (
	"github.com/nwidger/m65go2"
	"testing"
)

func TestStore(t *testing.T) {
	mem := NewMappedMemory(m65go2.NewBasicMemory())
	clock := m65go2.NewClock(m65go2.NTSC_CLOCK_RATE)
	cpu := NewRP2A03(mem, clock, NTSC_CLOCK_DIVISOR)
	cpu.Reset()
	go clock.Start()

	cpu.APU.Registers.Pulse1[0] = 0xde
	cpu.Memory.Store(0x4000, 0xff)

	if cpu.APU.Registers.Pulse1[0] != 0xff {
		t.Error("Register is not 0xff")
	}

	clock.Stop()
}
