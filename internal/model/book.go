package model

type Book struct {
	Title       string  `json:"title"`
	Author      Authors `json:"author"`
	Year_public int     `json:"year_public"`
	Year_read   int     `json:"year_read"`
	Rate        int     `json:"rate"`
}

type Authors struct {
	Author_name string `json:"author_name"`
	Year_born   int    `json:"year_b"`
	Country     string `json:"country"`
}

type Book2 struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Year_public int    `json:"year_public"`
	Year_read   int    `json:"year_read"`
	Rate        int    `json:"rate"`
	C_id        int    `json:"c_id"`
}
