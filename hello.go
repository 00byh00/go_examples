package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println(quote.Go())
	fmt.Println("running main func...")

	//compute(5)
	//compute(5)
	go compute(5)
	go compute(5)

	fmt.Scanln()
	fmt.Println("Done...")
}
