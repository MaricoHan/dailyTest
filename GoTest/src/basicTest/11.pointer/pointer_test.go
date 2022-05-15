package pointer

import (
	"fmt"
	"testing"
)

// go所有的传递都是值传递，没有
func TestPointer(t *testing.T) {
	s := "123123"
	fmt.Printf("main s: %v \n", s)
	fmt.Printf("main &s: %v \n", &s)
	hello(&s)
	fmt.Println(s)
}
func hello(s *string) { // 将指针型的入参拷贝一份，存入新的内存空间s
	fmt.Printf("hello s: %v \n", s)
	fmt.Printf("hello &s: %v \n", &s)
	*s = "456456"
}
