package basictest

import (
	"fmt"
	"testing"
	"time"
)

func doubleForAdd() {
	before := time.Now().UnixNano()
	for i := 0; i < 1e3; i++ {
		for j := 0; j < 1e7; j++ {
		}
	}
	after := time.Now().UnixNano()
	fmt.Println("外层次数少，内层次数多", after-before) // 577001181

	before = time.Now().UnixNano()
	for i := 0; i < 1e7; i++ {
		for j := 0; j < 1e3; j++ {
		}
	}
	after = time.Now().UnixNano()

	fmt.Println("外层次数多，内层次数少", after-before) // 290771085
}

func doubleForSub() {
	before := time.Now().UnixNano()
	for i := 1e3; i > 0; i-- {
		for j := 1e7; j > 0; j-- {
		}
	}
	after := time.Now().UnixNano()
	fmt.Println("外层次数少，内层次数多", after-before) // 1138045495

	before = time.Now().UnixNano()
	for i := 1e7; i > 0; i-- {
		for j := 1e3; j > 0; j-- {
		}
	}
	after = time.Now().UnixNano()

	fmt.Println("外层次数多，内层次数少", after-before) // 1191986287
}

// TestFor
// @Description:
// 结论：（只针对go语言）
// 1.“外少内多”更耗时
// 2.“自减”比“自增”更耗时
// @param t
func TestFor(t *testing.T) {
	fmt.Println("双层for循环，自增")
	doubleForAdd()
	fmt.Println("双层for循环，自减")
	doubleForSub()
}
