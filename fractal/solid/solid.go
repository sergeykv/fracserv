package solid

import (
	"github.com/sergeykv/fracserv/fractal"
	"fmt"
	"image"
	"image/color"
	"strconv"
)

type solid struct {
	*image.Uniform
	bounds image.Rectangle
}

func newSolidFromParams(p fractal.Params) (fractal.Fractal, error) {
	var r, g, b byte
	if _, err := fmt.Sscanf(p.Get("c"), "%2x%2x%2x", &r, &g, &b); err != nil {
		return nil, fmt.Errorf("Unable to parse colour '%s': %s", p.Get("c"), err)
	}
	w, err := strconv.Atoi(p.Get("w"))
	if err != nil {
		return nil, fmt.Errorf("Unable to parse width '%s': %s", p.Get("w"), err)
	}
	h, err := strconv.Atoi(p.Get("h"))
	if err != nil {
		return nil, fmt.Errorf("Unable to parse height '%s': %s", p.Get("h"), err)
	}
	return newSolid(w, h, color.RGBA{r, g, b, 255}), nil
}

func newSolid(w, h int, c color.Color) fractal.Fractal {
	return &solid{image.NewUniform(c), image.Rect(0, 0, w, h)}
}

func (s *solid) Bounds() image.Rectangle {
	return s.bounds
}

func init() {
	fractal.RegisterFractal("solid", newSolidFromParams)
}
