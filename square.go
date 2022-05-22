package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (sq Square) End() Point {
	return Point{x: sq.start.x + int(sq.a),
		y: sq.start.y + int(sq.a)}
}

func (sq Square) Area() uint {
	return sq.a * sq.a
}

func (sq Square) Perimeter() uint {
	return sq.a * 4
}

func NewSquare(startX int, startY int, a uint) Square {
	return Square{start: Point{x: startX, y: startY}, a: a}
}
