package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("一直运行1", time.Now().Unix())
	}
}
