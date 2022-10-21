package productdto

type ProductRequest struct {
	Name   string `json:"name" from:"name" gorm:"type: varchar(255)"`
	Desc   string `json:"desc" gorm:"type:text" form:"desc"`
	Price  int    `json:"price" from:"price" gorm:"type: int"`
	Image  string `json:"image" from:"image" gorm:"type: varchar(255)"`
	UserID int    `json:"user_id" gorm:"type: int"`
}
