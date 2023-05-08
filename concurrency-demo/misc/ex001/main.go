package main

import "fmt"

type VisitFn func(value int) int

func square(v int) int {
	return v * v
}
func cube(v int) int {
	return v * v * v
}

func execute(ch chan VisitFn) {

	for i := 0; i < 20; i++ {
		if i < 10 {
			ch <- cube
		} else {
			ch <- square
		}
	}
}

func main() {
	ch := make(chan VisitFn)
	go execute(ch)

	for i := 0; i < 20; i++ {
		ex := <-ch
		fmt.Println(ex(i))
	}
}
