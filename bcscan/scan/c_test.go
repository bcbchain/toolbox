package scan

import (
	"fmt"
	"sort"
	"testing"
)

func TestName(t *testing.T) {

	//c := make(chan int, 10)
	//for i := 0; i < 200; i++ {
	//	go func(ii int) {
	//		c <- ii
	//	}(i)
	//}
	//
	//r := make([]int, 200)
	//for {
	//	a := <-c
	//	r[a] = a
	//	if len(r) == 200 {
	//		return
	//	}
	//}

	type A struct {
		A int
		B int
	}

	sl := []A{
		{1, 5}, {11, 16}, {6, 10},
	}

	sort.Slice(sl, func(i, j int) bool {
		if sl[i].A < sl[j].A {
			return true
		}
		return false
	})
	fmt.Println(sl)

	sl = sl[3:]
	fmt.Println(sl)
}
