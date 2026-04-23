package health

import "context"

type HealthResponse struct {
	Status string `json:"status"`
}

// Ping returns backend liveness status.
//encore:api public method=GET path=/healthz
func Ping(ctx context.Context) (*HealthResponse, error) {
	return &HealthResponse{Status: "ok"}, nil
}
