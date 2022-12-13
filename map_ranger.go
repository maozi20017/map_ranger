package main

import (
	"fmt"
	"math"

	"github.com/go-vgo/robotgo"
	gohook "github.com/robotn/gohook"
)

var (
	meter, thebase float64
	x1, y1, x2, y2 int
)

func main() {
	fmt.Println("輸入比例尺")
	fmt.Scanln(&meter)
	fmt.Println("設定一格長度")
	thebase = thefirst()
	fmt.Printf("%d", int(thebase))
	for {
		getrange()
	}
}

func thefirst() float64 {
	fmt.Println("按下滑鼠中鍵以設定初始起點")
	ok := gohook.AddMouse("center") //原本是robotgo.AddEvent 已改成 gohook.AddEvent
	if ok {
		x1, y1 = robotgo.GetMousePos()
		fmt.Println(("按下滑鼠中鍵以設定初始終點"))
	}
	ok = gohook.AddMouse("center")
	if ok {
		x2, y2 = robotgo.GetMousePos()
	}
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func getrange() {

	fmt.Println("按下滑鼠中鍵以設定起點")
	ok := gohook.AddMouse("center") //原本是robotgo.AddEvent 已改成 gohook.AddEvent
	if ok {
		x1, y1 = robotgo.GetMousePos()
		fmt.Println(("按下滑鼠中鍵以設定終點"))
	}
	ok = gohook.AddMouse("center")
	if ok {
		x2, y2 = robotgo.GetMousePos()
		therange := int(math.Sqrt((math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))) / thebase * meter)
		fmt.Printf("%d\n", int(math.Sqrt((math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2)))))
		fmt.Printf("距離:%d\n", therange)
	}
}
