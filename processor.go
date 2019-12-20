package main

import (
	"gocv.io/x/gocv"
	"math"
	"sort"
)

func getLineStartingPointsForImage(path string) []int {
	edges := gocv.NewMat()
	lines := gocv.NewMat()
	img := gocv.IMRead(path, gocv.IMReadColor)
	gocv.Canny(img, &edges, 30, 90)
	gocv.HoughLinesPWithParams(edges, &lines, 1, math.Pi/180.0, 6, 200, 6)
	allPoints := []*Point{}
	points := []*Point{}
	// try to get points without duplicating ones
	for j := 0; j < lines.Rows(); j++ {
		vec := lines.GetVeciAt(j, 0)
		p1 := &Point{X: int(vec[0]), Y: int(vec[1])}
		allPoints = append(allPoints, p1)
		if hasClosePoint(p1, allPoints) {
			continue
		}
		points = append(points, p1)
	}
	// sort from top do bottom
	sort.Slice(points, func(i, j int) bool {
		return points[i].Y < points[j].Y
	})
	// get only values
	res := []int{}
	for _, point := range points {
		res = append(res, point.Y)
	}
	return res
}

func hasClosePoint(p *Point, points []*Point) bool {
	for _, point := range points {
		dis := p.Distance(*point)
		if dis > 0 && dis < 10 {
			return true
		}
	}
	return false
}

type Point struct {
	X int
	Y int
}

// New returns a Point based on X and Y positions on a graph.
func New(x int, y int) Point {
	return Point{x, y}
}

// Distance finds the length of the hypotenuse between two points.
// Forumula is the square root of (x2 - x1)^2 + (y2 - y1)^2
func (p Point) Distance(p2 Point) float64 {
	first := math.Pow(float64(p2.X-p.X), 2)
	second := math.Pow(float64(p2.Y-p.Y), 2)
	return math.Sqrt(first + second)
}
