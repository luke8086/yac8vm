package src

import (
	"github.com/veandco/go-sdl2/sdl"
)

type GUI struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
}

func NewGUI() (*GUI, error) {
	var err error

	window, err := sdl.CreateWindow(AppTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WindowW, WindowH, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		window.Destroy()
		return nil, err
	}

	renderer.SetDrawColor(0x00, 0x00, 0x00, 0xFF)
	renderer.Clear()
	renderer.Present()

	return &GUI{Window: window, Renderer: renderer}, nil
}

func (g *GUI) Destroy() {
	if g.Renderer != nil {
		g.Renderer.Destroy()
	}

	if g.Window != nil {
		g.Window.Destroy()
	}

	sdl.Quit()
}

func (g *GUI) Draw(m *Machine) {
	s := &m.ScreenBuf
	w := int32(WindowW / len(s))
	h := int32(WindowH / len(s[0]))
	rect := sdl.Rect{0, 0, w, h}

	for x := 0; x < len(s); x++ {
		for y := 0; y < len(s[x]); y++ {
			rect.X = int32(x) * rect.W
			rect.Y = int32(y) * rect.H
			c := uint8(0x00)
			if s[x][y] {
				c = 0xFF
			}
			g.Renderer.SetDrawColor(c, c, c, 0xFF)
			g.Renderer.FillRect(&rect)
		}
	}

	g.Renderer.Present()
}

func (g *GUI) UpdateTitle(m *Machine) {
	if m.Paused {
		g.Window.SetTitle(AppTitle + " (paused)")
	} else {
		g.Window.SetTitle(AppTitle)
	}
}

func (g *GUI) ProcessKey(m *Machine, code sdl.Keycode, down bool) {
	kname := sdl.GetKeyName(code)

	if code == sdl.K_p && !down {
		m.TogglePaused(!m.Paused)
		g.UpdateTitle(m)
		return
	}

	if code == sdl.K_ESCAPE && !down {
		m.RequestExit()
		return
	}

	if key, ok := KeyMap[kname]; ok {
		m.ToggleKey(key, down)
	}
}

func (g *GUI) ProcessEvent(m *Machine) {
	ev := sdl.PollEvent()
	if ev == nil {
		return
	}

	switch t := ev.(type) {
	case *sdl.QuitEvent:
		m.RequestExit()

	case *sdl.WindowEvent:
		m.Draw()

	case *sdl.KeyboardEvent:
		g.ProcessKey(m, t.Keysym.Sym, t.Type == sdl.KEYDOWN)
	}
}
