package solid

import (
	"image"
	"image/color"
	"testing"
)

func TestBounds(t *testing.T) {
	expected := image.Rect(0, 0, 10, 10)
	solid := newSolid(10, 10, color.RGBA{1, 1, 1, 1})
	actual := solid.Bounds()
	if !expected.Eq(actual) {
		t.Errorf("Incorrect bounds, expected %v, actual %v", expected, actual)
	}
}

func TestColour(t *testing.T) {
	expected := color.RGBA{1, 2, 3, 4}
	solid := newSolid(10, 10, expected)
	actual := solid.At(2, 2)
	if expected != actual {
		t.Errorf("Incorrect colour, expected %v, actual %v", expected, actual)
	}
}

type testParams map[string]string

func (t testParams) Get(key string) string {
	return t[key]
}

func TestInitFromParams(t *testing.T) {
	solid, _ := newSolidFromParams(testParams{"w": "10", "h": "10", "c": "aabbcc"})
	expectedBounds := image.Rect(0, 0, 10, 10)
	if !expectedBounds.Eq(solid.Bounds()) {
		t.Errorf("Incorrect bounds, expected %v, actual %v", expectedBounds, solid.Bounds())
	}
	expectedColour := color.RGBA{0xaa, 0xbb, 0xcc, 0xff}
	if expectedColour != solid.At(2, 2) {
		t.Errorf("Incorrect colour, expected %v, actual %v", expectedColour, solid.At(2, 2))
	}
}

func TestInitParseFail(t *testing.T) {
	solid, err := newSolidFromParams(testParams{"w": "abc", "h": "10", "c": "aabbcc"})
	if err == nil {
		t.Errorf("Expected an error, got a fractal %v", solid)
	}
}
