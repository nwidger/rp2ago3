package rp2ago3

import (
	"errors"
	"github.com/nwidger/m65go2"
)

type MappedMemory struct {
	maps map[uint16]m65go2.Memory
	m65go2.Memory
}

func NewMappedMemory(base m65go2.Memory) *MappedMemory {
	return &MappedMemory{maps: make(map[uint16]m65go2.Memory), Memory: base}
}

func (mem *MappedMemory) AddMap(addresses []uint16, mmap m65go2.Memory) (error error) {
	for _, address := range addresses {
		if _, ok := mem.maps[address]; ok {
			return errors.New("Address is already mapped")
		}

		mem.maps[address] = mmap
	}

	return nil
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
