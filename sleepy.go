package main

import (
	"fmt"
	"time"
)

func main() {

	const n time.Duration = 2

	for {
		time.Sleep(n * time.Second)

		t := time.Now()
		fmt.Println("Ok, slept some time. Now:", t)
	}
}
