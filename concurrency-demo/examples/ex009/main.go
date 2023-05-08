package main

import "fmt"

func seq(start byte, end byte) <-chan byte { // Generator
	ch := make(chan byte)
	go func() {
		for {
			for i := start; i <= end; i++ {
				ch <- i
			}
		}
	}()
	return ch
}

func merge(inputs ...<-chan byte) <-chan byte { // AKA FanIn
	c := make(chan byte)
	for _, input := range inputs {
		ch := input
		go func() {
			for {
				c <- <-ch
			}
		}()
	}
	return c
}

func main() {
	ch := merge(seq('A', 'Z'), seq('0', '9'), seq('!', '~')) // HL
	for i := 0; i < 400; i++ {
		fmt.Printf("%c", <-ch)
	}
	fmt.Printf("\nLeaving ...\n")
}
