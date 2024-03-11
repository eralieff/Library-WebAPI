package model

type Reader struct {
	Id          int    `json:"id"`
	FullName    string `json:"full_name"`
	ListOfBooks []int  `json:"list_of_books"`
}

type ReaderBook struct {
	ReaderName string `json:"reader_name"`
	BookTitle  string `json:"book_title"`
	Genre      string `json:"genre"`
	ISBN       string `json:"isbn"`
}
