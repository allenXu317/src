package practice_1

import (
	"fmt"
	"os"
)

func Echo2() {

	for k, val := range os.Args {
		fmt.Println(k, val)
	}
}
