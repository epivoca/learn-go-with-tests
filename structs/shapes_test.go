package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Parallel() // marks TestPerimeter as capable of running in parallel with other tests
	perimeterTests := map[string]struct {
		shape        Shape
		hasPerimeter float64
	}{
		"Rectangle": {
			shape:        Rectangle{Width: 12, Height: 6},
			hasPerimeter: 36.0,
		},
		"Circle": {
			shape:        Circle{Radius: 10.0},
			hasPerimeter: 62.83185307179586,
		},
		"Triangle": {
			shape:        Triangle{Base: 10.0, Height: 7.0, SideA: 4.0, SideB: 8.0},
			hasPerimeter: 22.0,
		},
	}

	for name, tt := range perimeterTests {
		tt := tt // NOTE: /wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(name, func(t *testing.T) {
			if got, expected := tt.shape.Perimeter(), tt.hasPerimeter; got != expected {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	t.Parallel() // marks TestArea as capable of running in parallel with other tests

	areaTests := map[string]struct {
		shape   Shape
		hasArea float64
	}{
		"Rectangle": {
			shape:   Rectangle{Width: 12, Height: 6},
			hasArea: 72.0,
		},
		"Circle": {
			shape:   Circle{Radius: 10.0},
			hasArea: 314.1592653589793,
		},
		"Triangle": {
			shape:   Triangle{Base: 12, Height: 6},
			hasArea: 36,
		},
	}

	for name, tt := range areaTests {
		tt := tt // NOTE: /wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got, expected := tt.shape.Area(), tt.hasArea; got != expected {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
