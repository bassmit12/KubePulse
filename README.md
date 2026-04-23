# KubePulse AI

Integrated Kubernetes observability + anomaly detection platform.

## Stack
- Frontend: Next.js (TypeScript)
- Backend: Go (Encore-ready service boundaries)
- API: GraphQL
- Identity: Keycloak (OIDC)
- Observability: Prometheus + Cilium + Hubble

## Project Structure
- `frontend/` Next.js dashboard UI (Notion-style)
- `backend-encore/` Encore Go backend APIs
- `infra/` Kubernetes manifests and deployment scaffolding
- `docs/` architecture docs and planning

## Quick Start (local)
1. Backend (Encore)
   - `cd backend-encore`
   - `encore run`
2. Frontend
   - `cd frontend`
   - `npm run dev`

Set frontend API base if needed:
- PowerShell: `$env:NEXT_PUBLIC_API_BASE="http://localhost:4000"`

## Current status
Sprint 1 foundation includes:
- `GET /healthz`
- `GET /services`
- `GET /incidents`
- `POST /auth/verify` (Keycloak JWT verification via JWKS)
- `POST /graphql` (MVP query endpoint for health/services/incidents)

Next steps:
- Keycloak role-based route protection
- Prometheus/Hubble data adapters
- anomaly engine v1
