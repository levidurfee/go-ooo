package ween

import "fmt"

type Ween struct {
	cats int
}

type Weeen interface {
	Cats() int
}

func Poop(w Weeen) {
	fmt.Println(w.Cats())
}

func (w *Ween) SetCats(c int) {
	w.cats = c
}

func (w Ween) Cats() int {
	return w.cats
}
