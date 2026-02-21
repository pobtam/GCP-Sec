// Package models defines the core data structures for the GCP security analyzer.
package models

import "time"

// Finding represents a single GCP Security Command Center finding parsed from CSV.
type Finding struct {
	// Core identification
	Name                string `csv:"name" json:"name"`
	FindingClass        string `csv:"finding_class" json:"finding_class"`
	FindingType         string `csv:"finding_type" json:"finding_type"`
	Category            string `csv:"category" json:"category"`
	State               string `csv:"state" json:"state"`
	Severity            string `csv:"severity" json:"severity"`

	// Resource information
	ResourceName        string `csv:"resource_name" json:"resource_name"`
	ResourceDisplayName string `csv:"resource_display_name" json:"resource_display_name"`
	ResourceType        string `csv:"resource_type" json:"resource_type"`
	ProjectDisplayName  string `csv:"project_display_name" json:"project_display_name"`
	ProjectID           string `csv:"project_id" json:"project_id"`

	// Timestamps
	EventTime  time.Time `csv:"-" json:"event_time"`
	CreateTime time.Time `csv:"-" json:"create_time"`
	EventTimeRaw  string `csv:"event_time" json:"-"`
	CreateTimeRaw string `csv:"create_time" json:"-"`

	// Description and remediation
	Description string `csv:"description" json:"description"`
	NextSteps   string `csv:"finding.next_steps" json:"next_steps"`
	ExternalURI string `csv:"external_uri" json:"external_uri"`

	// Vulnerability data (raw JSON strings from CSV)
	CVSSv3Raw             string `csv:"finding.vulnerability.cve.cvssv3" json:"-"`
	CVEIDRaw              string `csv:"finding.vulnerability.cve.id" json:"-"`
	ObservedInWildRaw     string `csv:"finding.vulnerability.cve.observed_in_the_wild" json:"-"`
	ZeroDayRaw            string `csv:"finding.vulnerability.cve.zero_day" json:"-"`
	ExploitabilityRaw     string `csv:"finding.vulnerability.cve.exploitation_activity" json:"-"`

	// Parsed vulnerability data
	CVEID             string  `csv:"-" json:"cve_id,omitempty"`
	CVSSScore         float64 `csv:"-" json:"cvss_score,omitempty"`
	ObservedInWild    bool    `csv:"-" json:"observed_in_wild"`
	ZeroDay           bool    `csv:"-" json:"zero_day"`
	ExploitActivity   string  `csv:"-" json:"exploit_activity,omitempty"`

	// Exposure data (raw JSON)
	PublicIPAddressRaw string `csv:"finding.external_exposure.public_ip_address" json:"-"`
	PublicIPAddress    bool   `csv:"-" json:"public_ip_address"`

	// Compliance data (raw JSON strings)
	CompliancesRaw         string `csv:"finding.compliances" json:"-"`
	ComplianceDetailsRaw   string `csv:"finding.compliance_details.frameworks" json:"-"`
	CloudControlRaw        string `csv:"finding.compliance_details.cloud_control" json:"-"`

	// Parsed compliance data
	ComplianceFrameworks []string `csv:"-" json:"compliance_frameworks,omitempty"`
	ComplianceDetails    []ComplianceDetail `csv:"-" json:"compliance_details,omitempty"`

	// Labels/tags
	ResourceLabelsRaw string `csv:"resource.labels" json:"-"`
	FindingLabelsRaw  string `csv:"finding.source_properties" json:"-"`

	// Computed fields (populated during analysis)
	RiskScore    *RiskScore          `csv:"-" json:"risk_score,omitempty"`
	Priority     string              `csv:"-" json:"priority,omitempty"`
	Remediation  *RemediationStep    `csv:"-" json:"remediation,omitempty"`
	Violations   []ComplianceViolation `csv:"-" json:"violations,omitempty"`
}

// ComplianceDetail holds parsed compliance framework details.
type ComplianceDetail struct {
	Framework string `json:"framework"`
	Version   string `json:"version,omitempty"`
	Controls  []string `json:"controls,omitempty"`
}

// SeverityScore returns the numeric severity level for sorting.
func (f *Finding) SeverityScore() int {
	switch f.Severity {
	case "CRITICAL":
		return 4
	case "HIGH":
		return 3
	case "MEDIUM":
		return 2
	case "LOW":
		return 1
	default:
		return 0
	}
}

// IsActive returns true if the finding is in an active state.
func (f *Finding) IsActive() bool {
	return f.State == "ACTIVE"
}

// HasCVE returns true if the finding has a CVE identifier.
func (f *Finding) HasCVE() bool {
	return f.CVEID != ""
}

// ShortName returns the last component of the finding name path.
func (f *Finding) ShortName() string {
	if f.Name == "" {
		return ""
	}
	for i := len(f.Name) - 1; i >= 0; i-- {
		if f.Name[i] == '/' {
			return f.Name[i+1:]
		}
	}
	return f.Name
}
