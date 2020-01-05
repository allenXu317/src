package main

import (
	"bufio"
	"ch1/practice_1"
	"ch1/practice_2"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//书本例程
func dup1() {
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

func dup2() {
	counts := make(map[string]int)
	file := os.Args[1:]
	fmt.Println(file)
	if len(file) == 0 { //如果命令行参数个数为0，则从标准输入中读取数据
		//此时函数功能同dup1
		countLines(os.Stdin, counts)
	} else {
		//从文件中读取数据并进行比较
		for _, arg := range file {
			f, err := os.Open(arg) //打开文件
			if err != nil {        //打开文件出错
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts) //各个文件中进行查找
			f.Close()             //关闭文件
		}
	}

	for line, n := range counts {
		if n > 1 {
			//重复行的数据内容
			//相同行各自的个数
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		//line := input.Text()
		//count[line]++
	}
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] { //读取命令行参数
		data, err := ioutil.ReadFile(filename) //与os.Open 打开文件不同，这个是读取一整块文件内容
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			// fmt.Println(line, counts[line])
			counts[line]++
		}
	}
	for line, n := range counts {
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

	practice_2.Dup2()

	// dup1()
	// dup2()
	// dup3()

}
