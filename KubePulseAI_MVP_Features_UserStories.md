# KubePulse AI — Feature List & User Stories

## 0) Product Direction

**Working product name:** KubePulse AI  
**Goal:** Unified Kubernetes observability + network/security anomaly detection, with explainable alerts and role-based workflows.

**Core personas:**
- **Platform Engineer / SRE** (cluster health, SLO, incident response)
- **Security Engineer (SecOps)** (network threat detection, least privilege)
- **Developer** (service-level performance and dependencies)
- **Team Lead / Manager** (high-level status, trend reporting)

---

## 1) MVP Feature List (Prioritized)

## 1.1 Authentication & Access (Keycloak)

### Features
- OIDC login via Keycloak
- Role-based access control (Admin, SRE, SecOps, Developer, Viewer)
- Project/namespace-level access scopes
- Session management and logout

### User stories
1. **As an SRE**, I want to log in with my organizational account so I can securely access cluster dashboards.  
   **Acceptance criteria:**
   - Login redirects to Keycloak
   - On success, user profile and roles are available in app
   - Unauthorized users are denied access

2. **As an Admin**, I want role-based access so sensitive security data is only visible to authorized roles.  
   **Acceptance criteria:**
   - Roles map to permissions (view-only, investigate, manage)
   - Role changes in Keycloak reflect in app after token refresh

3. **As a Team Lead**, I want namespace-level restrictions so teams only see their own workloads.  
   **Acceptance criteria:**
   - User can access only authorized namespaces
   - GraphQL/API queries automatically filter unauthorized data

---

## 1.2 Cluster Overview Dashboard (Next.js + Prometheus)

### Features
- Live cluster health summary (CPU, memory, pod status, restarts)
- Namespace and service health cards
- Error-rate and latency trend charts
- Drill-down from cluster → namespace → workload

### User stories
4. **As an SRE**, I want a single overview page with cluster KPIs so I can quickly detect degradation.  
   **Acceptance criteria:**
   - Dashboard loads key KPIs in <3 seconds for baseline dataset
   - Highlights unhealthy workloads and restart spikes

5. **As a Developer**, I want to drill into my service metrics so I can diagnose regressions quickly.  
   **Acceptance criteria:**
   - Clicking service opens detailed charts (RPS, p95/p99, error rate)
   - Time range filters apply consistently across charts

6. **As a Viewer**, I want readable visual summaries so I can understand status without deep technical context.  
   **Acceptance criteria:**
   - Clear status indicators (healthy/warning/critical)
   - Definitions/tooltips for each metric

---

## 1.3 Network Observability (Cilium + Hubble)

### Features
- Service-to-service network flow map
- Flow filters (namespace, source, destination, protocol, verdict)
- Denied/blocked traffic visibility
- Top talkers and unusual path changes

### User stories
7. **As a SecOps engineer**, I want to see denied traffic flows so I can detect policy and attack-related issues.  
   **Acceptance criteria:**
   - Denied flows are searchable and timestamped
   - Each flow shows source/destination context

8. **As an SRE**, I want a service dependency map so I can understand blast radius during incidents.  
   **Acceptance criteria:**
   - Nodes/edges represent active communication
   - Clicking an edge reveals recent flow and latency details

9. **As a Developer**, I want to filter network flows by my namespace so I can isolate my service traffic.  
   **Acceptance criteria:**
   - Namespace filter updates map and flow table
   - Export filtered results (CSV/JSON)

---

## 1.4 AI Anomaly Detection Engine

### Features
- Baseline modeling for key signals (latency, error rate, restart rate, flow volume)
- Real-time anomaly scoring
- Explainability panel (which signals deviated + confidence)
- Alert suppression windows and dedup

### User stories
10. **As an SRE**, I want automatic anomaly alerts so I can detect incidents before users report them.  
    **Acceptance criteria:**
    - Alert generated when anomaly score crosses threshold
    - Alert includes affected service/namespace and timestamp

11. **As a SecOps engineer**, I want network behavior anomalies flagged so I can detect suspicious lateral movement.  
    **Acceptance criteria:**
    - New/unusual communication paths trigger anomalies
    - Alert includes baseline vs observed delta

12. **As an incident responder**, I want explainable alerts so I can trust and act on the signal.  
    **Acceptance criteria:**
    - Alert details show top contributing factors
    - Confidence score and evidence links are visible

---

## 1.5 Incident Feed & Investigation Workspace

### Features
- Unified incident timeline (metrics + flows + alerts)
- Incident status workflow (Open, Investigating, Mitigated, Resolved)
- Assignment and notes
- Suggested next checks/playbook links

### User stories
13. **As an SRE**, I want a chronological incident timeline so I can reconstruct what happened quickly.  
    **Acceptance criteria:**
    - Timeline correlates alerts, metric spikes, and flow anomalies
    - Time cursor syncs all charts/tables

14. **As a SecOps engineer**, I want to annotate incidents so handover is smooth between shifts.  
    **Acceptance criteria:**
    - Notes and status changes are saved with actor + timestamp

15. **As a Team Lead**, I want incident ownership and status visibility so accountability is clear.  
    **Acceptance criteria:**
    - Each incident has owner + current status
    - Filter incidents by status/owner/date

---

## 1.6 GraphQL API Layer

### Features
- GraphQL gateway over Encore.go services
- Typed query access to metrics, flows, anomalies, incidents
- Pagination + filtering primitives
- Auth-aware resolvers (role and namespace scoped)

### User stories
16. **As a frontend engineer**, I want a single GraphQL endpoint so UI development is faster and consistent.  
    **Acceptance criteria:**
    - GraphQL schema documented and introspectable
    - Main dashboards can be served from GraphQL only

17. **As an API consumer**, I want flexible queries so I can fetch only needed fields and reduce payload size.  
    **Acceptance criteria:**
    - Query complexity controls in place
    - Pagination and filtering supported for large result sets

18. **As an Admin**, I want API access enforcement so unauthorized queries are blocked by role/scope.  
    **Acceptance criteria:**
    - Unauthorized fields/records return permission errors

---

## 1.7 Performance Monitoring & SLO Module

### Features
- SLI definitions (latency, availability, error rate)
- SLO target tracking and burn-rate indicators
- Performance anomaly overlays
- Weekly performance trend summary

### User stories
19. **As an SRE**, I want SLO burn-rate indicators so I can prioritize high-risk services.  
    **Acceptance criteria:**
    - Burn-rate shown per service with threshold bands
    - Services sorted by risk level

20. **As a Developer**, I want p95/p99 overlays with deployment markers so I can link regressions to changes.  
    **Acceptance criteria:**
    - Deploy events shown on latency charts
    - Regression windows can be selected and compared

21. **As a Team Lead**, I want weekly trend summaries so I can report service health improvements or decline.  
    **Acceptance criteria:**
    - Exportable summary with key KPIs and top incidents

---

## 2) Non-Functional Requirements (MVP)

- **Security:** JWT/OIDC validation; tenant/namespace authorization on every request
- **Performance:** P95 API response under defined threshold for dashboard queries
- **Scalability:** Handle increasing flow/metric volume via async ingestion pipelines
- **Reliability:** Retry + backoff for collectors; partial failure isolation
- **Auditability:** Incident actions and admin operations logged
- **Observability of the platform itself:** self-monitoring metrics for backend components

---

## 3) Suggested Epic Breakdown

1. **Epic A — Identity & Access**
   - Keycloak integration
   - Role model + access middleware

2. **Epic B — Data Ingestion Foundation**
   - Prometheus query adapters
   - Hubble flow ingestion/storage pipeline

3. **Epic C — Core UI Dashboards**
   - Cluster overview
   - Service detail pages
   - Network map

4. **Epic D — Anomaly Engine v1**
   - Baseline logic + scoring
   - Alert generation + explainability

5. **Epic E — Incident Workflow**
   - Timeline
   - Assignment/status/notes

6. **Epic F — SLO/Performance Module**
   - SLI/SLO definitions
   - Burn-rate and trend reports

---

## 4) MVP Exclusions (to keep scope realistic)

- Multi-cluster federation (single-cluster first)
- Full auto-remediation (recommendation-only in MVP)
- Advanced ML (start with statistical baseline + thresholding)
- Policy auto-apply (generate suggestions first)
- External SIEM integrations (phase 2)

---

## 5) Suggested Next Step (immediate)

Take 8–10 high-value stories (from #4, #7, #10, #13, #16, #19 plus access-control stories) and convert them into:
- Sprint backlog items
- Technical tasks per service (Next.js, Encore.go services, GraphQL, collectors)
- Definition of Done and demo scenarios
