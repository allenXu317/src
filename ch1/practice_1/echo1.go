package practice_1

import (
	"fmt"
	"os"
)

func Echo1() {
	fmt.Println("This is practice 1.1")
	s := os.Args[0]
	fmt.Println(s)
}
