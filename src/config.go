/*
 * Copyright (c) 2015 Łukasz S.
 * Distributed under the terms of GPL-2 License.
 */

package yac8vm

const (
	AppVersion = "0.0.1"
	AppTitle   = "Yet Another Chip-8 Virtual Machine"
	AppShort   = "yac8vm"
	MemSize    = 4096
	StackSize  = 16
	ProgStart  = 0x200
	RegCount   = 16
	KeyCount   = 16
	ScreenW    = 64
	ScreenH    = 32
	ClockFreq  = 60
	ClockMul   = 30
	WindowW    = ScreenW * 10
	WindowH    = ScreenH * 10
)

var KeyMap = map[string]int8{
	"1": 0x1,
	"2": 0x2,
	"3": 0x3,
	"4": 0xC,
	"Q": 0x4,
	"W": 0x5,
	"E": 0x6,
	"R": 0xD,
	"A": 0x7,
	"S": 0x8,
	"D": 0x9,
	"F": 0xE,
	"Z": 0xA,
	"X": 0x0,
	"C": 0xB,
	"V": 0xF,
}

var InitialMemory = []byte{
	0xF0, // ****
	0x90, // *  *
	0x90, // *  *
	0x90, // *  *
	0xF0, // ****

	0x20, //   *
	0x60, //  **
	0x20, //   *
	0x20, //   *
	0x70, //  ***

	0xF0, // ****
	0x10, //    *
	0xF0, // ****
	0x80, // *
	0xF0, // ****

	0xF0, // ****
	0x10, //    *
	0xF0, // ****
	0x10, //    *
	0xF0, // ****

	0x90, // *  *
	0x90, // *  *
	0xF0, // ****
	0x10, //    *
	0x10, //    *

	0xF0, // ****
	0x80, // *
	0xF0, // ****
	0x10, //    *
	0xF0, // ****

	0xF0, // ****
	0x80, // *
	0xF0, // ****
	0x90, // *  *
	0xF0, // ****

	0xF0, // ****
	0x10, //    *
	0x20, //   *
	0x40, //  *
	0x40, //  *

	0xF0, // ****
	0x90, // *  *
	0xF0, // ****
	0x90, // *  *
	0xF0, // ****

	0xF0, // ****
	0x90, // *  *
	0xF0, // ****
	0x10, //    *
	0xF0, // ****

	0xF0, // ****
	0x90, // *  *
	0xF0, // ****
	0x90, // *  *
	0x90, // *  *

	0xE0, // ***
	0x90, // *  *
	0xE0, // ***
	0x90, // *  *
	0xE0, // ***

	0xF0, // ****
	0x80, // *
	0x80, // *
	0x80, // *
	0xF0, // ****

	0xE0, // ***
	0x90, // *  *
	0x90, // *  *
	0x90, // *  *
	0xE0, // ***

	0xF0, // ****
	0x80, // *
	0xF0, // ****
	0x80, // *
	0xF0, // ****

	0xF0, // ****
	0x80, // *
	0xF0, // ****
	0x80, // *
	0x80, // *
}