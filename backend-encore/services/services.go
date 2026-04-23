package services

import "context"

type ServiceStatus string

const (
	Healthy  ServiceStatus = "HEALTHY"
	Warning  ServiceStatus = "WARNING"
	Critical ServiceStatus = "CRITICAL"
)

type Service struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Namespace    string        `json:"namespace"`
	Status       ServiceStatus `json:"status"`
	P95LatencyMS float64       `json:"p95LatencyMs"`
	ErrorRate    float64       `json:"errorRate"`
}

type ListParams struct {
	Namespace string `query:"namespace"`
}

type ListResponse struct {
	Items []*Service `json:"items"`
}

var sample = []*Service{
	{ID: "svc-1", Name: "frontend", Namespace: "kubepulse", Status: Healthy, P95LatencyMS: 110, ErrorRate: 0.004},
	{ID: "svc-2", Name: "api", Namespace: "kubepulse", Status: Warning, P95LatencyMS: 360, ErrorRate: 0.019},
	{ID: "svc-3", Name: "anomaly-engine", Namespace: "kubepulse", Status: Healthy, P95LatencyMS: 95, ErrorRate: 0.001},
}

// List returns service status for dashboard bootstrapping.
//encore:api public method=GET path=/services
func List(ctx context.Context, p *ListParams) (*ListResponse, error) {
	if p == nil || p.Namespace == "" {
		return &ListResponse{Items: sample}, nil
	}

	filtered := make([]*Service, 0)
	for _, s := range sample {
		if s.Namespace == p.Namespace {
			filtered = append(filtered, s)
		}
	}
	return &ListResponse{Items: filtered}, nil
}
