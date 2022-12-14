package tlb

import (
	"log"

	"gitlab.com/akita/mem/v2/vm"
	"gitlab.com/akita/util/v2/ca"
)

type mshrEntry struct {
	pid         ca.PID
	vAddr       uint64
	Requests    []*vm.TranslationReq
	reqToBottom *vm.TranslationReq
	page        vm.Page
}

// newMSHREntry returns a new MSHR entry object
func newMSHREntry() *mshrEntry {
	e := new(mshrEntry)
	return e
}

// mshr is an interface that controls MSHR entries
type mshr interface {
	Query(pid ca.PID, addr uint64) *mshrEntry
	Add(pid ca.PID, addr uint64) *mshrEntry
	Remove(pid ca.PID, addr uint64) *mshrEntry
	AllEntries() []*mshrEntry
	IsFull() bool
	Reset()
	GetEntry(pid ca.PID, vAddr uint64) *mshrEntry
	IsEntryPresent(pid ca.PID, vAddr uint64) bool
}

type mshrImpl struct {
	capacity int
	entries  []*mshrEntry
}

// newMSHR returns a new mshr object
func newMSHR(capacity int) mshr {
	m := new(mshrImpl)
	m.capacity = capacity
	return m
}

func (m *mshrImpl) Add(pid ca.PID, vAddr uint64) *mshrEntry {
	for _, e := range m.entries {
		if e.pid == pid && e.vAddr == vAddr {
			panic("entry already in mshr")
		}
	}

	if len(m.entries) >= m.capacity {
		log.Panic("MSHR is full")
	}

	entry := newMSHREntry()
	entry.pid = pid
	entry.vAddr = vAddr
	m.entries = append(m.entries, entry)
	return entry
}

func (m *mshrImpl) Query(pid ca.PID, vAddr uint64) *mshrEntry {
	for _, e := range m.entries {
		if e.pid == pid && e.vAddr == vAddr {
			return e
		}
	}
	return nil
}

func (m *mshrImpl) Remove(pid ca.PID, vAddr uint64) *mshrEntry {
	for i, e := range m.entries {
		if e.pid == pid && e.vAddr == vAddr {
			m.entries = append(m.entries[:i], m.entries[i+1:]...)
			return e
		}
	}
	panic("trying to remove an non-exist entry")
}

func (m *mshrImpl) AllEntries() []*mshrEntry {
	return m.entries
}

func (m *mshrImpl) IsFull() bool {
	return len(m.entries) >= m.capacity
}

func (m *mshrImpl) Reset() {
	m.entries = nil
}

func (m *mshrImpl) GetEntry(pid ca.PID, vAddr uint64) *mshrEntry {
	for _, e := range m.entries {
		if e.pid == pid && e.vAddr == vAddr {
			return e
		}
	}
	return nil
}

func (m *mshrImpl) IsEntryPresent(pid ca.PID, vAddr uint64) bool {
	for _, e := range m.entries {
		if e.pid == pid && e.vAddr == vAddr {
			return true
		}
	}
	return false
}
