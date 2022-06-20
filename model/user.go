package model

import (
	"errors"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;unique" json:"password"`
	Nickname string `gorm:"size:255;not null;unique" json:"nickname"`
}

func (s *Server) ValidateEmail(email string) error {
	if email == "" {
		return errors.New("required email")
	}
	if email != "" {
		if err := checkmail.ValidateFormat(email); err != nil {
			return errors.New("invalid email")
		}
	}
	return nil
}

func (s *Server) CreateUser(user *User) (*User, error) {
	emailErr := s.ValidateEmail(user.Email)
	if emailErr != nil {
		return nil, emailErr
	}
	err := s.DB.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Server) GetUserByEmail(email string, password string) (*User, error) {
	user := &User{}
	err := s.DB.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}
	err = VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.New("Incorrect Password")
	}
	return user, nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
