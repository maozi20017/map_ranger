package main

import (
	"fmt"
	"math"

	"github.com/go-vgo/robotgo"
	gohook "github.com/robotn/gohook"
)

var (
	meter, baselength float64
	x1, y1, x2, y2    int
)

func main() {
	fmt.Println("輸入比例尺")
	fmt.Scanln(&meter)
	fmt.Println("設定一格長度")
	baselength = firststep()
	for {
		getrange()
	}
}

func firststep() float64 {
	fmt.Println("按下滑鼠中鍵以設定初始起點")
	x1, y1 = togetmousepos()
	fmt.Println(("按下滑鼠中鍵以設定初始終點"))
	x2, y2 = togetmousepos()
	return math.Hypot(math.Abs(float64(x2-x1)), math.Abs(float64(y2-y1)))
}

func getrange() {
	fmt.Println("按下滑鼠中鍵以設定起點")
	x1, y1 = togetmousepos()
	fmt.Println(("按下滑鼠中鍵以設定終點"))
	x2, y2 = togetmousepos()
	therange := int(math.Hypot(math.Abs(float64(x2-x1)), math.Abs(float64(y2-y1))) / baselength * meter)
	fmt.Printf("距離:%d\n", therange)
}

func togetmousepos() (x int, y int) {
	ok := gohook.AddMouse("center") //原本是robotgo.AddEvent 已改成 gohook.AddEvent
	if ok {
		x, y = robotgo.GetMousePos()
	}
	return x, y
}
