package repositories

import (
	"waysfood/models"

	"gorm.io/gorm"
)

// type repository struct {
// 	db *gorm.DB
// }

type ProductRepository interface {
	FindProduct() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(user models.Product) (models.Product, error)
	UpdateProduct(user models.Product) (models.Product, error)
	DeleteProduct(user models.Product) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProduct() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Find(&products).Error
	// err := r.db.Find(&product).Error
	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	// err := r.db.First(&user, ID).Error
	err := r.db.Preload("User").First(&product, ID).Error

	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error

	return product, err
}

func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}
