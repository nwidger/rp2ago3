package rp2ago3

import (
	"github.com/nwidger/m65go2"
)

const NTSC_CLOCK_DIVISOR uint64 = 12
const PAL_CLOCK_DIVISOR uint64 = 16

type RP2A03 struct {
	*m65go2.M6502
	*APU
	clock *m65go2.Divider
}

func NewRP2A03(mem *MappedMemory, clock m65go2.Clocker, divisor uint64) *RP2A03 {
	divider := m65go2.NewDivider(clock, divisor)
	cpu := m65go2.NewM6502(mem, divider)
	apu := NewAPU(divider)

	// APU memory maps
	mem.AddMap([]uint16{
		0x4000, 0x4001, 0x4002, 0x4003, 0x4004,
		0x4005, 0x4006, 0x4007, 0x4008, 0x400a,
		0x400b, 0x400c, 0x400e, 0x400f, 0x4010,
		0x4011, 0x4012, 0x4013, 0x4015, 0x4017,
	}, apu)

	return &RP2A03{M6502: cpu, APU: apu}
}
