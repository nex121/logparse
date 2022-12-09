package main

import (
	"fmt"
	"logparseProject/plugins"
	"time"
)

func main() {
	start := time.Now()
	plugins.GetRegeditList()
	//common.Flag()
	t := time.Since(start)
	fmt.Printf("[*] 收集结束,耗时: %s", t)
}
