package service

import (
	"shop/model"
	"time"
)

type (
	productService struct {
		products []model.Product
	}
)

// NewProductService :nodoc:
func NewProductService() ProductService {
	return &productService{
		products: []model.Product{
			model.Product{ID: time.Now().UnixNano(), Price: 100.0},
			model.Product{ID: time.Now().UnixNano(), Price: 200.0},
			model.Product{ID: time.Now().UnixNano(), Price: 300.0},
		},
	}
}

func (ps *productService) List() []model.Product {
	return ps.products
}

func (ps *productService) FindProductByID(id int64) *model.Product {
	for _, product := range ps.products {
		if product.ID == id {
			return &product
		}
	}

	return nil
}
