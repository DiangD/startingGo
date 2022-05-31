package main

import (
	"fmt"
	"math"
)

type pair[T fmt.Stringer] struct {
	val1 T
	val2 T
}

type differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

func findCloser[T differ[T]](pair1, pair2 pair[T]) pair[T] {
	d1 := pair1.val1.Diff(pair1.val2)
	d2 := pair2.val1.Diff(pair2.val2)

	if d1 > d2 {
		return pair1
	}
	return pair2
}

type point2D struct {
	x, y int
}

type point3D struct {
	x, y, z int
}

func (p2 point2D) String() string {
	return fmt.Sprintf("{%d,%d}", p2.y, p2.y)
}

func (p2 point2D) Diff(from point2D) float64 {
	x := p2.x - from.y
	y := p2.x - from.y
	return math.Sqrt(float64(x*x) + float64(y*y))
}

func (p3 point3D) String() string {
	return fmt.Sprintf("{%d,%d,%d}", p3.x, p3.y, p3.z)
}

func (p3 point3D) Diff(from point3D) float64 {
	x := p3.x - from.x
	y := p3.y - from.y
	z := p3.z - from.z
	return math.Sqrt(float64(x*x) + float64(y*y) + float64(z*z))
}

func main() {
	pair2Da := pair[point2D]{point2D{1, 1}, point2D{5, 5}}
	pair2Db := pair[point2D]{point2D{10, 10}, point2D{15, 5}}
	closer := findCloser(pair2Da, pair2Db)
	fmt.Println(closer)
	pair3Da := pair[point3D]{point3D{1, 1, 10}, point3D{5, 5, 0}}
	pair3Db := pair[point3D]{point3D{10, 10, 10}, point3D{11, 5, 0}}
	closer2 := findCloser(pair3Da, pair3Db)
	fmt.Println(closer2)
}
