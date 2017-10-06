package ween

type Ween struct {
	cats int
}

func (w *Ween) SetCats(c int) {
	w.cats = c
}

func (w Ween) Cats() int {
	return w.cats
}
