package geometry

import (
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("%s got %.2f want %.2f", reflect.TypeOf(shape).Name(), got, want)
		}
	}

	shapesToTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{12.0}, 452.3893421169302},
	}

	for _, testCase := range shapesToTest {
		checkArea(t, testCase.shape, testCase.want)
	}
}
