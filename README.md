# KubePulse AI

Integrated Kubernetes observability + anomaly detection platform.

## Stack
- Frontend: Next.js (TypeScript)
- Backend: Go (Encore-ready service boundaries)
- API: GraphQL
- Identity: Keycloak (OIDC)
- Observability: Prometheus + Cilium + Hubble

## Project Structure
- `frontend/` Next.js dashboard UI
- `backend/` Go services + GraphQL API
- `infra/` Kubernetes manifests and deployment scaffolding
- `docs/` architecture docs and planning

## Quick Start (local)
1. Frontend
   - `cd frontend`
   - `npm run dev`
2. Backend (placeholder scaffolding in progress)
   - `cd backend`
   - `go mod tidy`
   - `go run ./cmd/api`

## Current status
Initial scaffold generated. Next steps:
- Keycloak OIDC integration
- GraphQL schema + resolvers
- Prometheus/Hubble ingestion adapters
- anomaly engine v1
