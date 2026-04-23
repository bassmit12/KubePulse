package incidents

import "context"

type IncidentStatus string

type Severity string

const (
	Open          IncidentStatus = "OPEN"
	Investigating IncidentStatus = "INVESTIGATING"
	Mitigated     IncidentStatus = "MITIGATED"
	Resolved      IncidentStatus = "RESOLVED"
)

const (
	Low      Severity = "LOW"
	Medium   Severity = "MEDIUM"
	High     Severity = "HIGH"
	Critical Severity = "CRITICAL"
)

type Incident struct {
	ID        string         `json:"id"`
	Title     string         `json:"title"`
	Status    IncidentStatus `json:"status"`
	Severity  Severity       `json:"severity"`
	StartedAt string         `json:"startedAt"`
}

type ListParams struct {
	Status string `query:"status"`
}

type ListResponse struct {
	Items []*Incident `json:"items"`
}

var sample = []*Incident{
	{ID: "inc-1001", Title: "Spike in API p95 latency", Status: Investigating, Severity: High, StartedAt: "2026-04-23T18:20:00Z"},
	{ID: "inc-1002", Title: "Denied flow burst from checkout to db", Status: Open, Severity: Medium, StartedAt: "2026-04-23T18:43:00Z"},
}

// List returns the active incident feed.
//encore:api public method=GET path=/incidents
func List(ctx context.Context, p *ListParams) (*ListResponse, error) {
	if p == nil || p.Status == "" {
		return &ListResponse{Items: sample}, nil
	}
	filtered := make([]*Incident, 0)
	for _, in := range sample {
		if string(in.Status) == p.Status {
			filtered = append(filtered, in)
		}
	}
	return &ListResponse{Items: filtered}, nil
}
