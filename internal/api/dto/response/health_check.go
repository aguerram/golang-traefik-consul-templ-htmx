package response

type HealthCheckResponse struct {
	Status     string `json:"status"`
	Components []struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	} `json:"components"`
}

func NewHealthCheckResponse() *HealthCheckResponse {
	return &HealthCheckResponse{
		Status: "UP",
		Components: make([]struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		}, 0, 1),
	}
}

func (r *HealthCheckResponse) AddComponentStatus(name string, isUp bool) {
	status := "UP"
	if !isUp {
		status = "DOWN"
		if r.Status == "UP" {
			r.Status = "DOWN"
		}
	}
	r.Components = append(r.Components, struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}{Name: name, Status: status})
}
