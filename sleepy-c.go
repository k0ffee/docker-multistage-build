package main

import (
	"fmt"
	"time"
)

// #include <unistd.h>
import "C"

func main() {

	const n = C.uint(2)
	const timeSecond = C.uint(1000 * 1000)

	for {
		C.usleep(n * timeSecond)

		t := time.Now()
		fmt.Println("Ok, slept some time. Now:", t)
	}
}
