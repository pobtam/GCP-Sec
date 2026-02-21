# Risk Scoring Methodology

## Overview

The GCP Security Analyzer uses a multi-factor risk scoring system to assign each finding a numeric risk score between 0 and 100. This score is then mapped to a priority level (CRITICAL, HIGH, MEDIUM, LOW) that drives remediation scheduling.

The scoring system is intentionally transparent and configurable — every component is broken out in the JSON output so security teams can audit and adjust the weights.

---

## Scoring Components

### 1. Base Severity Score (0–40 points)

The severity assigned by GCP Security Command Center forms the foundation of the risk score.

| Severity | Points |
|----------|-------:|
| CRITICAL | 40 |
| HIGH     | 30 |
| MEDIUM   | 20 |
| LOW      | 10 |

GCP's severity classification reflects the potential impact of the finding, so it carries the highest individual weight.

---

### 2. CVSS Score Component (0–30 points)

For vulnerability findings, the CVSSv3 base score is extracted from the `finding.vulnerability.cve.cvssv3` JSON field. The 0–10 CVSS scale is linearly mapped to 0–30 by multiplying by 3.

```
cvss_component = min(cvss_base_score × 3, 30)
```

CVSS provides an industry-standard assessment of exploitability and impact, making it a reliable complement to GCP's severity classification.

**Field:** `finding.vulnerability.cve.cvssv3.base_score`

---

### 3. Exploitability Score (0–20 points)

Active exploitation dramatically increases the urgency of remediation. This component rewards findings with evidence of real-world exploitation:

| Condition | Points |
|-----------|-------:|
| Observed in the wild (`observed_in_the_wild = true`) | +10 |
| Zero-day vulnerability (`zero_day = true`)           | +8  |
| Exploitation activity = HIGH or ACTIVE               | +6  |
| Exploitation activity = MEDIUM or PoC               | +4  |
| Exploitation activity = LOW                          | +2  |
| Has CVE identifier                                   | +2  |

Points are additive but capped at 20.

**Fields:**
- `finding.vulnerability.cve.observed_in_the_wild`
- `finding.vulnerability.cve.zero_day`
- `finding.vulnerability.cve.exploitation_activity`
- `finding.vulnerability.cve.id`

---

### 4. Finding Class Modifier (0–10 points)

The finding class reflects the nature of the security issue:

| Finding Class    | Points |
|------------------|-------:|
| THREAT           | 10 |
| VULNERABILITY    | 7  |
| MISCONFIGURATION | 5  |
| OBSERVATION      | 2  |

Threats represent active attack activity and receive maximum weight. Misconfigurations, while important, reflect potential risk rather than active exploitation.

**Field:** `finding_class`

---

### 5. Resource Exposure Score (0–10 bonus points)

The attack surface of the affected resource increases risk. Internet-facing and high-value resources receive additional weight:

| Condition | Points |
|-----------|-------:|
| Resource has a public IP address | +5 |
| Internet-facing service (HTTP, Load Balancer in category) | +3 |
| Critical resource type (Database, Storage, Bucket) | +2 |

Points are additive but capped at 10.

**Fields:**
- `finding.external_exposure.public_ip_address`
- `category` (examined for keywords)

---

### 6. Compliance Impact (0–10 bonus points)

Findings that violate regulatory compliance frameworks carry additional organizational risk:

| Condition | Points |
|-----------|-------:|
| Has compliance frameworks (`finding.compliances` not empty) | +5 |
| Has detailed compliance information (`finding.compliance_details.frameworks` not empty) | +3 |
| Audit/logging/monitoring category | +2 |

Points are additive but capped at 10.

**Fields:**
- `finding.compliances`
- `finding.compliance_details.frameworks`
- `category`

---

### 7. Category Risk Weight (multiplier: 0.8–1.2×)

A category-based multiplier is applied to the raw sum of all components. This reflects that some types of findings inherently carry higher systematic risk:

| Risk Level | Categories | Multiplier |
|------------|-----------|----------:|
| High       | VULNERABILITY, PRIVILEGE_ESCALATION, WEAK_SSL, PUBLIC, EXTERNAL, CONTAINER_IMAGE_VULN | 1.2× |
| Medium     | AUDIT_LOGGING, FLOW_LOGS, BUCKET_LOGGING, SERVICE_ACCOUNT_KEY, FIREWALL, IAM, NETWORK | 1.0× |
| Low        | All others (misconfigurations, observations) | 0.8× |

**Field:** `category`

---

## Final Formula

```
raw_score = base_severity + cvss_component + exploitability + class_modifier + exposure_score + compliance_score

risk_score = min(raw_score × category_weight, 100)
```

---

## Priority Assignment

| Priority | Risk Score Range | Remediation SLA |
|----------|-----------------|-----------------|
| CRITICAL | 75–100           | 24–48 hours     |
| HIGH     | 55–74            | 1 week          |
| MEDIUM   | 35–54            | 30 days         |
| LOW      | 0–34             | 90 days         |

---

## Example Calculations

### Example 1: Critical Log4Shell Finding

| Component | Calculation | Points |
|-----------|-------------|-------:|
| Base Severity (HIGH) | — | 30 |
| CVSS 10.0 | 10.0 × 3 | 30 |
| Exploitability | In-wild(+10) + CVE(+2) | 12 |
| Class (VULNERABILITY) | — | 7 |
| Exposure | No public IP | 0 |
| Compliance | CIS framework | 5 |
| **Raw sum** | | **84** |
| Category Weight | VULNERABILITY × 1.2 | ×1.2 |
| **Final Score** | min(84 × 1.2, 100) | **100** → CRITICAL |

### Example 2: Flow Logs Disabled

| Component | Calculation | Points |
|-----------|-------------|-------:|
| Base Severity (MEDIUM) | — | 20 |
| CVSS | No CVE | 0 |
| Exploitability | No CVE | 0 |
| Class (MISCONFIGURATION) | — | 5 |
| Exposure | No public IP | 0 |
| Compliance | CIS framework | 5 |
| **Raw sum** | | **30** |
| Category Weight | FLOW_LOGS × 1.0 | ×1.0 |
| **Final Score** | min(30 × 1.0, 100) | **30** → LOW |

---

## Customizing Scoring Weights

The scoring weights can be adjusted via a YAML configuration file:

```yaml
# scoring-config.yaml
critical_base: 40
high_base: 30
medium_base: 20
low_base: 10
cvss_max: 30
exploit_max: 20
exposure_max: 10
compliance_max: 10
high_risk_weight: 1.2
med_risk_weight: 1.0
low_risk_weight: 0.8
```

Pass it with `--scoring-config scoring-config.yaml`.

---

## Limitations and Caveats

1. **Severity is GCP-assigned**: The base severity reflects GCP's classification, not a universal standard. Different GCP detectors may assign severities differently.

2. **CVSS availability**: Many misconfigurations and observations do not have CVE identifiers or CVSS scores. Their score is driven primarily by severity, class, and category.

3. **Compliance coverage**: Compliance impact is based on metadata in the CSV export. Incomplete compliance mappings in GCP may undercount violations.

4. **Context not captured**: The tool cannot evaluate business context (e.g., whether an internet-facing service is intentionally public or whether a test environment is lower risk than production). Apply judgment when reviewing priorities.

5. **Score is relative**: The 0–100 score enables comparison and prioritization within a dataset, but should not be treated as an absolute severity measure comparable across different organizations.
