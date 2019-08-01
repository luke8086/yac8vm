/*
 * Copyright (c) 2015 Łukasz S.
 * Distributed under the terms of GPL-2 License.
 */

package src

type Opcode struct {
	Pattern     uint16
	Mask        uint16
	Format      string
	Implemented bool
	Handler     func(*Instr, *Machine)
}

var Opcodes = []Opcode{
	{
		Pattern:     0x00E0,
		Mask:        0xF0FF,
		Format:      "CLS",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			for x := 0; x < len(m.ScreenBuf); x++ {
				for y := 0; y < len(m.ScreenBuf[x]); y++ {
					m.ScreenBuf[x][y] = false
				}
			}
			m.Draw()
		},
	},
	{
		Pattern:     0x00EE,
		Mask:        0xF0FF,
		Format:      "RET",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			if m.SP > 0 {
				m.PC = m.Stack[m.SP]
				m.SP--
			}
		},
	},
	{
		Pattern:     0x0000,
		Mask:        0xF000,
		Format:      "SYS  #addr",
		Implemented: false,
		Handler: func(i *Instr, m *Machine) {
		},
	},
	{
		Pattern:     0x1000,
		Mask:        0xF000,
		Format:      "JP   #addr",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.PC = i.Addr
		},
	},
	{
		Pattern:     0x2000,
		Mask:        0xF000,
		Format:      "CALL #addr",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.SP++
			m.Stack[m.SP] = m.PC
			m.PC = i.Addr
		},
	},
	{
		Pattern:     0x3000,
		Mask:        0xF000,
		Format:      "SE   #vx, #byte",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			if m.V[i.X] == i.Byte {
				m.PC += 2
			}
		},
	},
	{
		Pattern:     0x4000,
		Mask:        0xF000,
		Format:      "SNE  #vx, #byte",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			if m.V[i.X] != i.Byte {
				m.PC += 2
			}
		},
	},
	{
		Pattern:     0x5000,
		Mask:        0xF00F,
		Format:      "SE   #vx, #vy",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			if m.V[i.X] == m.V[i.Y] {
				m.PC += 2
			}
		},
	},
	{
		Pattern:     0x6000,
		Mask:        0xF000,
		Format:      "LD   #vx, #byte",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[i.X] = i.Byte
		},
	},
	{
		Pattern:     0x7000,
		Mask:        0xF000,
		Format:      "ADD  #vx, #byte",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[i.X] += i.Byte
		},
	},
	{
		Pattern:     0x8000,
		Mask:        0xF00F,
		Format:      "LD   #vx, #vy",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[i.X] = m.V[i.Y]
		},
	},
	{
		Pattern:     0x8001,
		Mask:        0xF00F,
		Format:      "OR   #vx, #vy",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[i.X] = m.V[i.X] | m.V[i.Y]
		},
	},
	{
		Pattern:     0x8002,
		Mask:        0xF00F,
		Format:      "AND  #vx, #vy",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[i.X] = m.V[i.X] & m.V[i.Y]
		},
	},
	{
		Pattern:     0x8003,
		Mask:        0xF00F,
		Format:      "XOR  #vx, #vy",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[i.X] = m.V[i.X] ^ m.V[i.Y]
		},
	},
	{
		Pattern:     0x8004,
		Mask:        0xF00F,
		Format:      "ADD  #vx, #vy",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			x := uint16(m.V[i.X]) + uint16(m.V[i.Y])
			m.V[i.X] = uint8(x)
			if x > 0xFF {
				m.V[0xF] = 1
			} else {
				m.V[0xF] = 0
			}
		},
	},
	{
		Pattern:     0x8005,
		Mask:        0xF00F,
		Format:      "SUB  #vx, #vy",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[0xF] = BoolToUint8(m.V[i.X] > m.V[i.Y])
			m.V[i.X] = m.V[i.X] - m.V[i.Y]
		},
	},
	{
		Pattern:     0x8006,
		Mask:        0xF00F,
		Format:      "SHR  #vx, 1",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[0xF] = m.V[i.X] & 0x01
			m.V[i.X] = m.V[i.X] >> 1
		},
	},
	{
		Pattern:     0x8007,
		Mask:        0xF00F,
		Format:      "SUBN #vx, #vy",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[0xF] = BoolToUint8(m.V[i.Y] > m.V[i.X])
			m.V[i.X] = m.V[i.Y] - m.V[i.X]
		},
	},
	{
		Pattern:     0x800E,
		Mask:        0xF00F,
		Format:      "SHL  #vx, 1",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[0xF] = BoolToUint8(m.V[i.X]&0x80 == 0x80)
			m.V[i.X] = m.V[i.X] << 1
		},
	},
	{
		Pattern:     0x9000,
		Mask:        0xF00F,
		Format:      "SNE  #vx, #vy",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			if m.V[i.X] != m.V[i.Y] {
				m.PC += 2
			}
		},
	},
	{
		Pattern:     0xA000,
		Mask:        0xF000,
		Format:      "LD   I, #addr",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.I = i.Addr
		},
	},
	{
		Pattern:     0xB000,
		Mask:        0xF000,
		Format:      "JP   V0, #addr",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.PC = i.Addr + uint16(m.V[0])
		},
	},
	{
		Pattern:     0xC000,
		Mask:        0xF000,
		Format:      "RND  #vx, #byte",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[i.X] = RandUint8() & i.Byte
		},
	},
	{
		Pattern:     0xD000,
		Mask:        0xF000,
		Format:      "DRW  #vx, #vy, #nib",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			xofs := m.V[i.X]
			yofs := m.V[i.Y]
			m.V[0xF] = 0

			for y := uint8(0); y < i.Nib; y++ {
				pixel := m.Mem[m.I+uint16(y)]
				mask := uint8(0x80)

				for x := uint8(0); x < 8; x++ {
					if (pixel & (mask >> x)) == 0 {
						continue
					}

					xpos := (xofs + x) % uint8(len(m.ScreenBuf))
					ypos := (yofs + y) % uint8(len(m.ScreenBuf[0]))
					val := !m.ScreenBuf[xpos][ypos]
					m.ScreenBuf[xpos][ypos] = val

					if !val {
						m.V[0xF] = 1
					}
				}
			}

			m.Draw()
		},
	},
	{
		Pattern:     0xE09E,
		Mask:        0xF0FF,
		Format:      "SKP  #vx",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			if m.Keys[m.V[i.X]] {
				m.PC += 2
			}
		},
	},
	{
		Pattern:     0xE0A1,
		Mask:        0xF0FF,
		Format:      "SKNP #vx",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			if !m.Keys[m.V[i.X]] {
				m.PC += 2
			}
		},
	},
	{
		Pattern:     0xF007,
		Mask:        0xF0FF,
		Format:      "LD   #vx, DT",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.V[i.X] = uint8(m.DelayTimer / ClockMul)
		},
	},
	{
		Pattern:     0xF00A,
		Mask:        0xF0FF,
		Format:      "LD   #vx, K",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			if m.KeyRequested && m.KeyReceived != -1 {
				m.V[i.X] = uint8(m.KeyReceived)
				m.KeyRequested = false
				m.KeyReceived = -1
			} else {
				m.KeyRequested = true
				m.KeyReceived = -1
				m.PC -= 2
			}
		},
	},
	{
		Pattern:     0xF015,
		Mask:        0xF0FF,
		Format:      "LD   DT, #vx",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.DelayTimer = uint16(m.V[i.X]) * ClockMul
		},
	},
	{
		Pattern:     0xF018,
		Mask:        0xF0FF,
		Format:      "LD   ST, #vx",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.SoundTimer = m.V[i.X]
		},
	},
	{
		Pattern:     0xF01E,
		Mask:        0xF0FF,
		Format:      "ADD  I, #vx",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.I += uint16(m.V[i.X])
		},
	},
	{
		Pattern:     0xF029,
		Mask:        0xF0FF,
		Format:      "LD   F, #vx",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			m.I = uint16(m.V[i.X]) * uint16(0x05)
		},
	},
	{
		Pattern:     0xF033,
		Mask:        0xF0FF,
		Format:      "LD   B, #vx",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			mem := &m.Mem
			vx := m.V[i.X]
			mem[m.I+0] = (vx / 100) % 10
			mem[m.I+1] = (vx / 10) % 10
			mem[m.I+2] = (vx / 1) % 10
		},
	},
	{
		Pattern:     0xF055,
		Mask:        0xF0FF,
		Format:      "LD   [I], #vx",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			for j := uint8(0); j < i.X; j++ {
				m.Mem[m.I+uint16(j)] = m.V[j]
			}
		},
	},
	{
		Pattern:     0xF065,
		Mask:        0xF0FF,
		Format:      "LD   #vx, [I]",
		Implemented: true,
		Handler: func(i *Instr, m *Machine) {
			for j := 0; byte(j) <= i.X; j++ {
				m.V[j] = m.Mem[m.I+uint16(j)]
			}
		},
	},
}
