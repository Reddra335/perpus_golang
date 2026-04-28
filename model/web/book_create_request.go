package web

type BookCreateRequest struct {
	Title      string `validate:"required,max:200,min:1" json:"title"`
	Author     string `validate:"required,max:100,min:1" json:"author"`
	Isbn       string `validate:"required,max:20,min:1" json:"isbn"`
	Stock      int    `validate:"required,gte=0" json:"stock"`
	CategoryId int    `validate:"required" json:"categor_yid"`
	FilePath   string `json:"file_path"`
}
