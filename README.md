# 多边图形碰撞检测

1.下载并使用
~~~shell
$ go install github.com/SkywardSky/polygon-collision@latest
~~~

2.使用
~~~
package main

import (
	"fmt"
	"polygoncollision"
)

func main(){
	var (
		polygon1 polygoncollision.Polygon
		polygon2 polygoncollision.Polygon
	)

	// 不交叉也不包含
	polygon1.Points = append(polygon1.Points, polygoncollision.Point{0, 0}, polygoncollision.Point{0, 3}, polygoncollision.Point{3, 3}, polygoncollision.Point{3, 0})
	polygon2.Points = append(polygon2.Points, polygoncollision.Point{4, 4}, polygoncollision.Point{4, 6}, polygoncollision.Point{6, 6}, polygoncollision.Point{6, 4})
	fmt.Println("不交叉也不包含：",polygoncollision.Detect(polygon1,polygon2))

	polygon1.Points=make([]polygoncollision.Point,0,10)
	polygon2.Points=make([]polygoncollision.Point,0,10)
	// 交叉
	polygon1.Points = append(polygon1.Points, polygoncollision.Point{0, 0}, polygoncollision.Point{0, 5}, polygoncollision.Point{5, 4}, polygoncollision.Point{3, 0})
	polygon2.Points = append(polygon2.Points, polygoncollision.Point{4, 4}, polygoncollision.Point{4, 6}, polygoncollision.Point{6, 6}, polygoncollision.Point{6, 4})
	fmt.Println("交叉：",polygoncollision.Detect(polygon1,polygon2))

	polygon1.Points=make([]polygoncollision.Point,0,10)
	polygon2.Points=make([]polygoncollision.Point,0,10)
	// 包含
	polygon1.Points = append(polygon1.Points, polygoncollision.Point{0, 0}, polygoncollision.Point{0, 3}, polygoncollision.Point{3, 3}, polygoncollision.Point{3, 0})
	polygon2.Points = append(polygon2.Points, polygoncollision.Point{1, 1}, polygoncollision.Point{1, 2}, polygoncollision.Point{2, 2}, polygoncollision.Point{2, 0})
	fmt.Println("包含：",polygoncollision.Detect(polygon1,polygon2))

	polygon1.Points=make([]polygoncollision.Point,0,10)
	polygon2.Points=make([]polygoncollision.Point,0,10)
	// 复杂图形
	polygon1.Points = append(polygon1.Points, polygoncollision.Point{7, 0}, polygoncollision.Point{7, 200}, polygoncollision.Point{150, 100}, polygoncollision.Point{300, 320}, polygoncollision.Point{300, 0})
	polygon2.Points = append(polygon2.Points, polygoncollision.Point{80, 300}, polygoncollision.Point{70, 200}, polygoncollision.Point{150, 130}, polygoncollision.Point{180, 250})
	fmt.Println("复杂图形：",polygoncollision.Detect(polygon1,polygon2))
}
~~~