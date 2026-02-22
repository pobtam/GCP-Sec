package parser

import (
	"testing"
)

func TestBuildIndexAliases(t *testing.T) {
	// Simulates a GCP SCC export header with dotted API paths.
	header := []string{"finding.severity", "finding.category", "resource.name", "finding.name"}
	idx := buildIndex(header)

	// Verify canonical aliases were created.
	tests := []struct {
		canonical string
		wantIdx   int
	}{
		{"severity", 0},
		{"category", 1},
		{"resource_name", 2},
		{"name", 3},
	}

	for _, tc := range tests {
		got, ok := idx[tc.canonical]
		if !ok {
			t.Errorf("canonical %q not found in index", tc.canonical)
			continue
		}
		if got != tc.wantIdx {
			t.Errorf("idx[%q] = %d, want %d", tc.canonical, got, tc.wantIdx)
		}
	}

	// Verify originals are also still present.
	if _, ok := idx["finding.severity"]; !ok {
		t.Error("original 'finding.severity' should remain in index")
	}
}

func TestBuildIndexBareColumnTakesPrecedence(t *testing.T) {
	// Header has both bare "severity" and dotted "finding.severity".
	header := []string{"severity", "finding.severity"}
	idx := buildIndex(header)

	// "severity" canonical should point to index 0 (the bare column), not 1.
	got, ok := idx["severity"]
	if !ok {
		t.Fatal("'severity' not in index")
	}
	if got != 0 {
		t.Errorf("idx['severity'] = %d, want 0 (bare column should win)", got)
	}
}
