package structs_stuff

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectanlges", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got '%g' want '%g'", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got '%g' want '%g'", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
	t.Run("example of table testing", func(t *testing.T) {
		// this anonymous struct thing makes me happy
		areaTests := []struct {
			name         string
			shape        Shape
			expectedArea float64
		}{
			// optionally named args (fields)
			{name: "Test Rectangle Area", shape: Rectangle{10.0, 10.0}, expectedArea: 100.0},
			{name: "Test Circle Area", shape: Circle{10.0}, expectedArea: 314.1592653589793},
		}

		for _, tt := range areaTests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.shape.Area()

				if got != tt.expectedArea {
					t.Errorf("%#v - got '%g' want '%g'", tt.shape, got, tt.expectedArea)
				}
			})
		}
	})
}
