package service

import (
	"shop-backend/model"
	"shop-backend/repository"
)

type ProductsService struct {
	ProductsRepository *repository.ProductsRepository
}

func (s ProductsService) GetProductsByUser(id string) [] model.Product {
	return s.ProductsRepository.GetProductsByUser(id)
}

func (s ProductsService) GetAllProducts() [] model.Product {
	return s.ProductsRepository.GetAllProducts()
}

func (s ProductsService) Create(product model.Product) error {
	return s.ProductsRepository.Create(product)
}

func (s ProductsService) GetProductById(id string) (model.Product, error) {
	return s.ProductsRepository.GetProductById(id)
}

func (s ProductsService) Delete(id string) error {
	return s.ProductsRepository.Delete(id)
}

func (s ProductsService) Update(product model.Product) error {
	return s.ProductsRepository.Update(product)
}
