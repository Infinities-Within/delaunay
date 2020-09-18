## Delaunay Triangulation
Fast Delaunay triangulation implemented in Go.

This module has been derived from [Michael Fogelman's port](https://github.com/fogleman/delaunay) of the [Mapbox's Delaunator project](https://github.com/mapbox/delaunator) (JavaScript).

The only modifications from Mr. Fogelman's code are that the project has been updated to be a Go module, and all methods now accept and return the `Point2d` interface, to allow easier interoperation between this library and code that does not want to use this library's internal representation of 2d points.

### Installation

    $ go get -u github.com/Infinities-Within/delaunay

### Documentation

https://godoc.org/github.com/Infinities-Within/delaunay

See https://mapbox.github.io/delaunator/ for more information about the `Triangles` and `Halfedges` data structures.

### Usage

```go
var points []delaunay.Point2d
// populate points...
triangulation, err := delaunay.Triangulate(points)
// handle err...
// use triangulation.Triangles, triangulation.Halfedges
```

### Performance

2.6 GHz Intel Core i7

| # of Points | Time |
| ---: | ---: |
| 10 | 1.559µs |
| 100 | 37.645µs |
| 1,000 | 485.625µs |
| 10,000 | 5.552ms |
| 100,000 | 79.895ms |
| 1,000,000 | 1.272s |
| 10,000,000 | 23.481s |

![Example](https://i.imgur.com/xhfW1EV.png)
