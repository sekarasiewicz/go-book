package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + " " + sep + os.Args[i]
		sep = " / "
	}
	fmt.Println(s)
	fmt.Printf("%.d Microseconds upłynęło \n", time.Since(start).Nanoseconds())
}
