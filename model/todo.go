package model

import "errors"

type Todo struct {
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID uint64 `gorm:"not null" json:"user_id"`
	Title  string `gorm:"size:255;not null" json:"title" binding:"required"`
}

func (s *Server) CreateTodo(todo *Todo) (*Todo, error) {
	if todo.Title == "" {
		return nil, errors.New("Please Provide a Valid Title")
	}
	if todo.UserID == 0 {
		return nil, errors.New("a Valid User Id is Required")
	}
	err := s.DB.Debug().Create(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}
