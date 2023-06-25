package models

type User struct {
	ID       *int  `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "dineder_user2"
}