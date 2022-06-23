package model

import (
	"errors"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                            uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Email                         string `json:"email" binding:"required"`
	Password                      string `json:"password" binding:"required"`
	Device_Id                     string `gorm:"size:255;not null;unique" json:"device_id"`
	Device_Name                   string `gorm:"size:255;not null;unique" json:"Device_Name"`
	Temp_Password                 string `gorm:"size:100;not null;" json:"Temp_Password"`
	Name                          string `gorm:"size:100;not null;" json:"Name"`
	Photo                         string `gorm:"size:100;not null;" json:"Photo"`
	Last_Juz                      string `gorm:"size:100;not null;" json:"Last_Juz"`
	Last_Chapter                  string `gorm:"size:100;not null;" json:"Last_Chapter"`
	Last_Verse                    string `gorm:"size:100;not null;" json:"Last_Verse"`
	Last_Read_Date                string `gorm:"size:100;not null;" json:"Last_Read_Date"`
	Scheduled_Khatam_Date         string `gorm:"size:100;not null;" json:"Scheduled_Khatam_Date"`
	Default_Voice_Path            string `gorm:"size:100;not null;" json:"Default_Voice_Path"`
	Penyakit_Khusus               string `gorm:"size:100;not null;" json:"Penyakit_Khusus"`
	Obat_Khusus                   string `gorm:"size:100;not null;" json:"Obat_Khusus"`
	Gol_Darah                     string `gorm:"size:100;not null;" json:"Gol_Darah"`
	Ktp_Image                     string `gorm:"size:100;not null;" json:"Ktp_Image"`
	Passport_Image                string `gorm:"size:100;not null;" json:"Passport_Image"`
	Vaccination_Certificate_Image string `gorm:"size:100;not null;" json:"Vaccination_Certificate_Image"`
	Verification_Token            string `gorm:"size:100;not null;" json:"Verification_Token"`
	Verification_Token_Date       string `gorm:"size:100;not null;" json:"Verification_Token_Date"`
	Email_Verified                string `gorm:"size:100;not null;" json:"Email_Verified"`
	Lat                           string `gorm:"size:100;not null;" json:"Lat"`
	Lng                           string `gorm:"size:100;not null;" json:"Lng"`
	Current_Location              string `gorm:"size:100;not null;" json:"Current_Location"`
	Parking_Lat                   string `gorm:"size:100;not null;" json:"Parking_Lat"`
	Parking_Lng                   string `gorm:"size:100;not null;" json:"Parking_Lng"`
	Driver_Status                 string `gorm:"size:100;not null;" json:"Driver_Status"`
	Google_Uid                    string `gorm:"size:100;not null;" json:"Google_Uid"`
	Facebook_Uid                  string `gorm:"size:100;not null;" json:"Facebook_Uid"`
	Fcm_Key                       string `gorm:"size:100;not null;" json:"Fcm_Key"`
	Pushy_Token                   string `gorm:"size:100;not null;" json:"Pushy_Token"`
	Xmpp_Password                 string `gorm:"size:100;not null;" json:"Xmpp_Password"`
	Lang_Code                     string `gorm:"size:100;not null;" json:"Lang_Code"`
	Premium                       string `gorm:"size:100;not null;" json:"Premium"`
	Premium_Months                string `gorm:"size:100;not null;" json:"Premium_Months"`
	Notification_Channel_Id       string `gorm:"size:100;not null;" json:"Notification_Channel_Id"`
	Group_Notification_Channel_Id string `gorm:"size:100;not null;" json:"Group_Notification_Channel_Id"`
	Last_Premium_Date             string `gorm:"size:100;not null;" json:"Last_Premium_Date"`
}

type UserLogin struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Device_Id   string `gorm:"size:255;not null;unique" json:"device_id"`
	Device_Name string `gorm:"size:255;not null;unique" json:"Device_Name"`
	Lang_Code   string `gorm:"size:100;not null;" json:"Lang_Code"`
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
	hashedPassword, errd := Hash(user.Password)
	if errd != nil {
		return nil, errd
	}
	user.Password = string(hashedPassword)
	emailErr := s.ValidateEmail(user.Email)
	if emailErr != nil {
		return nil, emailErr
	}

	err := s.DB.Debug().Create(&user).Error
	if err != nil {
		return nil, errors.New("Already Registered")
	}
	return user, nil
}

func (s *Server) GetUserByEmail(email string, password string) (*User, error) {
	user := &User{}
	err := s.DB.Debug().Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, errors.New("User Not Found")
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
