package models

type User struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname" gorm:"type: varchar(255)"`
	Email     string `json:"email" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Location  string `json:"location" gorm:"type: varchar(255)"`
	Image     string `json:"image" gorm:"type: varchar(255)"`
	Role      string `json:"role" gorm:"type: varchar(255)"`
	Product   string `json:"product" gorm:"type: varchar(255)"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

// type UserProfileResponse struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// func (UserProfileResponse) TableName() string {
// 	return "users"
// }
