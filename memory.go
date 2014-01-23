package rp2ago3

import (
	"errors"
	"github.com/nwidger/m65go2"
)

type MappableMemory interface {
	m65go2.Memory
	Mappings() []uint16
}

type MappedMemory struct {
	maps map[uint16]m65go2.Memory
	m65go2.Memory
}

func NewMappedMemory(base m65go2.Memory) *MappedMemory {
	return &MappedMemory{maps: make(map[uint16]m65go2.Memory, 0xffff), Memory: base}
}

func (mem *MappedMemory) AddMappings(mappable MappableMemory) (err error) {
	addresses := mappable.Mappings()

	for _, address := range addresses {
		if _, ok := mem.maps[address]; ok {
			err = errors.New("Address is already mapped")
			return
		}

		mem.maps[address] = mappable
	}

	return
}

func (mem *MappedMemory) Reset() {
	// don't clear mappings
	mem.Memory.Reset()
}

func (mem *MappedMemory) Fetch(address uint16) (value uint8) {
	if mmap, ok := mem.maps[address]; ok {
		return mmap.Fetch(address)
	}

	return mem.Memory.Fetch(address)
}

func (mem *MappedMemory) Store(address uint16, value uint8) (oldValue uint8) {
	if mmap, ok := mem.maps[address]; ok {
		return mmap.Store(address, value)
	}

	return mem.Memory.Store(address, value)
}
