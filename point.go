package delaunay

import "math"

type point struct {
	x, y float64
}

func NewPoint(x, y float64) Point2d {
	return point{x, y}
}

type Point2d interface {
	X() float64
	Y() float64
	SquaredDistance(b Point2d) float64
	Distance(b Point2d) float64
	Sub(b Point2d) Point2d
}

func (a point) X() float64 {
	return a.x
}

func (a point) Y() float64 {
	return a.y
}

func (a point) SquaredDistance(b Point2d) float64 {
	dx := a.X() - b.X()
	dy := a.Y() - b.Y()
	return dx*dx + dy*dy
}

func (a point) Distance(b Point2d) float64 {
	return math.Hypot(a.X()-b.X(), a.Y()-b.Y())
}

func (a point) Sub(b Point2d) Point2d {
	return point{a.X() - b.X(), a.Y() - b.Y()}
}
