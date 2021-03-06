/*
 * Copyright (c) 2015 Łukasz S.
 * Distributed under the terms of GPL-2 License.
 */

package src

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func runApplication(romFileName string) {
	fp, err := os.Open(romFileName)
	handleError(err)
	defer fp.Close()

	g, err := NewGUI()
	handleError(err)
	defer g.Destroy()

	m, err := NewMachine(g, fp)
	handleError(err)
	m.Run()
}

func disApplication(romFileName string) {
	file, err := os.Open(romFileName)
	handleError(err)
	defer file.Close()

	buf := make([]byte, 2)
	pos := 0x200

	for {
		n, err := file.Read(buf)
		if err == io.EOF || n != len(buf) {
			break
		}
		handleError(err)

		raw := uint16(buf[0])<<8 | uint16(buf[1])
		i := NewInstr(raw)

		if !i.Valid {
			// break
		}

		fmt.Printf("%04X:      %04X      %v\n", pos, raw, i)

		pos += len(buf)
	}
}

func Main() {
	desc := AppShort + " " + AppVersion + " - " + AppTitle
	app := kingpin.New(AppShort, desc)
	app.Version(AppVersion)

	runCmd := app.Command("run", "Run an application")
	runCmdRomFile := runCmd.Arg("romfile", "ROM file to execute").Required().String()

	disCmd := app.Command("dis", "Disassemble an application")
	disCmdRomFile := disCmd.Arg("romfile", "ROM file to disassemble").Required().String()

	cmd := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch cmd {
	case runCmd.FullCommand():
		runApplication(*runCmdRomFile)
	case disCmd.FullCommand():
		disApplication(*disCmdRomFile)
	default:
		app.FatalUsage("")
	}
}
