package main

import "fmt"

func main() {
	s := make(chan int,1)
	s <- 5
	fmt.Println(<-s)
}
