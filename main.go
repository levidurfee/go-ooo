package main

import (
	"./ween"
	"./weenie"
	"./commands"
	"fmt"
)

func main() {
	weenie.Display("Hi")

	w := ween.Ween{}
	w.SetCats(4)
	c := w.Cats()
	fmt.Println(c)

	ween.Poop(w)
	commands.Shit()
	commands.Poop()

	i := 5
	j := &i
	*j = 6
	fmt.Println(i)
}
