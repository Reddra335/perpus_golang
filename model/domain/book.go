package domain

type Book struct {
	Id         int
	Title      string
	Author     string
	Isbn       string
	Stock      int
	CategoryId int
	FilePath   string
}
