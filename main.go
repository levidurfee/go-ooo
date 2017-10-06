package main

import (
	"./ween"
	"./weenie"
	"fmt"
)

func main() {
	weenie.Display("Hi")

	w := ween.Ween{}
	w.SetCats(4)
	c := w.Cats()
	fmt.Println(c)

	ween.Poop(w)
}
