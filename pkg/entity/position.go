package entity

// Position describing a position on the screen
type Position struct {
	X float64
	Y float64
	maxX float64
	maxY float64
}

func (p *Position) translate(tX float64, tY float64) bool {
	p.X += tX
	p.Y += tY

	if p.X > p.maxX {
		p.X = p.maxX
		return true
	}

	if p.X < 0 {
		p.X = 0
		return true
	}

	if p.Y > p.maxY {
		p.Y = p.maxY
		return true
	}

	if p.Y < 0 {
		p.Y = 0
		return true
	}

	return false
}
