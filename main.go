package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("一直运行12", time.Now().Unix())
	}
}
