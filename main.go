package main

import (
	"fmt"
	"logparseProject/common"
	"time"
)

func main() {
	start := time.Now()
	common.Flag()
	t := time.Since(start)
	fmt.Printf("[*] 收集结束,耗时: %s", t)
}
