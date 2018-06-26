// statisics
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"path/filepath"
	"runtime"
)

func init() {
	saneLength = makeBoundedIntFunc(1, 4096)
	saneRadius = makeBoundedIntFunc(1, 1024)
	saneSides = makeBoundedIntFunc(3, 60)
}

func makeBoundedIntFunc(minimum, maximim int) func(int) int {
	return func(x int) int {
		valid := x
		switch {
		case x < minimum:
			valid = minimum
		case x > maximim:
			valid = maximim
		}
		if valid != x {
			log.Printf("%s(): replaced %d width %d\n", caller(1), x, valid)
		}
		return valid
	}
}

func caller(steps int) string {
	/*
		runtime.Caller函数返回当前被调用函数的信息，并且不是在当前goroutine中返回。
		传入的参数int表示往后回退多远（即多少层函数）。如果为0，那么只查看当前函数信息。
		如果传入的参数为1，那么即查看该函数调用者信息。
		runtime.Caller 返回调用者，文件名，行数，是否取得了信息
	*/

	name := "?"
	if pc, _, _, ok := runtime.Caller(steps + 1); ok {
		name = filepath.Base(runtime.FuncForPC(pc).Name())
	}
	return name
}

func FilledImage(width, height int, fill color.Color) draw.Image {
	if fill == nil { // 默认将空的颜色设置为黑色
		fill = color.Black
	}
	width = saneLength(width)
	height = saneLength(height)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform(fill), image.ZP, draw.Src)
	return img
}

func DrawShapes(img draw.Image,x,y int,shapes ..Shaper)

func main() {
	author1 := Author1{Person{"Mr", []string{"Robert", "Louis", "Balfour"}, "Stevenson"}, []string{"Kidnapped", "Treasure IsLand"}, 1850}
	fmt.Println(author1)
	author1.Names.Title = "Fu"
	author1.Names.Forenames = []string{"zhangSan", "Lisi", "xiaoMing"}
	author1.Names.Surname = "Well"
	author1.YearBorn += 20
	fmt.Println(author1)
}
