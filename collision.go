package polygoncollision

import "math"

// Detect 探测两个多边形是否碰撞
// polygon1 多边形1
// polygon2 多边形2
// bool true-碰撞，false-不碰撞
func Detect(polygon1, polygon2 Polygon) bool {
	// 探测两个多边形是否存在线段的碰撞
	if detectPolygonSegments(polygon1, polygon2) {
		return true
	}
	// 探测两个多边形是否存在包含的碰撞
	if detectPolygonDots(polygon1, polygon2) {
		return true
	}
	return false
}

// detectPolygonSegment 探测两个多边形的线段是否碰撞
// polygon1 多边形1
// polygon2 多边形2
// bool true-相交，false-不相交
func detectPolygonSegments(polygon1, polygon2 Polygon) bool {
	for i := 1; i < len(polygon1.Points); i++ {
		if detectPolygonSegment(Segment{polygon1.Points[i-1], polygon1.Points[i]}, polygon2) {
			return true
		}
	}

	return detectPolygonSegment(Segment{polygon1.Points[len(polygon1.Points)-1], polygon1.Points[0]}, polygon2)
}

// detectPolygonSegment 线段是否与多边形的所有线段碰撞
func detectPolygonSegment(segment Segment, polygon Polygon) bool {
	for i := 1; i < len(polygon.Points); i++ {
		if detectSegment(segment, Segment{polygon.Points[i-1], polygon.Points[i]}) {
			return true
		}
	}

	return detectSegment(segment, Segment{polygon.Points[len(polygon.Points)-1], polygon.Points[0]})
}

// detectSegment 探测两根线段是否碰撞
func detectSegment(seg1, seg2 Segment) bool {
	// 快速校验一下必不可相交的情况
	if math.Max(seg2.p2.X, seg2.p1.X) < math.Min(seg1.p2.X, seg1.p1.X) ||
		math.Max(seg2.p2.Y, seg2.p1.Y) < math.Min(seg1.p2.Y, seg1.p2.Y) ||
		math.Max(seg1.p2.X, seg1.p1.X) < math.Min(seg2.p2.X, seg2.p1.X) ||
		math.Max(seg1.p2.Y, seg1.p1.Y) < math.Min(seg2.p2.Y, seg2.p1.Y) {
		return false
	}

	// 线段1的向量
	vec1 := seg1.p2.sub(seg1.p1)
	// 线段1到线段2的两个端点的向量
	vec1_1 := seg2.p1.sub(seg1.p1)
	vec1_2 := seg2.p2.sub(seg1.p1)
	// 左右方向判断
	if !detectSegment2(vec1, vec1_1, vec1_2) {
		return false
	}

	vec2 := seg2.p2.sub(seg2.p1)
	vec2_1 := seg1.p1.sub(seg2.p1)
	vec2_2 := seg1.p2.sub(seg2.p1)

	if !detectSegment2(vec2, vec2_1, vec2_2) {
		return false
	}

	return true

}

// detectSegment2 探测向量v1、v2是否分别在向量v的左右2边
// bool true-是，false-不是
func detectSegment2(v, v1, v2 Point) bool {
	// 如果在左右两边，相乘是一个负数
	return v.cross(v1)*v.cross(v2) <= 0
}

// detectPolygonDots 探测两个多边形的点是否在另一个多边形内
// polygon1 多边形1
// polygon2 多边形2
// bool true-相交，false-不相交
func detectPolygonDots(polygon1, polygon2 Polygon) bool {
	for _, point := range polygon1.Points {
		if detectDot(point, polygon2) {
			return true
		}

	}
	for _, point := range polygon2.Points {
		if detectDot(point, polygon1) {
			return true
		}

	}
	return false
}

// detectDot 探测点是否在多边形的内部
// p 点
// polygon 多边形
// bool true-相交，false-不相交
func detectDot(p Point, polygon Polygon) bool {
	var (
		right int
	)

	for i := 1; i < len(polygon.Points); i++ {
		tmp := getIntersectionPoint(Segment{polygon.Points[i-1], polygon.Points[i]}, p)

		if nil == tmp {
			continue
		}

		if tmp.X > p.X {
			right++
		}
	}

	tmp := getIntersectionPoint(Segment{polygon.Points[len(polygon.Points)-1], polygon.Points[0]}, p)
	if nil != tmp && tmp.X > p.X {
		right++
	}

	return right%2 == 1
}

// getIntersectionPoint 获取水平线与线段的交点
// segment 线段
// point 点
// Point 交点
func getIntersectionPoint(segment Segment, point Point) *Point {
	if math.Min(segment.p1.Y, segment.p2.Y) >= point.Y ||
		math.Max(segment.p1.Y, segment.p2.Y) < point.Y {
		return nil
	}

	f := math.Abs((point.Y - segment.p1.Y) / (segment.p1.Y - segment.p2.Y))
	val := segment.p2.sub(segment.p1).mul(f).add(segment.p1)
	return &val
}
