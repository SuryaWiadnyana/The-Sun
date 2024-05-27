package Model

type Books struct {
	ID      string    `json:"id"`
	Judul   string `json:"judul"`
	Penulis string `json:"penulis"`
	Halaman int    `json:"halaman"`
	Status  string `json:"status"`
}
