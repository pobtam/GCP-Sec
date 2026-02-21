# GCP Security Findings Analysis Report

Generated: 2026-02-21 19:08:25 UTC
Input File: findings.csv

---

## Executive Summary

- **Total Active Findings**: 3889
- **Critical Priority**: 0 (0.0%)
- **High Priority**: 0 (0.0%)
- **Medium Priority**: 61 (1.6%)
- **Low Priority**: 3828 (98.4%)

**Risk Score Statistics:**
- Mean: 18.65
- Median: 13.6
- Range: 9.6 - 43.6
- Std Dev: 8.19

**Top Risk Categories:**
1.  (3889 findings)


---

## Priority Breakdown

| Priority | Count | Percentage | Avg Risk Score | Remediation SLA |
|----------|------:|----------:|---------------:|-----------------|
| CRITICAL | 0 | 0.0% | N/A | 24-48 hours |
| HIGH | 0 | 0.0% | N/A | 1 week |
| MEDIUM | 61 | 1.6% | 35.4 | 30 days |
| LOW | 3828 | 98.4% | 18.4 | 90 days |

---

## Risk Scoring Methodology

Findings are scored on a 0-100 scale using the following components:

| Component | Max Points | Description |
|-----------|----------:|-------------|
| Base Severity | 40 | CRITICAL=40, HIGH=30, MEDIUM=20, LOW=10 |
| CVSS Score | 30 | CVSS v3 base score × 3 |
| Exploitability | 20 | In-the-wild (+10), Zero-day (+8), Activity level (+2-6) |
| Finding Class | 10 | THREAT=10, VULNERABILITY=7, MISCONFIG=5, OBSERVATION=2 |
| Resource Exposure | 10 | Public IP (+5), Internet-facing (+3), Critical resource (+2) |
| Compliance Impact | 10 | Has frameworks (+5), Details (+3), Audit category (+2) |
| Category Weight | ×0.8–1.2 | High-risk categories get 1.2× multiplier |

**Priority Thresholds:** CRITICAL ≥75 | HIGH 55–74 | MEDIUM 35–54 | LOW <35

---

## Top Findings by Priority

### MEDIUM Priority (61 findings)

| # | Category | Resource | Risk Score | CVE | Project |
|--:|----------|----------|----------:|-----|---------|
| 1 |  |  | 35.20 | CVE-2025-68121 |  |
| 2 |  |  | 35.20 | CVE-2025-68121 |  |
| 3 |  |  | 35.20 | CVE-2025-68121 |  |
| 4 |  |  | 35.20 | CVE-2025-68121 |  |
| 5 |  |  | 35.20 | CVE-2025-68121 |  |
| 6 |  |  | 35.20 | CVE-2025-68121 |  |
| 7 |  |  | 35.20 | CVE-2025-68121 |  |
| 8 |  |  | 35.20 | CVE-2025-68121 |  |
| 9 |  |  | 35.20 | CVE-2025-68121 |  |
| 10 |  |  | 35.20 | CVE-2025-68121 |  |
| 11 |  |  | 35.20 | CVE-2025-68121 |  |
| 12 |  |  | 35.20 | CVE-2025-68121 |  |
| 13 |  |  | 35.20 | CVE-2025-68121 |  |
| 14 |  |  | 43.60 | CVE-2023-44487 |  |
| 15 |  |  | 35.20 | CVE-2025-68121 |  |
| 16 |  |  | 35.20 | CVE-2025-68121 |  |
| 17 |  |  | 35.20 | CVE-2025-68121 |  |
| 18 |  |  | 35.20 | CVE-2025-68121 |  |
| 19 |  |  | 35.20 | CVE-2025-68121 |  |
| 20 |  |  | 35.20 | CVE-2025-68121 |  |

### LOW Priority (3828 findings)

| # | Category | Resource | Risk Score | CVE | Project |
|--:|----------|----------|----------:|-----|---------|
| 1 |  |  | 12.00 | - |  |
| 2 |  |  | 12.00 | - |  |
| 3 |  |  | 12.00 | - |  |
| 4 |  |  | 12.00 | - |  |
| 5 |  |  | 12.00 | - |  |
| 6 |  |  | 12.00 | - |  |
| 7 |  |  | 12.00 | - |  |
| 8 |  |  | 12.00 | - |  |
| 9 |  |  | 12.00 | - |  |
| 10 |  |  | 12.00 | - |  |
| 11 |  |  | 12.00 | - |  |
| 12 |  |  | 12.00 | - |  |
| 13 |  |  | 12.00 | - |  |
| 14 |  |  | 12.00 | - |  |
| 15 |  |  | 12.00 | - |  |
| 16 |  |  | 12.00 | - |  |
| 17 |  |  | 12.00 | - |  |
| 18 |  |  | 12.00 | - |  |
| 19 |  |  | 12.00 | - |  |
| 20 |  |  | 12.00 | - |  |


---

## Compliance Framework Violations

### CCM

- **Total Violations**: 8
- **Top Violations**:
- CCM LOG-08: 142 finding(s)
- CCM DSP-17: 97 finding(s)
- CCM LOG-05: 38 finding(s)
- CCM CEK-03: 21 finding(s)
- CCM DSP-07: 18 finding(s)

### CIS

- **Total Violations**: 62
- **Top Violations**:
- CIS 3.8: 958 finding(s)
- CIS 3.9: 524 finding(s)
- CIS 1.4: 91 finding(s)
- CIS 1.3: 85 finding(s)
- CIS 1.6: 75 finding(s)

### CIS-CONTROLS

- **Total Violations**: 19
- **Top Violations**:
- CIS-CONTROLS 8.2: 180 finding(s)
- CIS-CONTROLS 8.5: 140 finding(s)
- CIS-CONTROLS 3.3: 61 finding(s)
- CIS-CONTROLS 8.11: 38 finding(s)
- CIS-CONTROLS 5.6: 32 finding(s)

### HIPAA

- **Total Violations**: 10
- **Top Violations**:
- HIPAA 164.312(b): 180 finding(s)
- HIPAA 164.312(a)(1): 61 finding(s)
- HIPAA 164.308(a)(3)(i): 61 finding(s)
- HIPAA 164.308(a)(3)(ii): 61 finding(s)
- HIPAA 164.308(a)(1)(ii): 38 finding(s)

### ISO

- **Total Violations**: 10
- **Top Violations**:
- ISO A.13.1.1: 630 finding(s)
- ISO A.14.1.3: 56 finding(s)
- ISO A.9.2.3: 32 finding(s)
- ISO A.12.4.1: 19 finding(s)
- ISO A.16.1.7: 19 finding(s)

### ISO27001

- **Total Violations**: 17
- **Top Violations**:
- ISO27001 A.8.15: 178 finding(s)
- ISO27001 A.8.20: 142 finding(s)
- ISO27001 A.5.15: 99 finding(s)
- ISO27001 A.8.4: 61 finding(s)
- ISO27001 A.8.3: 61 finding(s)

### NIST

- **Total Violations**: 21
- **Top Violations**:
- NIST SI-4: 571 finding(s)
- NIST PR-PT-1: 180 finding(s)
- NIST DE-AE-3: 178 finding(s)
- NIST SC-7: 138 finding(s)
- NIST PR-AC-4: 67 finding(s)

### NIST800-53

- **Total Violations**: 26
- **Top Violations**:
- NIST800-53 AU-7: 216 finding(s)
- NIST800-53 AU-12: 178 finding(s)
- NIST800-53 AU-2: 142 finding(s)
- NIST800-53 AC-6: 67 finding(s)
- NIST800-53 AC-3: 61 finding(s)

### PCI-DSS

- **Total Violations**: 47
- **Top Violations**:
- PCI-DSS 10.1: 590 finding(s)
- PCI-DSS 10.2: 590 finding(s)
- PCI-DSS 10.2.1.5: 178 finding(s)
- PCI-DSS 10.2.1: 178 finding(s)
- PCI-DSS 10.2.1.2: 178 finding(s)

### SOC2

- **Total Violations**: 32
- **Top Violations**:
- SOC2 CC5.2.3: 97 finding(s)
- SOC2 CC6.1.3: 88 finding(s)
- SOC2 CC6.1.7: 67 finding(s)
- SOC2 CC6.1.8: 59 finding(s)
- SOC2 CC4.1.7: 38 finding(s)


---

## Category Breakdown

| Category | Total | Critical | High | Medium | Low | Avg Score |
|----------|------:|---------:|-----:|-------:|----:|----------:|
|  | 3889 | 0 | 0 | 61 | 3828 | 18.6 |


---

## Project Breakdown

| Project | Total | Critical | High | Medium | Low | Avg Score |
|---------|------:|---------:|-----:|-------:|----:|----------:|


---

## Remediation Actions


*Report generated by [gcp-security-analyzer](https://github.com/wanaware/gcp-security-analyzer)*
