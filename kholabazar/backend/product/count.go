package product

func (svc *service) Count() (int64, error){
	return svc.productRepo.Count()
}