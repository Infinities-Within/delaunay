package delaunay

import "math"

var eps = math.Nextafter(1, 2) - 1

var infinity = math.Inf(1)

func pseudoAngle(dx, dy float64) float64 {
	p := dx / (math.Abs(dx) + math.Abs(dy))
	if dy > 0 {
		p = (3 - p) / 4
	} else {
		p = (1 + p) / 4
	}
	return math.Max(0, math.Min(1-eps, p))
}

func area(a, b, c Point2d) float64 {
	return (b.Y()-a.Y())*(c.X()-b.X()) - (b.X()-a.X())*(c.Y()-b.Y())
}

func inCircle(a, b, c, p Point2d) bool {
	dx := a.X() - p.X()
	dy := a.Y() - p.Y()
	ex := b.X() - p.X()
	ey := b.Y() - p.Y()
	fx := c.X() - p.X()
	fy := c.Y() - p.Y()

	ap := dx*dx + dy*dy
	bp := ex*ex + ey*ey
	cp := fx*fx + fy*fy

	return dx*(ey*cp-bp*fy)-dy*(ex*cp-bp*fx)+ap*(ex*fy-ey*fx) < 0
}

func circumradius(a, b, c Point2d) float64 {
	dx := b.X() - a.X()
	dy := b.Y() - a.Y()
	ex := c.X() - a.X()
	ey := c.Y() - a.Y()

	bl := dx*dx + dy*dy
	cl := ex*ex + ey*ey
	d := dx*ey - dy*ex

	x := (ey*bl - dy*cl) * 0.5 / d
	y := (dx*cl - ex*bl) * 0.5 / d

	r := x*x + y*y

	if bl == 0 || cl == 0 || d == 0 || r == 0 {
		return infinity
	}

	return r
}

func circumcenter(a, b, c Point2d) Point2d {
	dx := b.X() - a.X()
	dy := b.Y() - a.Y()
	ex := c.X() - a.X()
	ey := c.Y() - a.Y()

	bl := dx*dx + dy*dy
	cl := ex*ex + ey*ey
	d := dx*ey - dy*ex

	x := a.X() + (ey*bl-dy*cl)*0.5/d
	y := a.Y() + (dx*cl-ex*bl)*0.5/d

	return point{x, y}
}

func polygonArea(Point2ds []Point2d) float64 {
	var result float64
	for i, p := range Point2ds {
		q := Point2ds[(i+1)%len(Point2ds)]
		result += (p.X() - q.X()) * (p.Y() + q.Y())
	}
	return result / 2
}

func polygonPerimeter(Point2ds []Point2d) float64 {
	if len(Point2ds) == 0 {
		return 0
	}
	var result float64
	q := Point2ds[len(Point2ds)-1]
	for _, p := range Point2ds {
		result += p.Distance(q)
		q = p
	}
	return result
}
