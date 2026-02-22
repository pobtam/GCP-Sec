package fetcher

import (
	"testing"

	"cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func makeResult(
	name, category, description string,
	severity securitycenterpb.Finding_Severity,
	state securitycenterpb.Finding_State,
	class securitycenterpb.Finding_FindingClass,
) *securitycenterpb.ListFindingsResponse_ListFindingsResult {
	return &securitycenterpb.ListFindingsResponse_ListFindingsResult{
		Finding: &securitycenterpb.Finding{
			Name:         name,
			Category:     category,
			Description:  description,
			Severity:     severity,
			State:        state,
			FindingClass: class,
			EventTime:    timestamppb.Now(),
			CreateTime:   timestamppb.Now(),
		},
		Resource: &securitycenterpb.ListFindingsResponse_ListFindingsResult_Resource{
			Name:               "//storage.googleapis.com/prod-bucket",
			DisplayName:        "prod-bucket",
			Type:               "google.storage.Bucket",
			ProjectName:        "projects/prod-project",
			ProjectDisplayName: "Production Project",
		},
	}
}

func TestConvertFinding_BasicFields(t *testing.T) {
	result := makeResult(
		"organizations/123/sources/456/findings/abc",
		"PUBLIC_BUCKET_ACL",
		"Public bucket found.",
		securitycenterpb.Finding_CRITICAL,
		securitycenterpb.Finding_ACTIVE,
		securitycenterpb.Finding_MISCONFIGURATION,
	)

	f := ConvertFinding(result)

	if f.Name != "organizations/123/sources/456/findings/abc" {
		t.Errorf("Name = %q, want organizations/123/sources/456/findings/abc", f.Name)
	}
	if f.Category != "PUBLIC_BUCKET_ACL" {
		t.Errorf("Category = %q, want PUBLIC_BUCKET_ACL", f.Category)
	}
	if f.Severity != "CRITICAL" {
		t.Errorf("Severity = %q, want CRITICAL", f.Severity)
	}
	if f.State != "ACTIVE" {
		t.Errorf("State = %q, want ACTIVE", f.State)
	}
	if f.FindingClass != "MISCONFIGURATION" {
		t.Errorf("FindingClass = %q, want MISCONFIGURATION", f.FindingClass)
	}
	if f.Description != "Public bucket found." {
		t.Errorf("Description = %q, want 'Public bucket found.'", f.Description)
	}

	// Resource fields
	if f.ResourceName != "//storage.googleapis.com/prod-bucket" {
		t.Errorf("ResourceName = %q, unexpected", f.ResourceName)
	}
	if f.ResourceDisplayName != "prod-bucket" {
		t.Errorf("ResourceDisplayName = %q, want prod-bucket", f.ResourceDisplayName)
	}
	if f.ProjectID != "prod-project" {
		t.Errorf("ProjectID = %q, want prod-project", f.ProjectID)
	}
	if f.ProjectDisplayName != "Production Project" {
		t.Errorf("ProjectDisplayName = %q, want 'Production Project'", f.ProjectDisplayName)
	}

	// Timestamps should be populated
	if f.EventTime.IsZero() {
		t.Error("EventTime should not be zero")
	}
	if f.EventTimeRaw == "" {
		t.Error("EventTimeRaw should not be empty")
	}
}

func TestConvertFinding_Vulnerability(t *testing.T) {
	result := makeResult(
		"organizations/123/sources/456/findings/vuln1",
		"CONTAINER_IMAGE_VULNERABILITY",
		"CVE found.",
		securitycenterpb.Finding_HIGH,
		securitycenterpb.Finding_ACTIVE,
		securitycenterpb.Finding_VULNERABILITY,
	)
	result.Finding.Vulnerability = &securitycenterpb.Vulnerability{
		Cve: &securitycenterpb.Cve{
			Id:                  "CVE-2021-44228",
			ObservedInTheWild:   true,
			ZeroDay:             false,
			ExploitationActivity: securitycenterpb.Cve_WIDE,
			Cvssv3: &securitycenterpb.Cvssv3{
				BaseScore: 10.0,
			},
		},
	}

	f := ConvertFinding(result)

	if f.CVEID != "CVE-2021-44228" {
		t.Errorf("CVEID = %q, want CVE-2021-44228", f.CVEID)
	}
	if f.CVSSScore != 10.0 {
		t.Errorf("CVSSScore = %.1f, want 10.0", f.CVSSScore)
	}
	if !f.ObservedInWild {
		t.Error("ObservedInWild should be true")
	}
	if f.ZeroDay {
		t.Error("ZeroDay should be false")
	}
	if f.ExploitActivity != "ACTIVE" {
		t.Errorf("ExploitActivity = %q, want ACTIVE", f.ExploitActivity)
	}
	if f.Severity != "HIGH" {
		t.Errorf("Severity = %q, want HIGH", f.Severity)
	}
}

func TestConvertFinding_Compliance(t *testing.T) {
	result := makeResult(
		"organizations/123/sources/456/findings/comp1",
		"FLOW_LOGS_DISABLED",
		"Flow logs disabled.",
		securitycenterpb.Finding_MEDIUM,
		securitycenterpb.Finding_ACTIVE,
		securitycenterpb.Finding_MISCONFIGURATION,
	)
	result.Finding.Compliances = []*securitycenterpb.Compliance{
		{Standard: "CIS", Version: "1.2", Ids: []string{"3.9"}},
		{Standard: "PCI-DSS", Version: "3.2", Ids: []string{"6.2", "10.1"}},
	}

	f := ConvertFinding(result)

	if len(f.ComplianceFrameworks) != 2 {
		t.Fatalf("ComplianceFrameworks len = %d, want 2", len(f.ComplianceFrameworks))
	}
	if f.ComplianceFrameworks[0] != "CIS" {
		t.Errorf("ComplianceFrameworks[0] = %q, want CIS", f.ComplianceFrameworks[0])
	}
	if f.ComplianceFrameworks[1] != "PCI-DSS" {
		t.Errorf("ComplianceFrameworks[1] = %q, want PCI-DSS", f.ComplianceFrameworks[1])
	}
	if f.CompliancesRaw == "" {
		t.Error("CompliancesRaw should not be empty")
	}
}

func TestConvertFinding_NilFields(t *testing.T) {
	// Minimal result with no resource, no vulnerability, no compliance.
	result := &securitycenterpb.ListFindingsResponse_ListFindingsResult{
		Finding: &securitycenterpb.Finding{
			Name:     "organizations/123/sources/456/findings/minimal",
			Category: "SOME_CATEGORY",
			Severity: securitycenterpb.Finding_LOW,
			State:    securitycenterpb.Finding_ACTIVE,
		},
	}

	f := ConvertFinding(result)

	if f.Name != "organizations/123/sources/456/findings/minimal" {
		t.Errorf("Name = %q, unexpected", f.Name)
	}
	if f.Severity != "LOW" {
		t.Errorf("Severity = %q, want LOW", f.Severity)
	}
	if f.ResourceName != "" {
		t.Errorf("ResourceName should be empty, got %q", f.ResourceName)
	}
	if f.CVEID != "" {
		t.Errorf("CVEID should be empty, got %q", f.CVEID)
	}
	if len(f.ComplianceFrameworks) != 0 {
		t.Errorf("ComplianceFrameworks should be empty, got %v", f.ComplianceFrameworks)
	}
}

func TestSanitizeEnum(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"CRITICAL", "CRITICAL"},
		{"SEVERITY_UNSPECIFIED", ""},
		{"FINDING_CLASS_VULNERABILITY", "VULNERABILITY"},
		{"STATE_ACTIVE", "ACTIVE"},
		{"EXPLOITATION_ACTIVITY_WIDE", "WIDE"},
		{"UNSPECIFIED", ""},
		{"", ""},
		{"THREAT", "THREAT"},
	}
	for _, tc := range tests {
		got := sanitizeEnum(tc.input)
		if got != tc.want {
			t.Errorf("sanitizeEnum(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestSanitizeSeverity(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"CRITICAL", "CRITICAL"},
		{"HIGH", "HIGH"},
		{"SEVERITY_UNSPECIFIED", "LOW"},
		{"", "LOW"},
	}
	for _, tc := range tests {
		got := sanitizeSeverity(tc.input)
		if got != tc.want {
			t.Errorf("sanitizeSeverity(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestSanitizeExploitActivity(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"WIDE", "ACTIVE"},
		{"CONFIRMED", "ACTIVE"},
		{"AVAILABLE", "POC"},
		{"ANTICIPATED", "LOW"},
		{"NO_KNOWN", ""},
		{"EXPLOITATION_ACTIVITY_WIDE", "ACTIVE"},
		{"", ""},
	}
	for _, tc := range tests {
		got := sanitizeExploitActivity(tc.input)
		if got != tc.want {
			t.Errorf("sanitizeExploitActivity(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestExtractProjectID(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"projects/my-project", "my-project"},
		{"organizations/123/projects/prod", "prod"},
		{"", ""},
		{"some-other-format", "some-other-format"},
	}
	for _, tc := range tests {
		got := extractProjectID(tc.input)
		if got != tc.want {
			t.Errorf("extractProjectID(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

func TestSerializeCompliances(t *testing.T) {
	cs := []*securitycenterpb.Compliance{
		{Standard: "CIS", Version: "1.2", Ids: []string{"3.1", "3.2"}},
	}
	raw := serializeCompliances(cs)
	if raw == "" {
		t.Fatal("serializeCompliances returned empty string")
	}
	want := `[{"standard":"CIS","version":"1.2","ids":["3.1","3.2"]}]`
	if raw != want {
		t.Errorf("serializeCompliances = %q, want %q", raw, want)
	}
}
