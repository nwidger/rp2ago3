package rp2ago3

import (
	"github.com/nwidger/m65go2"
	"time"
)

const NTSC_CLOCK_RATE time.Duration = 46 * time.Nanosecond // 21.477272Mhz
const PAL_CLOCK_RATE time.Duration = 37 * time.Nanosecond  // 26.601712MHz

const NTSC_CPU_CLOCK_DIVISOR uint64 = 12
const PAL_CPU_CLOCK_DIVISOR uint64 = 16

type RP2A03 struct {
	*m65go2.M6502
	*APU
	clock  *m65go2.Divider
	Memory *MirroredMemory
}

func NewRP2A03(mem *MappedMemory, clock m65go2.Clocker, divisor uint64) *RP2A03 {
	divider := m65go2.NewDivider(clock, divisor)
	cpu := m65go2.NewM6502(mem, divider)
	cpu.DisableDecimalMode()
	apu := NewAPU(divider)

	// APU memory maps
	mem.AddMappings(apu)

	mirrored := NewMirroredMemory(mem)
	mirrors := make(map[uint16]uint16)

	// Mirrored 2KB internal RAM
	for i := uint16(0x0800); i <= 0x1fff; i++ {
		mirrors[i] = i % 0x0800
	}

	mirrored.AddMirrors(mirrors)

	return &RP2A03{Memory: mirrored, M6502: cpu, APU: apu, clock: divider}
}

func (cpu *RP2A03) Reset() {
	cpu.M6502.Reset()
	cpu.APU.Reset()
	cpu.Memory.Reset()
}
