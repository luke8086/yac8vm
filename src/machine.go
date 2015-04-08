/*
 * Copyright (c) 2015 Łukasz S.
 * Distributed under the terms of GPL-2 License.
 */

package yac8vm

import (
	"bytes"
	"errors"
	"io"
	"time"
)

type ScreenBuf [ScreenW][ScreenH]bool

type Machine struct {
	Mem   [MemSize]uint8
	Stack [StackSize]uint16

	V          [RegCount]uint8
	I          uint16
	PC         uint16
	SP         uint8
	DelayTimer uint16
	SoundTimer uint8

	KeyRequested  bool
	KeyReceived   int8
	Paused        bool
	ExitRequested bool

	ScreenBuf ScreenBuf
	Keys      [KeyCount]bool

	GUI *GUI
}

func NewMachine(gui *GUI, romReader io.Reader) (*Machine, error) {
	m := &Machine{GUI: gui, PC: ProgStart}

	if _, err := bytes.NewReader(InitialMemory).Read(m.Mem[0:]); err != nil {
		return nil, err
	}

	if _, err := romReader.Read(m.Mem[ProgStart:]); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Machine) Run() error {
	dt := time.Second / ClockFreq / ClockMul

	for {
		for timeout := time.Now().Add(dt); time.Now().Before(timeout); {
			m.GUI.ProcessEvent(m)
		}

		if m.ExitRequested {
			return nil
		}

		if m.Paused {
			continue
		}

		if err := m.Tick(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Machine) FetchInstr() (*Instr, error) {
	if int(s.PC)+1 >= len(s.Mem) {
		return nil, errors.New("invalid PC")
	}

	hb := s.Mem[s.PC]
	lb := s.Mem[s.PC+1]
	w := uint16(hb)<<8 | uint16(lb)
	op := NewInstr(w)

	s.PC += 2

	return op, nil
}

func (m *Machine) Tick() error {
	i, err := m.FetchInstr()
	if err != nil {
		return err
	}

	if i.Valid {
		i.Opcode.Handler(i, m)
	} else {
		return errors.New("invalid instruction")
	}

	if m.DelayTimer > 0 {
		m.DelayTimer--
	}

	if m.SoundTimer > 0 {
		m.SoundTimer--
	}

	return nil
}

func (m *Machine) Draw() {
	m.GUI.Draw(m)
}

func (m *Machine) ToggleKey(key int8, pressed bool) {
	m.Keys[key] = pressed
	if m.KeyRequested && pressed {
		m.KeyReceived = key
	}
}

func (m *Machine) TogglePaused(paused bool) {
	m.Paused = paused
}

func (m *Machine) RequestExit() {
	m.ExitRequested = true
}
