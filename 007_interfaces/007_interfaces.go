package main

import (
    "fmt"
    "math"
)

// 定义一个接口 `Shape`
type Shape interface {
    area() float64
}

// 定义一个矩形结构体
type Rectangle struct {
    width, height float64
}

// 实现 `Shape` 接口中的 `area` 方法
func (r Rectangle) area() float64 {
    return r.width * r.height
}

// 定义一个圆形结构体
type Circle struct {
    radius float64
}

// 实现 `Shape` 接口中的 `area` 方法
func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func main() {
    // 创建一个矩形
    rectangle := Rectangle{width: 3, height: 4}
    // 创建一个圆形
    circle := Circle{radius: 5}

    // 调用接口中的方法来计算不同形状的面积
    fmt.Println("矩形的面积:", getArea(rectangle)) // 输出：矩形的面积: 12
    fmt.Println("圆形的面积:", getArea(circle))     // 输出：圆形的面积: 78.53981633974483
}

// 接收一个 `Shape` 类型的参数，并调用其 `area` 方法
func getArea(shape Shape) float64 {
    return shape.area()
}