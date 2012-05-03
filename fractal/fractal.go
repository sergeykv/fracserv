package fractal

import (
	"fmt"
	"image"
)

type Fractal interface {
	image.Image
}

type Params interface {
	Get(string) string
}

var registry = make(map[string]func(Params) (Fractal, error))

func RegisterFractal(name string, factory func(Params) (Fractal, error)) {
	registry[name] = factory
}

func NewFractal(name string, params Params) (Fractal, error) {
	factory, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("Unrecognized fractal name: '%s'", name)
	}
	return factory(params)
}