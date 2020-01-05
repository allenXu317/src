package main

import (
	"bufio"
	"ch1/practice_1"
	"fmt"
	"os"
)

//书本例程
func readTest() {
	counts := make(map[string]int)      //键值对存放重复行  键为字符串，值为整型
	input := bufio.NewScanner(os.Stdin) //声明一个新的Scanner类型

	for input.Scan() && input.Text() != "!" { //以"!"作为结束标志
		counts[input.Text()]++ //有重复行就在数量上+1

		//等价于：
		//line := input.Text()
		//counts[line] = counts[line]+1
	}

	for line, n := range counts { //遍历map进行输出
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	fmt.Println("This is ch1 ")
	practice_1.Echo1()
	practice_1.Echo2()
	practice_1.Test1()
	practice_1.Test2()

	readTest()
}
