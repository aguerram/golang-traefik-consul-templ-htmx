package services

type Service struct {
	HealthService HealthService
}

func NewService() *Service {
	return &Service{
		HealthService: NewHealthService(),
	}
}
