package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""

	for i, arg := range os.Args[1:] {
		s += strconv.Itoa(i) + sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%.2fs upłynęło \n", time.Since(start).Seconds())
}
