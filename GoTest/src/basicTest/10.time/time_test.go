package time

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	format := time.Now().Format("2006-01-02")
	local, _ := time.LoadLocation("Local")
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", format+"_23:59:59", local)
	fmt.Println(end, end.Unix())
	fmt.Println(format)
	var a uint64
	a = 1231
	fmt.Println(string(a))
}
