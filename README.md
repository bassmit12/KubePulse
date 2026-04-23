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
- `backend/` plain Go + GraphQL starter scaffold
- `backend-encore/` Encore Go app (active backend scaffold)
- `infra/` Kubernetes manifests and deployment scaffolding
- `docs/` architecture docs and planning

## Quick Start (local)
1. Frontend
   - `cd frontend`
   - `npm run dev`
2. Encore backend
   - `cd backend-encore`
   - `encore run`
3. Plain Go backend (optional fallback)
   - `cd backend`
   - `go mod tidy`
   - `go run ./cmd/api`

## Current status
Initial app scaffold generated and Encore backend bootstrapped with endpoints:
- `GET /healthz`
- `GET /services?namespace=kubepulse`
- `GET /incidents?status=OPEN`

Next steps:
- Keycloak OIDC integration in backend-encore
- GraphQL gateway in backend-encore
- Prometheus/Hubble ingestion adapters
- anomaly engine v1
