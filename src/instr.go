/*
 * Copyright (c) 2015 Łukasz S.
 * Distributed under the terms of GPL-2 License.
 */

package yac8vm

import (
	"fmt"
	"strings"
)

type Instr struct {
	Raw    uint16
	X      uint8
	Y      uint8
	Byte   uint8
	Nib    uint8
	Addr   uint16
	Valid  bool
	Opcode *Opcode
}

func NewInstr(raw uint16) *Instr {
	i := &Instr{
		Raw:    raw,
		Valid:  false,
		Opcode: nil,
		X:      uint8((raw & 0x0F00) >> 8),
		Y:      uint8((raw & 0x00F0) >> 4),
		Byte:   uint8(raw & 0x00FF),
		Nib:    uint8(raw & 0x000F),
		Addr:   uint16(raw & 0x0FFF),
	}

	for _, o := range Opcodes {
		if i.Raw&o.Mask == o.Pattern {
			i.Valid = true
			i.Opcode = &o
			break
		}
	}

	return i
}

func (i *Instr) String() string {
	ret := "INV  (#raw)"
	rep := map[string]string{
		"#addr": fmt.Sprintf("0x%04X", i.Addr),
		"#byte": fmt.Sprintf("0x%02X", i.Byte),
		"#nib":  fmt.Sprintf("0x%01X", i.Nib),
		"#vx":   fmt.Sprintf("V%01X", i.X),
		"#vy":   fmt.Sprintf("V%01X", i.Y),
		"#raw":  fmt.Sprintf("%04X", i.Raw),
	}

	if i.Valid {
		ret = i.Opcode.Format
	}

	for key, val := range rep {
		ret = strings.Replace(ret, key, val, -1)
	}

	if i.Valid && !i.Opcode.Implemented {
		ret = fmt.Sprintf("%-20s (not implemented yet)", ret)
	}

	return ret
}
