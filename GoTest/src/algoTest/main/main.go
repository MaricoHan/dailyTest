package main

import (
	"fmt"
	"time"
)

func doubleForAdd() {
	before := time.Now().UnixNano()
	for i := 0; i < 1e3; i++ {
		for j := 0; j < 1e6; j++ {
		}
	}
	after := time.Now().UnixNano()
	fmt.Println("外层次数少，内层次数多", after-before) // 577001181

	before = time.Now().UnixNano()
	for i := 0; i < 1e6; i++ {
		for j := 0; j < 1e3; j++ {
		}
	}
	after = time.Now().UnixNano()

	fmt.Println("外层次数多，内层次数少", after-before) // 290771085
}
func main() {
	before := time.Now().UnixNano()
	for i := 0; i < 1e3; i++ {
		for j := 0; j < 1e6; j++ {
		}
	}
	after := time.Now().UnixNano()
	fmt.Println("外层次数少，内层次数多", after-before) // 577001181

	before = time.Now().UnixNano()
	for i := 0; i < 1e6; i++ {
		for j := 0; j < 1e3; j++ {
		}
	}
	after = time.Now().UnixNano()

	fmt.Println("外层次数多，内层次数少", after-before) // 290771085
}
