package main

import (
	"fmt"
	"strconv"
)

type (
	byte int8
	_w   string
)

func main() {
	var a int = 65
	b := strconv.Itoa(a)
	c, _ := strconv.Atoi(b)
	fmt.Println(b)
	fmt.Println(c)
}
