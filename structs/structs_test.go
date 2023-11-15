package structs

import (
	"testing"
)

func TestPerimiter(t *testing.T) {
	r := Rectangle{2.0, 2.0}
	got := Perimiter(r)
	hasArea := 8.0

	if got != hasArea {
		t.Errorf("want %.2f got %.2f", hasArea, got)
	}
}

func TestArea(t *testing.T) {

	// checkArea := func(t testing.TB, shape Shape, hasArea float64) {
	// 	t.Helper()
	// 	got := shape.Area()

	// 	if got != hasArea {
	// 		t.Errorf("hasArea %g got %g", want, got)
	// 	}
	// }
	// t.Run("rectangle", func(t *testing.T) {
	// 	r := Rectangle{2.0, 2.0}
	// 	hasArea := 4.0
	// 	checkArea(t, r, hasArea)
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	c := Circle{10}
	// 	hasArea := 314.1592653589793
	// 	checkArea(t, c, hasArea)
	// })

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 2.0, Height: 2.0}, hasArea: 4.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12.0, Height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v want %g got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}
