package main

import (
	"fmt"
	"strconv"
)

type Test interface {
	Go() string
}

type One struct {
	first  int
	second int
}

type Two struct {
	first  string
	second string
}

func (o One) Go() string {
	a := o.first + o.second
	return strconv.Itoa(int(a))
}

func (s Two) Go() string {
	a := s.first + s.second
	return a
}

func PrintGo(g Test) {
	fmt.Println(g.Go())
}

func main() {
	first := One{1, 2}
	second := Two{"one", "two"}

	PrintGo(first)
	PrintGo(second)
}
