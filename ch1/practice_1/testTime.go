package practice_1

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Test1() {
	testTime1()
}
func Test2() {
	testTime2()
}

func testTime1() {
	start := time.Now()

	var res string
	s := ""

	for _, v := range os.Args {
		res += v
		res += s
	}
	fmt.Println(res)
	fmt.Printf("%.60fs elapse\n", time.Since(start).Seconds())
}

func testTime2() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("%.60fs elapse\n", time.Since(start).Seconds())
}
