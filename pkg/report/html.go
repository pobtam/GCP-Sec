package report

import (
	"fmt"
	"html/template"
	"io"
	"sort"
	"time"

	"github.com/wanaware/gcp-security-analyzer/internal/models"
	"github.com/wanaware/gcp-security-analyzer/internal/utils"
)

const htmlTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>GCP Security Findings Report</title>
<style>
  :root {
    --critical: #dc2626; --high: #ea580c; --medium: #d97706; --low: #16a34a;
    --bg: #f8fafc; --card: #ffffff; --border: #e2e8f0; --text: #1e293b;
    --muted: #64748b;
  }
  * { box-sizing: border-box; margin: 0; padding: 0; }
  body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
         background: var(--bg); color: var(--text); line-height: 1.5; }
  .container { max-width: 1400px; margin: 0 auto; padding: 2rem 1rem; }
  h1 { font-size: 1.875rem; font-weight: 700; margin-bottom: 0.25rem; }
  h2 { font-size: 1.25rem; font-weight: 600; margin: 1.5rem 0 0.75rem; }
  .meta { color: var(--muted); font-size: 0.875rem; margin-bottom: 2rem; }
  .grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(180px, 1fr)); gap: 1rem; margin-bottom: 2rem; }
  .card { background: var(--card); border: 1px solid var(--border); border-radius: 0.5rem; padding: 1rem; }
  .card-label { font-size: 0.75rem; font-weight: 600; text-transform: uppercase; color: var(--muted); }
  .card-value { font-size: 2rem; font-weight: 700; }
  .card-sub { font-size: 0.875rem; color: var(--muted); }
  .critical .card-value { color: var(--critical); }
  .high .card-value { color: var(--high); }
  .medium .card-value { color: var(--medium); }
  .low .card-value { color: var(--low); }
  table { width: 100%; border-collapse: collapse; font-size: 0.875rem; background: var(--card);
          border: 1px solid var(--border); border-radius: 0.5rem; overflow: hidden; margin-bottom: 1.5rem; }
  th { background: #f1f5f9; padding: 0.625rem 0.75rem; text-align: left; font-weight: 600;
       font-size: 0.75rem; text-transform: uppercase; color: var(--muted);
       border-bottom: 1px solid var(--border); white-space: nowrap; }
  td { padding: 0.5rem 0.75rem; border-bottom: 1px solid var(--border); }
  tr:last-child td { border-bottom: none; }
  tr:hover { background: #f8fafc; }
  .badge { display: inline-block; padding: 0.125rem 0.5rem; border-radius: 9999px;
            font-size: 0.75rem; font-weight: 600; }
  .badge-CRITICAL { background: #fee2e2; color: var(--critical); }
  .badge-HIGH { background: #ffedd5; color: var(--high); }
  .badge-MEDIUM { background: #fef3c7; color: var(--medium); }
  .badge-LOW { background: #dcfce7; color: var(--low); }
  .score-bar { display: flex; align-items: center; gap: 0.5rem; }
  .bar { height: 6px; border-radius: 3px; background: #e2e8f0; flex: 1; }
  .bar-fill { height: 100%; border-radius: 3px; background: #3b82f6; }
  .section { margin-bottom: 2rem; }
  input[type=text] { padding: 0.5rem 0.75rem; border: 1px solid var(--border); border-radius: 0.375rem;
                     font-size: 0.875rem; width: 100%; max-width: 24rem; margin-bottom: 1rem; }
  .truncate { max-width: 200px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; }
  footer { text-align: center; color: var(--muted); font-size: 0.75rem; margin-top: 3rem; padding-top: 1rem;
           border-top: 1px solid var(--border); }
  details { border: 1px solid var(--border); border-radius: 0.375rem; margin-top: 0.5rem; }
  summary { cursor: pointer; padding: 0.5rem 0.75rem; font-weight: 600; font-size: 0.875rem;
            background: #f1f5f9; border-radius: 0.375rem; user-select: none; }
  details[open] summary { border-bottom: 1px solid var(--border); border-radius: 0.375rem 0.375rem 0 0; }
  .script-pre { background: #1e293b; color: #e2e8f0; padding: 1rem; border-radius: 0 0 0.375rem 0.375rem;
                overflow-x: auto; font-family: 'Menlo', 'Consolas', monospace; font-size: 0.78rem;
                line-height: 1.6; margin: 0; white-space: pre; }
  .script-card { background: var(--card); border: 1px solid var(--border); border-left: 4px solid var(--critical);
                 border-radius: 0.5rem; padding: 1rem; margin-bottom: 1rem; }
  .script-card h3 { font-size: 1rem; font-weight: 600; margin-bottom: 0.25rem; }
  .script-meta { font-size: 0.8rem; color: var(--muted); margin-bottom: 0.5rem; }
</style>
</head>
<body>
<div class="container">
  <h1>GCP Security Findings Report</h1>
  <p class="meta">Generated: {{ .GeneratedAt.UTC.Format "2006-01-02 15:04:05 UTC" }} &middot; Source: {{ .InputFile }}</p>

  <div class="grid">
    <div class="card">
      <div class="card-label">Total Findings</div>
      <div class="card-value">{{ .Stats.Total }}</div>
    </div>
    <div class="card critical">
      <div class="card-label">Critical</div>
      <div class="card-value">{{ .Stats.Critical }}</div>
      <div class="card-sub">{{ pct .Stats.Critical .Stats.Total }}%</div>
    </div>
    <div class="card high">
      <div class="card-label">High</div>
      <div class="card-value">{{ .Stats.High }}</div>
      <div class="card-sub">{{ pct .Stats.High .Stats.Total }}%</div>
    </div>
    <div class="card medium">
      <div class="card-label">Medium</div>
      <div class="card-value">{{ .Stats.Medium }}</div>
      <div class="card-sub">{{ pct .Stats.Medium .Stats.Total }}%</div>
    </div>
    <div class="card low">
      <div class="card-label">Low</div>
      <div class="card-value">{{ .Stats.Low }}</div>
      <div class="card-sub">{{ pct .Stats.Low .Stats.Total }}%</div>
    </div>
    <div class="card">
      <div class="card-label">Mean Risk Score</div>
      <div class="card-value">{{ .Stats.RiskStats.Mean }}</div>
      <div class="card-sub">Median: {{ .Stats.RiskStats.Median }}</div>
    </div>
  </div>

  <div class="section">
    <h2>All Findings</h2>
    <input type="text" id="searchInput" placeholder="Filter by category, resource, project..." onkeyup="filterTable()">
    <table id="findingsTable">
      <thead>
        <tr>
          <th>#</th>
          <th>Priority</th>
          <th>Risk Score</th>
          <th>Category</th>
          <th>Resource</th>
          <th>Project</th>
          <th>Severity</th>
          <th>CVE</th>
        </tr>
      </thead>
      <tbody>
        {{ range $i, $f := .Findings -}}
        <tr>
          <td>{{ inc $i }}</td>
          <td><span class="badge badge-{{ $f.Priority }}">{{ $f.Priority }}</span></td>
          <td>
            <div class="score-bar">
              <span>{{ riskScore $f }}</span>
              <div class="bar"><div class="bar-fill" style="width: {{ barWidth $f }}%"></div></div>
            </div>
          </td>
          <td>{{ $f.Category }}</td>
          <td class="truncate" title="{{ $f.ResourceName }}">{{ $f.ResourceDisplayName }}</td>
          <td>{{ $f.ProjectDisplayName }}</td>
          <td>{{ $f.Severity }}</td>
          <td>{{ if $f.HasCVE }}{{ $f.CVEID }}{{ else }}-{{ end }}</td>
        </tr>
        {{ end -}}
      </tbody>
    </table>
  </div>

  <div class="section">
    <h2>Top Categories</h2>
    <table>
      <thead>
        <tr><th>Category</th><th>Total</th><th>Critical</th><th>High</th><th>Medium</th><th>Low</th><th>Avg Score</th></tr>
      </thead>
      <tbody>
        {{ range sortedCategories .CategoryBreakdown -}}
        <tr>
          <td>{{ .Category }}</td>
          <td>{{ .Count }}</td>
          <td>{{ .Critical }}</td>
          <td>{{ .High }}</td>
          <td>{{ .Medium }}</td>
          <td>{{ .Low }}</td>
          <td>{{ printf "%.1f" .AvgRiskScore }}</td>
        </tr>
        {{ end -}}
      </tbody>
    </table>
  </div>

  {{ if .ComplianceSummary }}
  <div class="section">
    <h2>Compliance Violations</h2>
    <table>
      <thead>
        <tr><th>Framework</th><th>Control</th><th>Findings</th></tr>
      </thead>
      <tbody>
        {{ range $fw, $vs := .ComplianceSummary -}}
        {{ range $vs -}}
        <tr>
          <td><strong>{{ .Framework }}</strong></td>
          <td>{{ .Control }}</td>
          <td>{{ .Count }}</td>
        </tr>
        {{ end -}}
        {{ end -}}
      </tbody>
    </table>
  </div>
  {{ end }}

  {{ if criticalWithScript .Findings }}
  <div class="section">
    <h2>&#x1F6E1; Critical Finding Remediation Scripts</h2>
    <p style="color:var(--muted);font-size:0.875rem;margin-bottom:1rem;">
      The scripts below are auto-generated and resource-specific.
      Set <code>DRY_RUN=true</code> to preview without making changes.
    </p>
    {{ range criticalWithScript .Findings }}
    <div class="script-card">
      <h3>{{ .Category }}</h3>
      <div class="script-meta">
        <span class="badge badge-CRITICAL">CRITICAL</span>
        &nbsp;Risk Score: {{ riskScore . }}
        &nbsp;&middot;&nbsp;Resource: {{ .ResourceDisplayName }}
        &nbsp;&middot;&nbsp;Project: {{ .ProjectDisplayName }}
        {{ if .HasCVE }}&nbsp;&middot;&nbsp;CVE: {{ .CVEID }}{{ end }}
      </div>
      <details>
        <summary>View Remediation Script ({{ scriptLang . }})</summary>
        <pre class="script-pre"><code>{{ scriptBody . }}</code></pre>
      </details>
    </div>
    {{ end }}
  </div>
  {{ end }}

  <footer>
    Generated by <a href="https://github.com/wanaware/gcp-security-analyzer">gcp-security-analyzer</a>
  </footer>
</div>
<script>
function filterTable() {
  const input = document.getElementById('searchInput').value.toLowerCase();
  const rows = document.querySelectorAll('#findingsTable tbody tr');
  rows.forEach(row => {
    const text = row.textContent.toLowerCase();
    row.style.display = text.includes(input) ? '' : 'none';
  });
}
</script>
</body>
</html>`

// HTMLGenerator writes an interactive HTML report.
type HTMLGenerator struct{}

// NewHTMLGenerator creates a new HTMLGenerator.
func NewHTMLGenerator() *HTMLGenerator { return &HTMLGenerator{} }

// Generate writes an HTML report for r to w.
func (g *HTMLGenerator) Generate(r *models.Report, w io.Writer) error {
	if r.GeneratedAt.IsZero() {
		r.GeneratedAt = time.Now().UTC()
	}

	funcMap := template.FuncMap{
		"pct":              func(part, total int) string { return fmt.Sprintf("%.1f", utils.SafePercentage(part, total)) },
		"inc":              func(i int) int { return i + 1 },
		"riskScore":        func(f *models.Finding) string { return riskScoreStr(f) },
		"barWidth":         func(f *models.Finding) string { return barWidthStr(f) },
		"sortedCategories": sortedCategoryStats,
		// Critical remediation script helpers
		"criticalWithScript": criticalFindingsWithScript,
		"scriptLang": func(f *models.Finding) string {
			if f.Remediation != nil {
				return f.Remediation.RemediationScriptLang
			}
			return "bash"
		},
		// scriptBody returns the script as template.HTML so html/template does not
		// double-escape it — the content is already safely escaped by the template engine
		// when rendered inside <pre><code>{{ scriptBody . }}</code></pre>.
		"scriptBody": func(f *models.Finding) template.HTML {
			if f.Remediation == nil {
				return ""
			}
			return template.HTML(template.HTMLEscapeString(f.Remediation.RemediationScript))
		},
	}

	tmpl, err := template.New("html").Funcs(funcMap).Parse(htmlTemplate)
	if err != nil {
		return fmt.Errorf("parsing HTML template: %w", err)
	}

	// Sort findings by risk score for HTML table
	sorted := make([]*models.Finding, len(r.Findings))
	copy(sorted, r.Findings)
	sort.Slice(sorted, func(i, j int) bool {
		si, sj := 0.0, 0.0
		if sorted[i].RiskScore != nil {
			si = sorted[i].RiskScore.Total
		}
		if sorted[j].RiskScore != nil {
			sj = sorted[j].RiskScore.Total
		}
		return si > sj
	})

	// Use a temporary report with sorted findings
	rSorted := *r
	rSorted.Findings = sorted

	if err := tmpl.Execute(w, &rSorted); err != nil {
		return fmt.Errorf("executing HTML template: %w", err)
	}
	return nil
}

func barWidthStr(f *models.Finding) string {
	if f.RiskScore == nil {
		return "0"
	}
	return fmt.Sprintf("%.0f", f.RiskScore.Total)
}

// criticalFindingsWithScript returns CRITICAL findings that have a generated remediation script.
func criticalFindingsWithScript(findings []*models.Finding) []*models.Finding {
	var out []*models.Finding
	for _, f := range findings {
		if f.Priority == "CRITICAL" && f.Remediation != nil && f.Remediation.RemediationScript != "" {
			out = append(out, f)
		}
	}
	return out
}
