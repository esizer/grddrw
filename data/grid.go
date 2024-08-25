package data

type Pixel struct {
	cell uint8
	R    uint8
	G    uint8
	B    uint8
}

type Grid struct {
	Id     int
	Pixels [256]*Pixel
}

func (g *Grid) Paint(idx uint16, p Pixel) {
	g.Pixels[idx] = &Pixel{
		R: p.R,
		G: p.G,
		B: p.B,
	}
}

func NewGrid() Grid {
	pxs := [256]*Pixel{}

	for i := range pxs {
		pxs[i] = &Pixel{
			R: 255,
			G: 255,
			B: 255,
		}
	}

	return Grid{
		Id:     1,
		Pixels: pxs,
	}
}
