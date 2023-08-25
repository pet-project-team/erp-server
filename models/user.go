package models

type User struct {
	BaseModel
	FirstName string `json:"first_name" gorm:"column:first_name;type:varchar(50);not null"`
	LastName  string `json:"last_name" gorm:"column:last_name;type:varchar(50);not null"`
	Email     string `json:"email" gorm:"column:email;type:varchar(100);not null"`
	Password  string `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Social    string `json:"social"`
	SocialID  string `json:"social_id"`
	RoleKey   string `json:"role_key" gorm:"column:role_key;type:varchar(50);not null"`
}
