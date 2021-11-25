package polygoncollision

// Point 点
type Point struct {
	X, Y float64
}

// sub 向量的减
func (sp Point) sub(p Point) Point {
	return Point{sp.X - p.X, sp.Y - p.Y}
}

// mul 向量的乘
func (sp Point) mul(val float64) Point {
	return Point{sp.X * val, sp.Y * val}
}

// add 向量的加
func (sp Point) add(p Point) Point {
	return Point{sp.X + p.X, sp.Y + p.Y}
}

// cross 向量的叉乘
func (sp Point) cross(p Point) float64 {
	return sp.X*p.Y - sp.Y*p.X
}
