package practice_2

import (
	"bufio"
	"fmt"
	"os"
)

func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			//直接调用
			fileName(countLines(f, counts), arg)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) bool { //使计算重复行的函数返回一个bool
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			//返回true存在重复行
			return true
		}
	}
	//返回false不存在重复行
	return false
}

func fileName(flag bool, arg string) {
	if flag {
		fmt.Printf("This file is : %s\n", arg)
	}
}
