package models

import "time"

// Report is the top-level structure holding all analysis results.
type Report struct {
	// Metadata
	GeneratedAt   time.Time `json:"generated_at"`
	InputFile     string    `json:"input_file"`
	TotalRows     int       `json:"total_rows"`
	ParseErrors   int       `json:"parse_errors"`

	// Findings
	Findings []*Finding `json:"findings"`

	// Statistics
	Stats ReportStats `json:"stats"`

	// Compliance summary
	ComplianceSummary map[string][]*ComplianceViolation `json:"compliance_summary,omitempty"`

	// Category breakdown
	CategoryBreakdown map[string]CategoryStats `json:"category_breakdown"`

	// Project breakdown
	ProjectBreakdown map[string]ProjectStats `json:"project_breakdown"`
}

// ReportStats holds aggregate statistics about the findings.
type ReportStats struct {
	Total    int `json:"total"`
	Critical int `json:"critical"`
	High     int `json:"high"`
	Medium   int `json:"medium"`
	Low      int `json:"low"`

	RiskStats RiskScoreStats `json:"risk_stats"`

	TopCategories []CategoryCount `json:"top_categories"`
	TopProjects   []ProjectCount  `json:"top_projects"`
}

// CategoryStats holds statistics for a specific finding category.
type CategoryStats struct {
	Category    string  `json:"category"`
	Count       int     `json:"count"`
	Critical    int     `json:"critical"`
	High        int     `json:"high"`
	Medium      int     `json:"medium"`
	Low         int     `json:"low"`
	AvgRiskScore float64 `json:"avg_risk_score"`
}

// ProjectStats holds statistics for a specific GCP project.
type ProjectStats struct {
	ProjectID  string  `json:"project_id"`
	ProjectName string `json:"project_name"`
	Count      int     `json:"count"`
	Critical   int     `json:"critical"`
	High       int     `json:"high"`
	Medium     int     `json:"medium"`
	Low        int     `json:"low"`
	AvgRiskScore float64 `json:"avg_risk_score"`
}

// CategoryCount is a simple count by category for top-N lists.
type CategoryCount struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

// ProjectCount is a simple count by project for top-N lists.
type ProjectCount struct {
	Project string `json:"project"`
	Count   int    `json:"count"`
}

// FilterOptions defines the criteria for filtering findings.
type FilterOptions struct {
	Priorities   []string
	Categories   []string
	Projects     []string
	MinRiskScore float64
	MaxRiskScore float64
}

// ReportOptions controls what is included in the generated report.
type ReportOptions struct {
	Format             string
	OutputPath         string
	OutputDir          string
	Formats            []string
	IncludeRemediation bool
	IncludeCompliance  bool
	SplitByPriority    bool
	Filter             FilterOptions
	ScoringConfigPath  string
}
