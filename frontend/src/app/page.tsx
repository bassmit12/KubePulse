type Service = {
  id: string;
  name: string;
  namespace: string;
  status: "HEALTHY" | "WARNING" | "CRITICAL";
  p95LatencyMs: number;
  errorRate: number;
};

type Incident = {
  id: string;
  title: string;
  status: "OPEN" | "INVESTIGATING" | "MITIGATED" | "RESOLVED";
  severity: "LOW" | "MEDIUM" | "HIGH" | "CRITICAL";
  startedAt: string;
};

const apiBase =
  process.env.NEXT_PUBLIC_API_BASE?.replace(/\/$/, "") ?? "http://localhost:4000";

async function getServices(): Promise<Service[]> {
  const res = await fetch(`${apiBase}/services`, { cache: "no-store" });
  if (!res.ok) return [];
  const data = (await res.json()) as { items: Service[] };
  return data.items ?? [];
}

async function getIncidents(): Promise<Incident[]> {
  const res = await fetch(`${apiBase}/incidents`, { cache: "no-store" });
  if (!res.ok) return [];
  const data = (await res.json()) as { items: Incident[] };
  return data.items ?? [];
}

async function getHealth(): Promise<string> {
  const res = await fetch(`${apiBase}/healthz`, { cache: "no-store" });
  if (!res.ok) return "down";
  const data = (await res.json()) as { status: string };
  return data.status ?? "unknown";
}

function statusClass(status: Service["status"]) {
  if (status === "HEALTHY") return "pill pill--healthy";
  if (status === "WARNING") return "pill pill--warning";
  return "pill pill--critical";
}

export default async function Home() {
  const [health, services, incidents] = await Promise.all([
    getHealth(),
    getServices(),
    getIncidents(),
  ]);

  return (
    <div className="notion-shell">
      <aside className="notion-sidebar">
        <h2>KubePulse AI</h2>
        <nav>
          <a className="active" href="#">
            Dashboard
          </a>
          <a href="#">Services</a>
          <a href="#">Incidents</a>
          <a href="#">Auth</a>
        </nav>
      </aside>

      <main className="notion-main">
        <header className="page-header">
          <div>
            <h1>Operations Dashboard</h1>
            <p>Sprint 1 foundation · Encore backend + Next.js UI</p>
          </div>
          <div className={`pill ${health === "ok" ? "pill--healthy" : "pill--critical"}`}>
            Backend: {health}
          </div>
        </header>

        <section className="card-grid">
          <article className="card">
            <h3>Services</h3>
            <p className="metric">{services.length}</p>
            <span>tracked workloads</span>
          </article>
          <article className="card">
            <h3>Open incidents</h3>
            <p className="metric">{incidents.filter((i) => i.status !== "RESOLVED").length}</p>
            <span>active investigations</span>
          </article>
          <article className="card">
            <h3>Critical services</h3>
            <p className="metric">
              {services.filter((s) => s.status === "CRITICAL").length}
            </p>
            <span>require immediate action</span>
          </article>
        </section>

        <section className="card table-card">
          <h3>Service Health</h3>
          <table>
            <thead>
              <tr>
                <th>Name</th>
                <th>Namespace</th>
                <th>Status</th>
                <th>P95 Latency</th>
                <th>Error Rate</th>
              </tr>
            </thead>
            <tbody>
              {services.map((s) => (
                <tr key={s.id}>
                  <td>{s.name}</td>
                  <td>{s.namespace}</td>
                  <td>
                    <span className={statusClass(s.status)}>{s.status}</span>
                  </td>
                  <td>{s.p95LatencyMs} ms</td>
                  <td>{(s.errorRate * 100).toFixed(2)}%</td>
                </tr>
              ))}
            </tbody>
          </table>
        </section>

        <section className="card table-card">
          <h3>Incident Feed</h3>
          <table>
            <thead>
              <tr>
                <th>Title</th>
                <th>Status</th>
                <th>Severity</th>
                <th>Started</th>
              </tr>
            </thead>
            <tbody>
              {incidents.map((i) => (
                <tr key={i.id}>
                  <td>{i.title}</td>
                  <td>{i.status}</td>
                  <td>{i.severity}</td>
                  <td>{new Date(i.startedAt).toLocaleString()}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </section>
      </main>
    </div>
  );
}
