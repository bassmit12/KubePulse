package graphql

import (
	"context"
	"strings"

	"encore.app/health"
	"encore.app/incidents"
	"encore.app/services"
)

type Request struct {
	Query string `json:"query"`
}

type Data struct {
	Health    *health.HealthResponse      `json:"health,omitempty"`
	Services  []*services.Service         `json:"services,omitempty"`
	Incidents []*incidents.Incident       `json:"incidents,omitempty"`
}

type Response struct {
	Data   *Data    `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// Query provides a minimal GraphQL-like endpoint for MVP integration.
//encore:api public method=POST path=/graphql
func Query(ctx context.Context, req *Request) (*Response, error) {
	if req == nil || strings.TrimSpace(req.Query) == "" {
		return &Response{Errors: []string{"query is required"}}, nil
	}

	q := strings.ToLower(req.Query)
	resp := &Response{Data: &Data{}}

	if strings.Contains(q, "health") {
		h, err := health.Ping(ctx)
		if err != nil {
			return nil, err
		}
		resp.Data.Health = h
	}

	if strings.Contains(q, "services") {
		s, err := services.List(ctx, &services.ListParams{})
		if err != nil {
			return nil, err
		}
		resp.Data.Services = s.Items
	}

	if strings.Contains(q, "incidents") {
		i, err := incidents.List(ctx, &incidents.ListParams{})
		if err != nil {
			return nil, err
		}
		resp.Data.Incidents = i.Items
	}

	if resp.Data.Health == nil && len(resp.Data.Services) == 0 && len(resp.Data.Incidents) == 0 {
		resp.Errors = []string{"unsupported query fields; try health, services, incidents"}
	}

	return resp, nil
}
