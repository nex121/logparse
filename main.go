package main

import (
	"bufio"
	"fmt"
	"logparseProject/common"
	"os"
	"time"
)

func main() {
	start := time.Now()
	common.Flag()
	t := time.Since(start)
	fmt.Printf("[*] 收集结束,耗时: %s", t)
	fmt.Println()
	fmt.Println("按任意键退出")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
