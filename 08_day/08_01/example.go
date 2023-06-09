package example

import (
	"fmt"
)

type Sumer interface {
	Sum() int
}

func f1(v Sumer) {
	fmt.Println(v.Sum())
}

type Struct01 struct {
	w int
	h int
}

func (s Struct01) Sum() int { return s.w + s.h }

type Struct02 struct {
	x int
	y int
}

func (s Struct02) Sum() int { return s.x + s.y }

// func main() {
// 	ttt := Struct01{w: 3, h: 4}
// 	rrr := Struct02{x: 6, y: 8}
// 	f1(ttt) // 7
// 	f1(rrr) // 14
// }
