package container

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	a := map[string]int{
		"12": 3,
	}
	c := map[string]int{
		"12": 5,
	}
	// 引用类型，slice、map和chan
	// 引用，接收者是源的别名，接收者即时源
	b := a
	b["12"] = 4
	fmt.Println(a, b) //改变b即是改变a
	b = c
	fmt.Println(b)
	b["12"] = 6
	fmt.Println(a, b, c)

}
