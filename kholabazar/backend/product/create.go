package product

import "kholabazar/domain"

func (svc *service) Create(product domain.Product) (*domain.Product, error) {
   return svc.productRepo.Create(product)
}
