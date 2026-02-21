package models

// RiskScore holds the component breakdown of a finding's risk score.
type RiskScore struct {
	// Component scores
	BaseSeverity    float64 `json:"base_severity"`
	CVSSComponent   float64 `json:"cvss_component"`
	Exploitability  float64 `json:"exploitability"`
	ClassModifier   float64 `json:"class_modifier"`
	ExposureScore   float64 `json:"exposure_score"`
	ComplianceScore float64 `json:"compliance_score"`
	CategoryWeight  float64 `json:"category_weight"`

	// Final score
	Total float64 `json:"total"`

	// Explanation
	Rationale string `json:"rationale,omitempty"`
}

// RiskScoreStats holds aggregated statistics for a collection of risk scores.
type RiskScoreStats struct {
	Count  int     `json:"count"`
	Mean   float64 `json:"mean"`
	Median float64 `json:"median"`
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
	StdDev float64 `json:"std_dev"`
}

// ComplianceViolation represents a specific compliance framework violation.
type ComplianceViolation struct {
	Framework   string   `json:"framework"`
	Version     string   `json:"version,omitempty"`
	Control     string   `json:"control"`
	Description string   `json:"description,omitempty"`
	Findings    []string `json:"findings,omitempty"` // Finding names
	Count       int      `json:"count"`
}

// RemediationStep provides structured guidance for fixing a finding.
type RemediationStep struct {
	Summary             string   `json:"summary"`
	NextSteps           []string `json:"next_steps"`
	ResourceLinks       []string `json:"resource_links,omitempty"`
	EstimatedEffort     string   `json:"estimated_effort"`
	AutomationPotential string   `json:"automation_potential"`
	AutomationHint      string   `json:"automation_hint,omitempty"`
	PriorityRationale   string   `json:"priority_rationale"`

	// Finding-specific remediation script — only populated for CRITICAL findings.
	RemediationScript     string `json:"remediation_script,omitempty"`
	RemediationScriptLang string `json:"remediation_script_lang,omitempty"` // "bash" or "python3"
}

// Priority levels for findings.
const (
	PriorityCritical = "CRITICAL"
	PriorityHigh     = "HIGH"
	PriorityMedium   = "MEDIUM"
	PriorityLow      = "LOW"
)

// PriorityThresholds defines the score ranges for each priority level.
var PriorityThresholds = map[string][2]float64{
	PriorityCritical: {75, 100},
	PriorityHigh:     {55, 74},
	PriorityMedium:   {35, 54},
	PriorityLow:      {0, 34},
}

// RemediationSLA maps priority levels to remediation SLA strings.
var RemediationSLA = map[string]string{
	PriorityCritical: "24-48 hours",
	PriorityHigh:     "1 week",
	PriorityMedium:   "30 days",
	PriorityLow:      "90 days",
}

// PriorityFromScore returns the priority level for a given risk score.
func PriorityFromScore(score float64) string {
	switch {
	case score >= 75:
		return PriorityCritical
	case score >= 55:
		return PriorityHigh
	case score >= 35:
		return PriorityMedium
	default:
		return PriorityLow
	}
}
