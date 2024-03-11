package model

type Reader struct {
	Id          int    `json:"id"`
	FullName    string `json:"full_name"`
	ListOfBooks []int  `json:"list_of_books"`
}
