package fractal

import (
	"testing"
)

func TestNames(t *testing.T) {
	RegisterFractal("frac!", func(Params) (Fractal, error) {return nil, nil})
	RegisterFractal("test", func(Params) (Fractal, error) {return nil, nil})
	a := Names()
	if len(a) != 2 {
		t.Errorf("Expected to have two fractals in the registry, got %v", len(a))
		return
	}
	if !(a[0] == "test" && a[1] == "frac!" || a[0] == "frac!" && a[1] == "test") {
		t.Errorf("Expected to have 'test' and 'frac!' in the registry, got %v", a)
	}
}

type testParams map[string]string

func (t testParams) Get(key string) string {
	return t[key]
}

func TestNewFractal(t *testing.T) {
	wasCalled := false
	RegisterFractal("frac!", func(Params) (Fractal, error) {
		wasCalled = true
		return nil, nil
	})
	NewFractal("frac!", testParams{})
	if !wasCalled {
		t.Errorf("The callback was not called")
	}
}
