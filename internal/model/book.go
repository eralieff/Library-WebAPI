package model

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	ISBN     string `json:"isbn"`
	AuthorId int    `json:"author_id"`
}
