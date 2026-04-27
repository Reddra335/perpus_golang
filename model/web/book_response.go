package web

type BookResponse struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Isbn       string `json:"isbn"`
	Stock      int    `json:"stock"`
	CategoryId int    `json:"category_id"`
	FilePath   string `json:"file_path"`
}
