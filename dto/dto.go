package dto

import (
	"books-list/err"
)

type BookResponse struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type NewBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type NewBookResponse struct {
	Id int `json:"id"`
}

type UpdateBookRequest struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

func (r NewBookRequest) Validate() *err.Error {
	var err err.Error
	if r.Author == "" || r.Title == "" || r.Year == "" {
		err.Message = "Enter missing fields."
		return &err
	}
	return nil
}

func (r UpdateBookRequest) Validate() *err.Error {
	var err err.Error
	if &r.Id == nil || r.Id <= 0 || r.Author == "" || r.Title == "" || r.Year == "" {
		err.Message = "All fields should be filled in."
		return &err
	}
	return nil
}