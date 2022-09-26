package model

type Book struct {
	Id        string `uri:"id" json:"id"`
	CreatedAt int64 
	UpdatedAt int64
	Title     string `json:"title"`
	Author    string `json:"author"`
}
