package service

import (
	"errors"
	"log"
	"synapsis/internal/model"
	"synapsis/internal/repository"
)

type productService struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
}

type ProductService interface {
	GetAllProductByCategoryName(name string) ([]model.ProductResponse, error)
}

func NewProductService(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository) *productService {
	return &productService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
	}
}

func (s *productService) GetAllProductByCategoryName(name string) ([]model.ProductResponse, error) {
	var productResponses []model.ProductResponse

	products, err := s.productRepository.GetAllProductByCategoryName(name)
	if err != nil {
		log.Println("error: " + err.Error())
		return productResponses, err
	}

	if len(products) == 0 {
		log.Println("error: product not found")
		return productResponses, errors.New("product not found")
	}

	category, err := s.categoryRepository.GetNameCategoryByIdCategory(products[0].IdCategory)
	if err != nil {
		log.Println("error: " + err.Error())
		return productResponses, err
	}

	for _, product := range products {
		productResponse := model.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Category:    category,
			Price:       product.Price,
			Stok:        product.Stok,
			Image:       product.Image,
		}
		productResponses = append(productResponses, productResponse)
	}

	return productResponses, nil
}
