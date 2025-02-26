package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"go_learn/src/day0226/mathutils"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(mathutils.Minus(10, 20))
	fmt.Println(cmp.Diff("Hello", "World"))
}
