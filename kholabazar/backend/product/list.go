package product

import "kholabazar/domain"

func (svc *service) List() ([]*domain.Product, error) {
	return svc.productRepo.List()
}
