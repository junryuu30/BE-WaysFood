package productdto

type ProductRequest struct {
	Title  string `json:"title" form:"name" gorm:"type: varchar(255)" validate:"required"`
	Image  string `json:"image" from:"image" gorm:"type: varchar(255)"`
	Price  int    `json:"price" form:"name" gorm:"type: varchar(255)" validate:"required"`
	Qty    int    `json:"qty" form:"qty" gorm:"type: int" validate:"required"`
	UserID int    `json:"user_id"`
}

// type UpdateProductRequst struct {
// 	Name  string `json:"name" form:"name" gorm:"type: varchar(255)"`
// 	Desc  string `json:"desc" form:"desc" gorm:"type:varchar(255)"`
// 	Price int    `json:"price" from:"price" gorm:"type: varchar(255)"`
// 	Qty   int    `json:"qty" from:"qty" gorm:"type: int"`
// }
