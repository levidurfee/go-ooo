package pages

type Page struct {
	Id int
	Url string
	Title string
	Content string
}

func NewPage(id int, url string, title string, content string) Page {
	return Page {Id: id, Url: url, Title: title, Content: content}
}
