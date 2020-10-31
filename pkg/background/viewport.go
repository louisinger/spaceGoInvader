package background

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type viewport struct {
	x16 int
	y16 int
}

func (p *viewport) move(bgImage *ebiten.Image) {
	w, h := bgImage.Size()
	maxX16 := w * 16
	maxY16 := h * 16

	p.x16 += w / 128
	p.y16 += h / 128
	p.x16 %= maxX16
	p.y16 %= maxY16
}

func (p *viewport) position() (int, int) {
	return p.x16, p.y16
}
