package dto

type BookResponse struct {
	Id  	int `json:"id"`
	Title  	string `json:"title"`
	Author 	string `json:"author"`
	Year   	string `json:"year"`
}

type NewBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type NewBookResponse struct {
	Id  int `json:"id"`
}

type UpdateBookRequest struct {
	Id  	int 	`json:"id"`
	Title  	string `json:"title"`
	Author 	string `json:"author"`
	Year   	string	`json:"year"`
}