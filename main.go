package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("一直运行", time.Now().Unix())
	}
}
