package clockface

import (
	"math"
	"time"
)

// Um ponto representa uma coordenada cartesiana com dois eixos (bidimensional)
type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
