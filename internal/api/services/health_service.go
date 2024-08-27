package services

type HealthService interface {
	HealthCheck() (bool, error)
}

type healthService struct {
}

func NewHealthService() HealthService {
	return &healthService{}
}

func (h *healthService) HealthCheck() (bool, error) {
	return true, nil
}
