// Package compliance handles compliance framework detection and violation reporting.
package compliance

// Framework represents a supported compliance standard.
type Framework struct {
	ID          string
	Name        string
	Description string
	Aliases     []string // alternate names found in CSV data
}

// KnownFrameworks lists the supported compliance frameworks.
var KnownFrameworks = []Framework{
	{
		ID:          "CIS",
		Name:        "CIS Benchmarks",
		Description: "Center for Internet Security Benchmarks",
		Aliases:     []string{"CIS", "CIS_BENCHMARK", "CISGCP", "CIS_GCP"},
	},
	{
		ID:          "PCI-DSS",
		Name:        "PCI-DSS",
		Description: "Payment Card Industry Data Security Standard",
		Aliases:     []string{"PCI", "PCIDSS", "PCI_DSS", "PCI-DSS"},
	},
	{
		ID:          "HIPAA",
		Name:        "HIPAA",
		Description: "Health Insurance Portability and Accountability Act",
		Aliases:     []string{"HIPAA"},
	},
	{
		ID:          "SOC2",
		Name:        "SOC 2",
		Description: "Service Organization Control 2",
		Aliases:     []string{"SOC2", "SOC 2", "SOC_2"},
	},
	{
		ID:          "ISO27001",
		Name:        "ISO 27001",
		Description: "Information Security Management System standard",
		Aliases:     []string{"ISO27001", "ISO 27001", "ISO_27001"},
	},
	{
		ID:          "NIST",
		Name:        "NIST",
		Description: "NIST Cybersecurity Framework",
		Aliases:     []string{"NIST", "NIST_CSF", "NIST CSF", "NISTCSF"},
	},
	{
		ID:          "GDPR",
		Name:        "GDPR",
		Description: "General Data Protection Regulation",
		Aliases:     []string{"GDPR"},
	},
}

// aliasMap maps normalized alias strings to canonical framework IDs.
var aliasMap map[string]string

func init() {
	aliasMap = make(map[string]string)
	for _, fw := range KnownFrameworks {
		for _, alias := range fw.Aliases {
			aliasMap[normalize(alias)] = fw.ID
		}
	}
}

// normalize returns a simplified lowercase key suitable for map lookups.
func normalize(s string) string {
	out := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			out = append(out, c+32)
		} else if c != '-' && c != '_' && c != ' ' {
			out = append(out, c)
		}
	}
	return string(out)
}

// CanonicalID maps an observed framework string to its canonical ID.
// Returns the input string unchanged if no match is found.
func CanonicalID(observed string) string {
	if id, ok := aliasMap[normalize(observed)]; ok {
		return id
	}
	return observed
}

// FrameworkByID returns the Framework definition for a given canonical ID.
func FrameworkByID(id string) (Framework, bool) {
	for _, fw := range KnownFrameworks {
		if fw.ID == id {
			return fw, true
		}
	}
	return Framework{}, false
}
