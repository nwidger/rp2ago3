package rp2ago3

import (
	"github.com/nwidger/m65go2"
)

type DMA struct {
	Memory  m65go2.Memory
	Clock   m65go2.Clocker
	pending chan uint32
}

func NewDMA(memory m65go2.Memory, clock m65go2.Clocker) *DMA {
	return &DMA{
		Memory:  memory,
		Clock:   clock,
		pending: make(chan uint32, 1),
	}
}

func (dma *DMA) PerformDMA() {
	select {
	case start := <-dma.pending:
		ticks := dma.Clock.Ticks()
		end := uint32(start + 0x0100)

		for address := start; address < end; address++ {
			dma.Memory.Store(0x2004, dma.Memory.Fetch(uint16(address)))
		}

		dma.Clock.Await(ticks + 513 + (ticks & 0x1))
	default:
	}
}

func (dma *DMA) Reset() {
	dma.pending = make(chan uint32, 1)
}

func (dma *DMA) Mappings(which Mapping) (fetch, store []uint16) {
	switch which {
	case CPU:
		store = []uint16{0x4014}
	}

	return
}

func (dma *DMA) Fetch(address uint16) (value uint8) {
	// nothing to do
	return
}

func (dma *DMA) Store(address uint16, value uint8) (oldValue uint8) {
	switch address {
	case 0x4014:
		dma.pending <- (uint32(value) << 8)
	}

	return
}
