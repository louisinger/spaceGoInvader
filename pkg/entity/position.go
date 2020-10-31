package entity

// Position describing a position on the screen
type Position struct {
	X float64
	Y float64
	maxX float64
	maxY float64
}

func (p *Position) translate(tX float64, tY float64) {
	p.X += tX
	p.Y += tY

	if p.X > p.maxX {
		p.X = p.maxX
	}

	if p.X < 0 {
		p.X = 0
	}

	if p.Y > p.maxY {
		p.Y = p.maxY
	}

	if p.Y < 0 {
		p.Y = 0
	}
}
