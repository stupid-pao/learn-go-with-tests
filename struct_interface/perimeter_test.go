package main

import (
	"math"
	"testing"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

//func TestArea(t *testing.T) {
//checkArea := func(t *testing.T, shape Shape, want float64) {
//t.Helper()
//got := shape.Area()
//if got != want {
//t.Errorf("got %.2f want %.2f", got, want)
//}
//}

//t.Run("rectangle", func(t *testing.T) {
//rectangle := Rectangle{12, 6}
//checkArea(t, rectangle, 72.0)
//})

//t.Run("circles", func(t *testing.T) {
//circle := Circle{10}
//checkArea(t, circle, math.Pi*100)
//})
//}

func TestArea(t *testing.T) {
	//列表驱动测试 -- 创建一系列相同的测试方式的测试用例

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: math.Pi * 100},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		// can use go test -run TestArea/Rectangle 测试指定用例
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})
	}
}

/*
- 声明结构体以创建我们自己的类型， 让我们把数据集合在一起并达到简化代码的目的
- 声明接口， 这样可以定义适合不同参数类型的函数（参数多态）
- 在自己的数据类型中添加方法以实现接口
- 列表驱动让断言更加清晰，这样可以使测试文件更易于扩展维护
接口作用：负责把从系统的其他不分隐藏起来
*/
