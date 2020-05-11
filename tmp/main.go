package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	end := time.Now()
	d := end.Sub(start) / time.Millisecond
	fmt.Println(d)
}
