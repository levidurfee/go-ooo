package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	go say("hello")
	go say("HOUSE")
	go say("light")
	go say("FRANK")

	time.Sleep(5 * 500 * time.Millisecond)
}
