package service

type Service struct {
	ProductService ProductService
	AlertService   AlertService
}

func NewService(product ProductService, alert AlertService) (*Service, error) {
	return &Service{product, alert}, nil
}
