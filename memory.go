package rp2ago3

import (
	"errors"
	"github.com/nwidger/m65go2"
)

type MappableMemory interface {
	m65go2.Memory
	Mappings() (fetch, store []uint16)
}

type MappedMemory struct {
	fetch map[uint16]m65go2.Memory
	store map[uint16]m65go2.Memory
	m65go2.Memory
}

func NewMappedMemory(base m65go2.Memory) *MappedMemory {
	return &MappedMemory{
		fetch:  make(map[uint16]m65go2.Memory, 0xffff),
		store:  make(map[uint16]m65go2.Memory, 0xffff),
		Memory: base,
	}
}

func (mem *MappedMemory) AddMappings(mappable MappableMemory) (err error) {
	fetch, store := mappable.Mappings()

	for _, address := range fetch {
		if _, ok := mem.fetch[address]; ok {
			err = errors.New("Address is already mapped for fetch")
			return
		}

		mem.fetch[address] = mappable
	}

	for _, address := range store {
		if _, ok := mem.store[address]; ok {
			err = errors.New("Address is already mapped for store")
			return
		}

		mem.store[address] = mappable
	}

	return
}

func (mem *MappedMemory) Reset() {
	// don't clear mappings
	mem.Memory.Reset()
}

func (mem *MappedMemory) Fetch(address uint16) (value uint8) {
	if mmap, ok := mem.fetch[address]; ok {
		return mmap.Fetch(address)
	}

	return mem.Memory.Fetch(address)
}

func (mem *MappedMemory) Store(address uint16, value uint8) (oldValue uint8) {
	if mmap, ok := mem.store[address]; ok {
		return mmap.Store(address, value)
	}

	return mem.Memory.Store(address, value)
}
