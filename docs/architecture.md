# Architecture Notes (initial)

## Planned runtime services
- auth-service (Keycloak integration)
- metrics-service (Prometheus queries)
- flow-service (Hubble/Cilium data)
- anomaly-service (baseline + anomaly scoring)
- incident-service (timeline/workflow)
- graphql-gateway (single client-facing endpoint)

## Data flow
1. Collect metrics and flows
2. Normalize into common event model
3. Run anomaly detection
4. Publish incidents/alerts via GraphQL
5. Render in Next.js dashboard
