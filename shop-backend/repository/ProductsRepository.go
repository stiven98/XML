package repository

import (
	"gorm.io/gorm"
	"shop-backend/model"
)

type ProductsRepository struct {
	Database *gorm.DB
}

func (r ProductsRepository) GetProductsByUser(id string) [] model.Product {
	var products [] model.Product
	r.Database.Model(&model.Product{}).Where("user_id = ?", id).Find(&products)
	return products
}

func (r ProductsRepository) GetAllProducts() []model.Product {
	var products [] model.Product
	r.Database.Model(&model.Product{}).Find(&products)
	return products
}

func (r ProductsRepository) Create(product model.Product) error {
	result := r.Database.Create(product)
	return result.Error
}

func (r ProductsRepository) GetProductById(id string) (model.Product, error) {
	var product model.Product
	var count int64
	res := r.Database.Model(&model.Product{}).Where("id = ?", id).First(&product).Count(&count)
	return product, res.Error
}

func (r ProductsRepository) Delete(id string) error {
	res := r.Database.Model(&model.Product{}).Where("id = ?", id).UpdateColumn("deleted", true)
	return res.Error
}

func (r ProductsRepository) Update(product model.Product) error {
	res := r.Database.Model(&model.Product{}).Where("id = ?", product.ID).Updates(product)
	return res.Error
}
