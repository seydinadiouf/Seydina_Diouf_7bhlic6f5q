package model

type User struct {
	UserId   int    `gorm:"column:user_id" json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (User) TableName() string {
	return "application_users"
}
