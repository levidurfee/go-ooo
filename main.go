package main

import (
	"fmt"
	"github.com/levidurfee/go/pages"
)

func main() {
	fmt.Println("Starting")
	home := pages.Page { Id: 1, Url: "/", Title: "Home", Content: "Welcome" }
	about := pages.NewPage(2, "/about/", "About", "Hii")
	fmt.Println(home.Id)
	fmt.Println(about.Id)
}
