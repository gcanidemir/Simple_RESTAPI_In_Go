package entity

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []Book{
	{ID: 1, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	{ID: 2, Title: "To Kill a Mockingbird", Author: "Harper Lee"},
	{ID: 3, Title: "1984", Author: "George Orwell"},
}
