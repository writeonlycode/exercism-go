// Package triangle defines a function to determine whether a figure is a
// equilateral, isosceles or scalene triangle.
package triangle

type Kind = int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides determines whether a figure is a equilateral, isosceles or
// scalene triangle based on its sides.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case a <= 0 || b <= 0 || c <= 0:
		return NaT
	case a+b < c || b+c < a || c+a < b:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b:
		return Iso
	case b == c:
		return Iso
	case c == a:
		return Iso
	}

	return Sca
}
