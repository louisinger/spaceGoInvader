package background

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Background represent the background image.
type Background struct {
	image *ebiten.Image
	viewport viewport
}

// NewBackground is the build function for the Background type
func NewBackground(image *ebiten.Image) *Background {
	return &Background{
		image: image,
		viewport: viewport{
			x16: 0,
			y16: 0,
		},
	}
}

// Draw function for background image
func (b *Background) Draw(screen *ebiten.Image) {
	const repeat = 3

	x16, y16 := b.viewport.position()
	offsetX, offsetY := float64(-x16)/16, float64(-y16)/16
	w, h := b.image.Size()

	for j := 0; j < repeat; j++ {
		for i := 0; i < repeat; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(w*i), float64(h*j))
			op.GeoM.Translate(offsetX, offsetY)
			screen.DrawImage(b.image, op)
		}
	}

}

// Update function for background image
func (b *Background) Update() {
	b.viewport.move(b.image)
}