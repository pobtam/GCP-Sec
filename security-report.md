# GCP Security Findings Analysis Report

Generated: 2026-02-22 05:05:44 UTC
Input File: GCP SCC org:320438219353 (last 4 days)

---

## Executive Summary

- **Total Active Findings**: 713
- **Critical Priority**: 199 (27.9%)
- **High Priority**: 364 (51.1%)
- **Medium Priority**: 0 (0.0%)
- **Low Priority**: 150 (21.0%)

**Risk Score Statistics:**
- Mean: 62.87
- Median: 73.8
- Range: 16 - 98.88
- Std Dev: 23.75

**Top Risk Categories:**
1. SOFTWARE_VULNERABILITY (338 findings)
2. OS_VULNERABILITY (95 findings)
3. PRIVATE_GOOGLE_ACCESS_DISABLED (42 findings)
4. FLOW_LOGS_DISABLED (42 findings)
5. POD THAT EXPOSES MANY VALUED RESOURCES (25 findings)
6. VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED (18 findings)
7. USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BUCKET OBJECTS. (13 findings)
8. FIREWALL THAT EXPOSES MANY VALUED RESOURCES (12 findings)
9. PERSISTENCE: NEW API METHOD (9 findings)
10. SERVICE ACCOUNT THAT EXPOSES MANY VALUED RESOURCES (8 findings)


---

## Priority Breakdown

| Priority | Count | Percentage | Avg Risk Score | Remediation SLA |
|----------|------:|----------:|---------------:|-----------------|
| CRITICAL | 199 | 27.9% | 81.8 | 24-48 hours |
| HIGH | 364 | 51.1% | 70.2 | 1 week |
| MEDIUM | 0 | 0.0% | N/A | 30 days |
| LOW | 150 | 21.0% | 19.9 | 90 days |

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

### CRITICAL Priority (199 findings)

| # | Category | Resource | Risk Score | CVE | Project |
|--:|----------|----------|----------:|-----|---------|
| 1 | POD THAT EXPOSES MANY VALUED RESOURCES | konnectivity-agent-autoscaler-d765574... | 75.00 | - | wanaware-core-prod |
| 2 | POD THAT EXPOSES MANY VALUED RESOURCES | fleet-manager-deployment-84f6db59df-t... | 75.00 | - | wanaware-core-prod |
| 3 | PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY. | owasp-test | 75.00 | - | wanaware-security-map-dev |
| 4 | POD THAT EXPOSES MANY VALUED RESOURCES | chisel-6fb5b877c6-4krgq | 75.00 | - | wanaware-core-prod |
| 5 | POD THAT EXPOSES MANY VALUED RESOURCES | reseller-wholesale-worker-deployment-... | 75.00 | - | wanaware-core-prod |
| 6 | POD THAT EXPOSES MANY VALUED RESOURCES | l7-default-backend-78858cccc9-7mzlq | 75.00 | - | wanaware-core-prod |
| 7 | SERVICE ACCOUNT THAT EXPOSES MANY VALUED RESOURCES | projects/wanaware-core-prod/serviceAc... | 75.00 | - | wanaware-core-prod |
| 8 | FIREWALL THAT EXPOSES MANY VALUED RESOURCES | wanaware-vpc-allow-https | 75.00 | - | wanaware-core-dev |
| 9 | FIREWALL THAT EXPOSES MANY VALUED RESOURCES | default-allow-rdp | 75.00 | - | wa-gcp-test |
| 10 | POD THAT EXPOSES MANY VALUED RESOURCES | task-scheduler-deployment-8477b65488-... | 75.00 | - | wanaware-core-prod |
| 11 | POD THAT EXPOSES MANY VALUED RESOURCES | postgres-client-578d77d76c-h62mx | 75.00 | - | wanaware-core-prod |
| 12 | FIREWALL THAT EXPOSES MANY VALUED RESOURCES | default-allow-rdp | 75.00 | - | wanaware-core-dev |
| 13 | SERVICE ACCOUNT KEY THAT EXPOSES MANY VALUED RESOURCES | projects/wanaware-core-prod/serviceAc... | 75.00 | - | wanaware-core-prod |
| 14 | USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO ASSUME OTHER SERVICE ACCOUNTS. | projects/wanaware-core-prod/serviceAc... | 75.00 | - | wanaware-core-prod |
| 15 | SERVICE ACCOUNT THAT EXPOSES MANY VALUED RESOURCES | projects/wa-gcp-test/serviceAccounts/... | 75.00 | - | wa-gcp-test |
| 16 | POD THAT EXPOSES MANY VALUED RESOURCES | metrics-server-v1.33.0-6b8795c6f6-6dhch | 75.00 | - | wanaware-core-prod |
| 17 | FIREWALL THAT EXPOSES MANY VALUED RESOURCES | default-allow-https | 75.00 | - | wanaware-security-map-dev |
| 18 | POD THAT EXPOSES MANY VALUED RESOURCES | gmp-operator-cf8d68dbd-xkl2f | 75.00 | - | wanaware-core-prod |
| 19 | POD THAT EXPOSES MANY VALUED RESOURCES | kube-dns-autoscaler-7fbb9fbfb-2rg6k | 75.00 | - | wanaware-core-prod |
| 20 | FIREWALL THAT EXPOSES MANY VALUED RESOURCES | default-allow-ssh | 75.00 | - | wa-gcp-test |

### HIGH Priority (364 findings)

| # | Category | Resource | Risk Score | CVE | Project |
|--:|----------|----------|----------:|-----|---------|
| 1 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 2 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 3 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 4 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 5 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 6 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 7 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 8 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 9 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 10 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 11 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 12 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 13 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 14 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 15 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 16 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 17 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 18 | VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED |  | 55.00 | - |  |
| 19 | GKE_SECURITY_BULLETIN | internet-mapper | 55.00 | - | wanaware-mapper-dev |
| 20 | GKE_SECURITY_BULLETIN | internet-mapper | 55.00 | - | wanaware-mapper-dev |

### LOW Priority (150 findings)

| # | Category | Resource | Risk Score | CVE | Project |
|--:|----------|----------|----------:|-----|---------|
| 1 | VERTEX_AI_METADATA_STORE_CMEK_DISABLED | projects/358972189709/locations/us-ce... | 20.00 | - | wanaware-stage |
| 2 | VERTEX_AI_MODEL_CMEK_DISABLED | projects/45062729948/locations/us-cen... | 20.00 | - | wanaware-dev |
| 3 | VERTEX_AI_METADATA_STORE_CMEK_DISABLED | projects/62518548529/locations/us-cen... | 20.00 | - | wanaware-security-map-dev |
| 4 | VERTEX_AI_METADATA_STORE_CMEK_DISABLED | projects/45062729948/locations/us-cen... | 20.00 | - | wanaware-dev |
| 5 | VERTEX_AI_MODEL_CMEK_DISABLED | projects/358972189709/locations/us-ce... | 20.00 | - | wanaware-stage |
| 6 | VERTEX_AI_TENSORBOARD_CMEK_DISABLED | Default TensorBoard 2026-01-08 00:28:... | 20.00 | - | wanaware-dev |
| 7 | VERTEX_AI_DATASET_CMEK_DISABLED | projects/62518548529/locations/us-cen... | 20.00 | - | wanaware-security-map-dev |
| 8 | VERTEX_AI_MODEL_CMEK_DISABLED | projects/62518548529/locations/us-cen... | 20.00 | - | wanaware-security-map-dev |
| 9 | VERTEX_AI_TRAINING_PIPELINE_CMEK_DISABLED | projects/62518548529/locations/us-cen... | 20.00 | - | wanaware-security-map-dev |
| 10 | VERTEX_AI_MODEL_CMEK_DISABLED | projects/45062729948/locations/us-cen... | 20.00 | - | wanaware-dev |
| 11 | VERTEX_AI_TENSORBOARD_CMEK_DISABLED | Default TensorBoard 2026-02-17 22:52:... | 20.00 | - | wanaware-stage |
| 12 | VERTEX_AI_DATASET_CMEK_DISABLED | projects/62518548529/locations/us-cen... | 20.00 | - | wanaware-security-map-dev |
| 13 | VERTEX_AI_MODEL_CMEK_DISABLED | projects/45062729948/locations/us-cen... | 20.00 | - | wanaware-dev |
| 14 | GKE_PRIVILEGE_ESCALATION | wanaware-deployment-cluster | 30.00 | - | wanaware-core-dev |
| 15 | GKE_PRIVILEGE_ESCALATION | internet-mapper | 30.00 | - | wanaware-mapper-dev |
| 16 | GKE_RUN_AS_NONROOT | internet-mapper | 20.00 | - | wanaware-mapper-dev |
| 17 | GKE_RUN_AS_NONROOT | wanaware-deployment-cluster | 20.00 | - | wanaware-core-dev |
| 18 | INITIAL ACCESS: GKE NODEPORT SERVICE CREATED | wanaware-deployment-cluster | 24.00 | - | wanaware-core-dev |
| 19 | PERSISTENCE: SERVICE ACCOUNT CREATED IN SENSITIVE NAMESPACE | internet-mapper | 16.00 | - | wanaware-mapper-dev |
| 20 | PERSISTENCE: NEW API METHOD | internet-mapper | 16.00 | - | wanaware-mapper-dev |


---

## Compliance Framework Violations

### CCM

- **Total Violations**: 5
- **Top Violations**:
- CCM LOG-08: 8 finding(s)
- CCM DSP-17: 3 finding(s)
- CCM LOG-05: 2 finding(s)
- CCM IVS-04: 1 finding(s)
- CCM IAM-09: 1 finding(s)

### CIS

- **Total Violations**: 23
- **Top Violations**:
- CIS 3.8: 84 finding(s)
- CIS 3.9: 42 finding(s)
- CIS 5.3: 2 finding(s)
- CIS 1.4: 2 finding(s)
- CIS 5.2: 2 finding(s)

### CIS-CONTROLS

- **Total Violations**: 12
- **Top Violations**:
- CIS-CONTROLS 8.2: 10 finding(s)
- CIS-CONTROLS 8.5: 6 finding(s)
- CIS-CONTROLS 3.3: 3 finding(s)
- CIS-CONTROLS 4.4: 2 finding(s)
- CIS-CONTROLS 8.11: 2 finding(s)

### HIPAA

- **Total Violations**: 5
- **Top Violations**:
- HIPAA 164.312(b): 10 finding(s)
- HIPAA 164.312(a)(1): 3 finding(s)
- HIPAA 164.308(a)(3)(i): 3 finding(s)
- HIPAA 164.308(a)(3)(ii): 3 finding(s)
- HIPAA 164.308(a)(1)(ii): 2 finding(s)

### ISO

- **Total Violations**: 5
- **Top Violations**:
- ISO A.13.1.1: 48 finding(s)
- ISO A.9.2.3: 6 finding(s)
- ISO A.12.4.1: 1 finding(s)
- ISO A.16.1.7: 1 finding(s)
- ISO A.18.1.3: 1 finding(s)

### ISO27001

- **Total Violations**: 9
- **Top Violations**:
- ISO27001 A.8.15: 8 finding(s)
- ISO27001 A.8.20: 8 finding(s)
- ISO27001 A.5.15: 5 finding(s)
- ISO27001 A.5.10: 3 finding(s)
- ISO27001 A.8.3: 3 finding(s)

### NIST

- **Total Violations**: 11
- **Top Violations**:
- NIST SI-4: 46 finding(s)
- NIST PR-PT-1: 10 finding(s)
- NIST DE-AE-3: 8 finding(s)
- NIST AC-6: 6 finding(s)
- NIST PR-AC-4: 4 finding(s)

### NIST800-53

- **Total Violations**: 16
- **Top Violations**:
- NIST800-53 AU-7: 10 finding(s)
- NIST800-53 AU-2: 8 finding(s)
- NIST800-53 AU-12: 8 finding(s)
- NIST800-53 AC-6: 4 finding(s)
- NIST800-53 AC-5: 3 finding(s)

### PCI-DSS

- **Total Violations**: 30
- **Top Violations**:
- PCI-DSS 10.1: 47 finding(s)
- PCI-DSS 10.2: 47 finding(s)
- PCI-DSS 10.2.1.6: 8 finding(s)
- PCI-DSS 6.4.1: 8 finding(s)
- PCI-DSS 10.2.1.2: 8 finding(s)

### SOC2

- **Total Violations**: 26
- **Top Violations**:
- SOC2 CC6.1.7: 4 finding(s)
- SOC2 CC6.1.3: 4 finding(s)
- SOC2 CC5.2.3: 3 finding(s)
- SOC2 CC4.1.3: 2 finding(s)
- SOC2 CC7.3.4: 2 finding(s)


---

## Category Breakdown

| Category | Total | Critical | High | Medium | Low | Avg Score |
|----------|------:|---------:|-----:|-------:|----:|----------:|
| SOFTWARE_VULNERABILITY | 338 | 87 | 251 | 0 | 0 | 76.6 |
| OS_VULNERABILITY | 95 | 48 | 47 | 0 | 0 | 78.8 |
| FLOW_LOGS_DISABLED | 42 | 0 | 0 | 0 | 42 | 22.0 |
| PRIVATE_GOOGLE_ACCESS_DISABLED | 42 | 0 | 0 | 0 | 42 | 16.0 |
| POD THAT EXPOSES MANY VALUED RESOURCES | 25 | 25 | 0 | 0 | 0 | 75.0 |
| VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED | 18 | 0 | 18 | 0 | 0 | 55.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BUCKET OBJECTS. | 13 | 1 | 12 | 0 | 0 | 56.5 |
| FIREWALL THAT EXPOSES MANY VALUED RESOURCES | 12 | 12 | 0 | 0 | 0 | 75.0 |
| PERSISTENCE: NEW API METHOD | 9 | 0 | 0 | 0 | 9 | 16.0 |
| SERVICE ACCOUNT THAT EXPOSES MANY VALUED RESOURCES | 8 | 8 | 0 | 0 | 0 | 75.0 |
| CONTAINER_IMAGE_VULNERABILITY | 6 | 2 | 4 | 0 | 0 | 69.0 |
| PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY. | 6 | 2 | 4 | 0 | 0 | 61.7 |
| INSTANCE THAT EXPOSES MANY VALUED RESOURCES | 5 | 5 | 0 | 0 | 0 | 75.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO ASSUME OTHER SERVICE ACCOUNTS. | 5 | 1 | 4 | 0 | 0 | 59.0 |
| SERVICE ACCOUNT KEY THAT EXPOSES MANY VALUED RESOURCES | 5 | 5 | 0 | 0 | 0 | 75.0 |
| PRIMITIVE_ROLES_USED | 5 | 0 | 0 | 0 | 5 | 24.0 |
| VERTEX_AI_MODEL_CMEK_DISABLED | 5 | 0 | 0 | 0 | 5 | 20.0 |
| PERSISTENCE: IAM ANOMALOUS GRANT | 4 | 0 | 4 | 0 | 0 | 55.0 |
| PERSISTENCE: SERVICE ACCOUNT CREATED IN SENSITIVE NAMESPACE | 4 | 0 | 0 | 0 | 4 | 16.0 |
| FIREWALL_RULE_LOGGING_DISABLED | 4 | 0 | 0 | 0 | 4 | 32.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BIGQUERY DATASETS. | 4 | 0 | 4 | 0 | 0 | 55.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PATCH KUBERNETES NODES. | 4 | 0 | 4 | 0 | 0 | 55.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO ASSUME OTHER SERVICE ACCOUNTS LEADING TO MANAGEMENT OPERATIONS ON VERTEX AI MODEL. | 3 | 0 | 3 | 0 | 0 | 55.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO MODIFY THE METADATA INFORMATION OF AN INSTANCE. | 3 | 1 | 2 | 0 | 0 | 61.7 |
| VERTEX_AI_METADATA_STORE_CMEK_DISABLED | 3 | 0 | 0 | 0 | 3 | 20.0 |
| PERSISTENCE: UNMANAGED ACCOUNT GRANTED SENSITIVE ROLE | 2 | 0 | 2 | 0 | 0 | 55.0 |
| VERTEX_AI_DATASET_CMEK_DISABLED | 2 | 0 | 0 | 0 | 2 | 20.0 |
| GKE_SECURITY_BULLETIN | 2 | 0 | 2 | 0 | 0 | 55.0 |
| VERTEX_AI_TENSORBOARD_CMEK_DISABLED | 2 | 0 | 0 | 0 | 2 | 20.0 |
| GKE_RUN_AS_NONROOT | 2 | 0 | 0 | 0 | 2 | 20.0 |
| INITIAL ACCESS: EXCESSIVE PERMISSION DENIED ACTIONS | 2 | 0 | 0 | 0 | 2 | 24.0 |
| PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY AND THE ABILITY TO ASSUME SERVICE ACCOUNTS. | 2 | 2 | 0 | 0 | 0 | 75.0 |
| GKE_PRIVILEGE_ESCALATION | 2 | 0 | 0 | 0 | 2 | 30.0 |
| BUCKET_LOGGING_DISABLED | 2 | 0 | 0 | 0 | 2 | 24.0 |
| BUCKET_POLICY_ONLY_DISABLED | 2 | 0 | 0 | 0 | 2 | 25.6 |
| PERSISTENCE: NEW USER AGENT | 2 | 0 | 0 | 0 | 2 | 16.0 |
| OVER_PRIVILEGED_SERVICE_ACCOUNT_USER | 1 | 0 | 0 | 0 | 1 | 24.0 |
| CUSTOM_ROLE_NOT_MONITORED | 1 | 0 | 0 | 0 | 1 | 16.0 |
| AUDIT_LOGGING_DISABLED | 1 | 0 | 0 | 0 | 1 | 22.0 |
| FIREWALL_NOT_MONITORED | 1 | 0 | 0 | 0 | 1 | 20.0 |
| NETWORK_NOT_MONITORED | 1 | 0 | 0 | 0 | 1 | 20.0 |
| DEFAULT_NETWORK | 1 | 0 | 0 | 0 | 1 | 30.0 |
| DNS_LOGGING_DISABLED | 1 | 0 | 0 | 0 | 1 | 25.6 |
| ADMIN_SERVICE_ACCOUNT | 1 | 0 | 0 | 0 | 1 | 24.0 |
| OPEN_RDP_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| EGRESS_DENY_RULE_NOT_SET | 1 | 0 | 0 | 0 | 1 | 16.0 |
| ROUTE_NOT_MONITORED | 1 | 0 | 0 | 0 | 1 | 16.0 |
| USER_MANAGED_SERVICE_ACCOUNT_KEY | 1 | 0 | 0 | 0 | 1 | 30.0 |
| OPEN_SSH_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OS_LOGIN_DISABLED | 1 | 0 | 0 | 0 | 1 | 25.6 |
| INITIAL ACCESS: GKE NODEPORT SERVICE CREATED | 1 | 0 | 0 | 0 | 1 | 24.0 |
| PERSISTENCE: NEW GEOGRAPHY | 1 | 0 | 0 | 0 | 1 | 16.0 |
| VERTEX_AI_TRAINING_PIPELINE_CMEK_DISABLED | 1 | 0 | 0 | 0 | 1 | 20.0 |
| OWNER_NOT_MONITORED | 1 | 0 | 0 | 0 | 1 | 16.0 |
| LOG_NOT_EXPORTED | 1 | 0 | 0 | 0 | 1 | 17.6 |
| BUCKET_IAM_NOT_MONITORED | 1 | 0 | 0 | 0 | 1 | 22.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PATCH KUBERNETES WORKLOADS. | 1 | 0 | 1 | 0 | 0 | 55.0 |
| PERSISTENCE: SERVICE ACCOUNT KEY CREATED | 1 | 0 | 0 | 0 | 1 | 16.0 |
| AUDIT_CONFIG_NOT_MONITORED | 1 | 0 | 0 | 0 | 1 | 17.6 |


---

## Project Breakdown

| Project | Total | Critical | High | Medium | Low | Avg Score |
|---------|------:|---------:|-----:|-------:|----:|----------:|
| wanaware-dev | 257 | 84 | 168 | 0 | 5 | 75.2 |
| wanaware-core-prod | 223 | 84 | 138 | 0 | 1 | 76.3 |
| wanaware-ai-dev | 102 | 0 | 6 | 0 | 96 | 21.9 |
| wanaware-security-map-dev | 32 | 13 | 14 | 0 | 5 | 59.0 |
| wanaware-core-dev | 24 | 11 | 8 | 0 | 5 | 60.2 |
| wanaware-mapper-dev | 17 | 0 | 2 | 0 | 15 | 22.6 |
| aerobic-pivot-488020-q5 | 10 | 0 | 0 | 0 | 10 | 19.1 |
| wanaware-stage | 9 | 0 | 1 | 0 | 8 | 26.5 |
| wanaware-core-stage | 7 | 2 | 5 | 0 | 0 | 60.7 |
| wa-gcp-test | 6 | 5 | 1 | 0 | 0 | 71.7 |
| wanaware-prod | 4 | 0 | 3 | 0 | 1 | 47.2 |


---

## Remediation Actions
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/45062729948/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/989564447475/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/866233621206/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/576796365153/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/1095920349335/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/folders/443458554953/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/96530720838/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/150060933135/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/358972189709/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/742666555197/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/503291607878/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/885453410960/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/157381923976/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/124563997292/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/folders/40046211539/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/62518548529/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/940510641271/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED finding on https://orgpolicy.googleapis.com/v2/projects/641194490051/policies/ainotebooks.disableFileDownloads:getEffectivePolicy

**Next Steps**:
- Turn off file downloading for the instance.
- 1. In the Google Cloud console, go to the [**Instances** page](https://console.cloud.google.com/vertex-ai/workbench/instances).
- 2. Click the instance that you want to configure.
- 3. In the **Software and security** tab, add the `notebook-disable-downloads` metadata key and set the value to `TRUE`.
- For more information, see [Update an instance's metadata](https://cloud.google.com/vertex-ai/docs/workbench/instances/manage-metadata#update).
**Estimated Effort**: Low-Medium (30 min - 2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### GKE_SECURITY_BULLETIN — internet-mapper

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: wanaware-mapper-dev
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate GKE_SECURITY_BULLETIN finding on internet-mapper

**Next Steps**:
- Follow the [instructions to upgrade your NodePool](https://cloud.google.com/kubernetes-engine/docs/how-to/upgrading-a-cluster#upgrading-nodes) to version 1.34.3-gke.1318000 or later
**Estimated Effort**: Medium-High (2 hours - 2 days)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### GKE_SECURITY_BULLETIN — internet-mapper

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: wanaware-mapper-dev
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate GKE_SECURITY_BULLETIN finding on internet-mapper

**Next Steps**:
- Follow the [instructions to upgrade your NodePool](https://cloud.google.com/kubernetes-engine/docs/how-to/upgrading-a-cluster#upgrading-nodes) to version 1.34.3-gke.1318000 or later
**Estimated Effort**: Medium-High (2 hours - 2 days)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### POD THAT EXPOSES MANY VALUED RESOURCES — konnectivity-agent-autoscaler-d7655748b-sqbts

- **Priority**: CRITICAL
- **Risk Score**: 75.00
- **Project**: wanaware-core-prod
- **Finding Class**: 9

- **Rationale**: CRITICAL severity (40 pts); category weight 0.8x; GCP severity floor applied (CRITICAL → 75); = 75.00

**Summary**: Remediate POD THAT EXPOSES MANY VALUED RESOURCES finding on konnectivity-agent-autoscaler-d7655748b-sqbts

**Next Steps**:
- Protecting Kubernetes Pods is fundamental. Secure your containerized workloads. Compromised Pods risk cluster integrity, data, and service availability. Mitigate Pod abuse with these best practices:
- __1. Pod and Container Configuration Security (SecurityContext):__
- Use SecurityContext. Restrict Pod and container privileges effectively.
- Run as non-root. Set `runAsUser` and `runAsGroup`. Ensure `runAsNonRoot` is `true`.
- Use a read-only root filesystem. Set `readOnlyRootFilesystem` to `true`.
- Drop unnecessary Linux capabilities. Add back only essential ones. Use `securityContext.capabilities.drop: ["ALL"]`.
- Disallow privilege escalation. Set `allowPrivilegeEscalation` to `false`.
- Define resource requests and limits. Set CPU and memory for all containers. Prevent resource exhaustion.
- Use Pod Security Admission (PSA). Enforce predefined security standards (e.g., baseline, restricted) cluster-wide.
- __2. Workload Identity and Access Control:__
- Use dedicated ServiceAccounts for Pods. Avoid the `default` ServiceAccount.
- Grant ServiceAccounts minimal RBAC permissions. Use Roles and RoleBindings for necessary API access.
- Disable auto-mounting ServiceAccount tokens. Set `automountServiceAccountToken: false` if Pods do not need Kubernetes API access.
- Mount secrets as volumes. Avoid using environment variables for sensitive data. Use Kubernetes Secret objects.
- Use cloud provider workload identity mechanisms. Allow Pods to securely access cloud APIs. Avoid storing long-lived cloud credentials in Pods.
- __3. Network Security for Pods:__
- Implement NetworkPolicies. Control Pod ingress and egress traffic. Start with a default-deny policy.
- Avoid `hostNetwork: true`. It bypasses NetworkPolicies and increases security risks significantly.
- Avoid `hostPort`. It can cause port conflicts and bypass standard service routing. Use Kubernetes Services instead.
- __4. Image and Runtime Security:__
- Use secure and minimal base images. Update base images regularly to include security patches.
- Scan images for vulnerabilities. Integrate image scanning into your CI/CD pipelines before deployment.
- Use immutable images. Do not patch running containers. Rebuild and redeploy new, patched images.
- Consider runtime sandboxing for sensitive workloads. Use gVisor or Kata Containers if available and appropriate.
- By implementing these best practices, you can significantly enhance the security of your Kubernetes Pod security and protect your applications and maintain cluster integrity.
- Please also review the full attack paths which will provide more detail about the actions an attacker is expected to attempt in order to abuse this Pod.
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
**Full Remediation Script** (bash):
```bash
#!/usr/bin/env bash
# ============================================================
# Finding:   01bc40a7b3a178147699e0718cf38708
# Category:  POD THAT EXPOSES MANY VALUED RESOURCES
# Priority:  CRITICAL (risk score 75.00)
# Resource:  konnectivity-agent-autoscaler-d7655748b-sqbts
# Project:   503291607878
# Generated: 2026-02-22
# ============================================================
set -euo pipefail

PROJECT="503291607878"
RESOURCE="konnectivity-agent-autoscaler-d7655748b-sqbts"
CATEGORY="POD THAT EXPOSES MANY VALUED RESOURCES"
DRY_RUN=${DRY_RUN:-false}

echo "[INFO] Inspecting ${CATEGORY} finding on resource: ${RESOURCE}"
echo "[INFO] Project: ${PROJECT}"
echo ""

echo "[INFO] Searching for the resource in Cloud Asset Inventory..."
gcloud asset search-all-resources \
  --project="${PROJECT}" \
  --query="name:${RESOURCE}" \
  --format="table(name,assetType,location,state)" || true

echo ""
echo "[INFO] View this finding in Security Command Center:"
echo "  https://console.cloud.google.com/security/command-center/findings?project=${PROJECT}"
echo ""
echo "[ACTION REQUIRED] Review the finding details and apply remediation steps from the report."

```
---
### POD THAT EXPOSES MANY VALUED RESOURCES — fleet-manager-deployment-84f6db59df-t4fcj

- **Priority**: CRITICAL
- **Risk Score**: 75.00
- **Project**: wanaware-core-prod
- **Finding Class**: 9

- **Rationale**: CRITICAL severity (40 pts); category weight 0.8x; GCP severity floor applied (CRITICAL → 75); = 75.00

**Summary**: Remediate POD THAT EXPOSES MANY VALUED RESOURCES finding on fleet-manager-deployment-84f6db59df-t4fcj

**Next Steps**:
- Protecting Kubernetes Pods is fundamental. Secure your containerized workloads. Compromised Pods risk cluster integrity, data, and service availability. Mitigate Pod abuse with these best practices:
- __1. Pod and Container Configuration Security (SecurityContext):__
- Use SecurityContext. Restrict Pod and container privileges effectively.
- Run as non-root. Set `runAsUser` and `runAsGroup`. Ensure `runAsNonRoot` is `true`.
- Use a read-only root filesystem. Set `readOnlyRootFilesystem` to `true`.
- Drop unnecessary Linux capabilities. Add back only essential ones. Use `securityContext.capabilities.drop: ["ALL"]`.
- Disallow privilege escalation. Set `allowPrivilegeEscalation` to `false`.
- Define resource requests and limits. Set CPU and memory for all containers. Prevent resource exhaustion.
- Use Pod Security Admission (PSA). Enforce predefined security standards (e.g., baseline, restricted) cluster-wide.
- __2. Workload Identity and Access Control:__
- Use dedicated ServiceAccounts for Pods. Avoid the `default` ServiceAccount.
- Grant ServiceAccounts minimal RBAC permissions. Use Roles and RoleBindings for necessary API access.
- Disable auto-mounting ServiceAccount tokens. Set `automountServiceAccountToken: false` if Pods do not need Kubernetes API access.
- Mount secrets as volumes. Avoid using environment variables for sensitive data. Use Kubernetes Secret objects.
- Use cloud provider workload identity mechanisms. Allow Pods to securely access cloud APIs. Avoid storing long-lived cloud credentials in Pods.
- __3. Network Security for Pods:__
- Implement NetworkPolicies. Control Pod ingress and egress traffic. Start with a default-deny policy.
- Avoid `hostNetwork: true`. It bypasses NetworkPolicies and increases security risks significantly.
- Avoid `hostPort`. It can cause port conflicts and bypass standard service routing. Use Kubernetes Services instead.
- __4. Image and Runtime Security:__
- Use secure and minimal base images. Update base images regularly to include security patches.
- Scan images for vulnerabilities. Integrate image scanning into your CI/CD pipelines before deployment.
- Use immutable images. Do not patch running containers. Rebuild and redeploy new, patched images.
- Consider runtime sandboxing for sensitive workloads. Use gVisor or Kata Containers if available and appropriate.
- By implementing these best practices, you can significantly enhance the security of your Kubernetes Pod security and protect your applications and maintain cluster integrity.
- Please also review the full attack paths which will provide more detail about the actions an attacker is expected to attempt in order to abuse this Pod.
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
**Full Remediation Script** (bash):
```bash
#!/usr/bin/env bash
# ============================================================
# Finding:   035e48a3a82240ff09d4f09a21c8ce6a
# Category:  POD THAT EXPOSES MANY VALUED RESOURCES
# Priority:  CRITICAL (risk score 75.00)
# Resource:  fleet-manager-deployment-84f6db59df-t4fcj
# Project:   503291607878
# Generated: 2026-02-22
# ============================================================
set -euo pipefail

PROJECT="503291607878"
RESOURCE="fleet-manager-deployment-84f6db59df-t4fcj"
CATEGORY="POD THAT EXPOSES MANY VALUED RESOURCES"
DRY_RUN=${DRY_RUN:-false}

echo "[INFO] Inspecting ${CATEGORY} finding on resource: ${RESOURCE}"
echo "[INFO] Project: ${PROJECT}"
echo ""

echo "[INFO] Searching for the resource in Cloud Asset Inventory..."
gcloud asset search-all-resources \
  --project="${PROJECT}" \
  --query="name:${RESOURCE}" \
  --format="table(name,assetType,location,state)" || true

echo ""
echo "[INFO] View this finding in Security Command Center:"
echo "  https://console.cloud.google.com/security/command-center/findings?project=${PROJECT}"
echo ""
echo "[ACTION REQUIRED] Review the finding details and apply remediation steps from the report."

```
---
### USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BUCKET OBJECTS. — projects/wanaware-core-stage/serviceAccounts/tf-deploy@wa...

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: wanaware-core-stage
- **Finding Class**: TOXIC_COMBINATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BUCKET OBJECTS. finding on projects/wanaware-core-stage/serviceAccounts/tf-deploy@wanaware-core-stage.iam.gserviceaccount.com

**Next Steps**:
- Service account keys can become [a security risk](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#credential-leakage) if not managed carefully. You should [choose a more secure alternative](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#alternatives) for authentication whenever possible.
- __Short-term mitigation__
- 1. Delete the key: [Delete the user-managed service account key](https://cloud.google.com/iam/docs/keys-create-delete#deleting) if they are not used by any application.
- 2. Rotate the key: Replace the key or, if automatic rotation of user-managed keys for service accounts is not enabled, [enable automatic rotation](https://cloud.google.com/iam/docs/key-rotation).
- 3. Revoke [excess permissions](https://cloud.google.com/policy-intelligence/docs/overview): Consider using more granular roles to remove any permissions that are not strictly required.
- __Long-term remediation__
- 1. Prefer short-lived credentials: Use user credentials, workload identity federation, or service account impersonation [instead of long-lived user-managed keys](https://cloud.google.com/docs/authentication#auth-decision-tree).
- 2. Use organization policy constraints: [Use organization policy constraints](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#orgpolicy-constraints) to prevent creating new service account keys or limit which projects that can create service account keys.
- 3. Enforce separation of duties: Prevent individuals and identities from having both key creation and IAM policy modification permissions and apply the principle of least privilege to all service accounts.
- 4. Monitor and alert: Set up alerts to [detect suspicious activity](https://cloud.google.com/security-command-center/docs/concepts-security-command-center-overview#manage-threats) related to IAM policies and service account keys.
- 5. Help [secure IAM with VPC Service Controls](https://cloud.google.com/iam/docs/secure-iam-vpc-sc): VPC Service Controls let you define a service perimeter for your resources. The service perimeter limits access to Google APIs and services within the defined perimeter.
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY. — my-fleet-instance

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: wanaware-core-dev
- **Finding Class**: TOXIC_COMBINATION

- **Rationale**: HIGH severity (30 pts); category weight 1.2x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY. finding on my-fleet-instance

**Next Steps**:
- Manual SSH key management [poses several risks](https://cloud.google.com/compute/docs/instances/access-overview#risks), including unrestricted sudo access for all users connected to instances with SSH keys stored in metadata. Additionally, you risk granting users, including users outside of your project, unintended access to your compute instances.
- To reduce the attack surface and secure your perimiter, avoid assigning public IP addresses to your compute instances. If you must use instances with public IPs, consider [some of the remediations](https://cloud.google.com/architecture/framework/security/network-security#secure_your_perimeter) listed below.
- __Short-term mitigation__
- 1. Isolate instance: [Revoke the instance's public IP address](https://cloud.google.com/compute/docs/ip-addresses/configure-static-external-ip-address#unassign_ip) or place the instance behind a [firewall with strict rules](https://cloud.google.com/firewall/docs/using-firewalls#creating_firewall_rules) allowing access only from trusted sources limited to the ports required for the application.
- 2. Disable project-wide SSH keys: If not essential, [disable the keys](https://cloud.google.com/compute/docs/connect/restrict-ssh-keys#remove-metadata-key). If they are required, review and restrict their access to specific instances or users.
- __Long-term remediation__
- 1. Use OS Login: OS Login [simplifies SSH access management](https://cloud.google.com/compute/docs/oslogin) by linking your instance user account to your Google identity without having to create and manage individual SSH keys. OS Login is the recommended way to manage many users across multiple instances or projects.
- 2. Configure Cloud VPN or Cloud Interconnect: [Establish a secure connection](https://cloud.google.com/architecture/framework/security/network-security#secure_connections_to_your_on-premises_or_multicloud_environments) between your on-premises network and your VPC network instead of using manually managed SSH keys to public instances.
- 3. Implement Identity-Aware Proxy (IAP): IAP can be used to [control access to instances](https://cloud.google.com/iap/docs/concepts-overview#compute-engine) based on user identity and context. By requiring authentication and authorization through IAP, you can prevent unauthorized public access or lateral movement.Set up an IAP TCP forwarding rule to connect to the instances with internal IP addresses.
- 4. Protecting resources with VPC Service Controls: VPC Service Controls let you [define a service perimeter](https://cloud.google.com/compute/docs/instances/protecting-resources-vpc-service-controls) for your resources. The service perimeter limits access to Google APIs and services within the defined perimeter.
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY. — owasp-test

- **Priority**: CRITICAL
- **Risk Score**: 75.00
- **Project**: wanaware-security-map-dev
- **Finding Class**: TOXIC_COMBINATION

- **Rationale**: CRITICAL severity (40 pts); category weight 1.2x; GCP severity floor applied (CRITICAL → 75); = 75.00

**Summary**: Remediate PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY. finding on owasp-test

**Next Steps**:
- Manual SSH key management [poses several risks](https://cloud.google.com/compute/docs/instances/access-overview#risks), including unrestricted sudo access for all users connected to instances with SSH keys stored in metadata. Additionally, you risk granting users, including users outside of your project, unintended access to your compute instances.
- To reduce the attack surface and secure your perimiter, avoid assigning public IP addresses to your compute instances. If you must use instances with public IPs, consider [some of the remediations](https://cloud.google.com/architecture/framework/security/network-security#secure_your_perimeter) listed below.
- __Short-term mitigation__
- 1. Isolate instance: [Revoke the instance's public IP address](https://cloud.google.com/compute/docs/ip-addresses/configure-static-external-ip-address#unassign_ip) or place the instance behind a [firewall with strict rules](https://cloud.google.com/firewall/docs/using-firewalls#creating_firewall_rules) allowing access only from trusted sources limited to the ports required for the application.
- 2. Disable project-wide SSH keys: If not essential, [disable the keys](https://cloud.google.com/compute/docs/connect/restrict-ssh-keys#remove-metadata-key). If they are required, review and restrict their access to specific instances or users.
- __Long-term remediation__
- 1. Use OS Login: OS Login [simplifies SSH access management](https://cloud.google.com/compute/docs/oslogin) by linking your instance user account to your Google identity without having to create and manage individual SSH keys. OS Login is the recommended way to manage many users across multiple instances or projects.
- 2. Configure Cloud VPN or Cloud Interconnect: [Establish a secure connection](https://cloud.google.com/architecture/framework/security/network-security#secure_connections_to_your_on-premises_or_multicloud_environments) between your on-premises network and your VPC network instead of using manually managed SSH keys to public instances.
- 3. Implement Identity-Aware Proxy (IAP): IAP can be used to [control access to instances](https://cloud.google.com/iap/docs/concepts-overview#compute-engine) based on user identity and context. By requiring authentication and authorization through IAP, you can prevent unauthorized public access or lateral movement.Set up an IAP TCP forwarding rule to connect to the instances with internal IP addresses.
- 4. Protecting resources with VPC Service Controls: VPC Service Controls let you [define a service perimeter](https://cloud.google.com/compute/docs/instances/protecting-resources-vpc-service-controls) for your resources. The service perimeter limits access to Google APIs and services within the defined perimeter.
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
**Full Remediation Script** (bash):
```bash
#!/usr/bin/env bash
# ============================================================
# Finding:   0a08c17a59002748d4962c3a4794b94f
# Category:  PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY.
# Priority:  CRITICAL (risk score 75.00)
# Resource:  owasp-test
# Project:   62518548529
# Generated: 2026-02-22
# ============================================================
set -euo pipefail

PROJECT="62518548529"
RESOURCE="owasp-test"
CATEGORY="PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY."
DRY_RUN=${DRY_RUN:-false}

echo "[INFO] Inspecting ${CATEGORY} finding on resource: ${RESOURCE}"
echo "[INFO] Project: ${PROJECT}"
echo ""

echo "[INFO] Searching for the resource in Cloud Asset Inventory..."
gcloud asset search-all-resources \
  --project="${PROJECT}" \
  --query="name:${RESOURCE}" \
  --format="table(name,assetType,location,state)" || true

echo ""
echo "[INFO] View this finding in Security Command Center:"
echo "  https://console.cloud.google.com/security/command-center/findings?project=${PROJECT}"
echo ""
echo "[ACTION REQUIRED] Review the finding details and apply remediation steps from the report."

```
---
### POD THAT EXPOSES MANY VALUED RESOURCES — chisel-6fb5b877c6-4krgq

- **Priority**: CRITICAL
- **Risk Score**: 75.00
- **Project**: wanaware-core-prod
- **Finding Class**: 9

- **Rationale**: CRITICAL severity (40 pts); category weight 0.8x; GCP severity floor applied (CRITICAL → 75); = 75.00

**Summary**: Remediate POD THAT EXPOSES MANY VALUED RESOURCES finding on chisel-6fb5b877c6-4krgq

**Next Steps**:
- Protecting Kubernetes Pods is fundamental. Secure your containerized workloads. Compromised Pods risk cluster integrity, data, and service availability. Mitigate Pod abuse with these best practices:
- __1. Pod and Container Configuration Security (SecurityContext):__
- Use SecurityContext. Restrict Pod and container privileges effectively.
- Run as non-root. Set `runAsUser` and `runAsGroup`. Ensure `runAsNonRoot` is `true`.
- Use a read-only root filesystem. Set `readOnlyRootFilesystem` to `true`.
- Drop unnecessary Linux capabilities. Add back only essential ones. Use `securityContext.capabilities.drop: ["ALL"]`.
- Disallow privilege escalation. Set `allowPrivilegeEscalation` to `false`.
- Define resource requests and limits. Set CPU and memory for all containers. Prevent resource exhaustion.
- Use Pod Security Admission (PSA). Enforce predefined security standards (e.g., baseline, restricted) cluster-wide.
- __2. Workload Identity and Access Control:__
- Use dedicated ServiceAccounts for Pods. Avoid the `default` ServiceAccount.
- Grant ServiceAccounts minimal RBAC permissions. Use Roles and RoleBindings for necessary API access.
- Disable auto-mounting ServiceAccount tokens. Set `automountServiceAccountToken: false` if Pods do not need Kubernetes API access.
- Mount secrets as volumes. Avoid using environment variables for sensitive data. Use Kubernetes Secret objects.
- Use cloud provider workload identity mechanisms. Allow Pods to securely access cloud APIs. Avoid storing long-lived cloud credentials in Pods.
- __3. Network Security for Pods:__
- Implement NetworkPolicies. Control Pod ingress and egress traffic. Start with a default-deny policy.
- Avoid `hostNetwork: true`. It bypasses NetworkPolicies and increases security risks significantly.
- Avoid `hostPort`. It can cause port conflicts and bypass standard service routing. Use Kubernetes Services instead.
- __4. Image and Runtime Security:__
- Use secure and minimal base images. Update base images regularly to include security patches.
- Scan images for vulnerabilities. Integrate image scanning into your CI/CD pipelines before deployment.
- Use immutable images. Do not patch running containers. Rebuild and redeploy new, patched images.
- Consider runtime sandboxing for sensitive workloads. Use gVisor or Kata Containers if available and appropriate.
- By implementing these best practices, you can significantly enhance the security of your Kubernetes Pod security and protect your applications and maintain cluster integrity.
- Please also review the full attack paths which will provide more detail about the actions an attacker is expected to attempt in order to abuse this Pod.
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
**Full Remediation Script** (bash):
```bash
#!/usr/bin/env bash
# ============================================================
# Finding:   0ec453a4767de76d0e65a305921ce52f
# Category:  POD THAT EXPOSES MANY VALUED RESOURCES
# Priority:  CRITICAL (risk score 75.00)
# Resource:  chisel-6fb5b877c6-4krgq
# Project:   503291607878
# Generated: 2026-02-22
# ============================================================
set -euo pipefail

PROJECT="503291607878"
RESOURCE="chisel-6fb5b877c6-4krgq"
CATEGORY="POD THAT EXPOSES MANY VALUED RESOURCES"
DRY_RUN=${DRY_RUN:-false}

echo "[INFO] Inspecting ${CATEGORY} finding on resource: ${RESOURCE}"
echo "[INFO] Project: ${PROJECT}"
echo ""

echo "[INFO] Searching for the resource in Cloud Asset Inventory..."
gcloud asset search-all-resources \
  --project="${PROJECT}" \
  --query="name:${RESOURCE}" \
  --format="table(name,assetType,location,state)" || true

echo ""
echo "[INFO] View this finding in Security Command Center:"
echo "  https://console.cloud.google.com/security/command-center/findings?project=${PROJECT}"
echo ""
echo "[ACTION REQUIRED] Review the finding details and apply remediation steps from the report."

```
---
### POD THAT EXPOSES MANY VALUED RESOURCES — reseller-wholesale-worker-deployment-84647999cf-ng6p4

- **Priority**: CRITICAL
- **Risk Score**: 75.00
- **Project**: wanaware-core-prod
- **Finding Class**: 9

- **Rationale**: CRITICAL severity (40 pts); category weight 0.8x; GCP severity floor applied (CRITICAL → 75); = 75.00

**Summary**: Remediate POD THAT EXPOSES MANY VALUED RESOURCES finding on reseller-wholesale-worker-deployment-84647999cf-ng6p4

**Next Steps**:
- Protecting Kubernetes Pods is fundamental. Secure your containerized workloads. Compromised Pods risk cluster integrity, data, and service availability. Mitigate Pod abuse with these best practices:
- __1. Pod and Container Configuration Security (SecurityContext):__
- Use SecurityContext. Restrict Pod and container privileges effectively.
- Run as non-root. Set `runAsUser` and `runAsGroup`. Ensure `runAsNonRoot` is `true`.
- Use a read-only root filesystem. Set `readOnlyRootFilesystem` to `true`.
- Drop unnecessary Linux capabilities. Add back only essential ones. Use `securityContext.capabilities.drop: ["ALL"]`.
- Disallow privilege escalation. Set `allowPrivilegeEscalation` to `false`.
- Define resource requests and limits. Set CPU and memory for all containers. Prevent resource exhaustion.
- Use Pod Security Admission (PSA). Enforce predefined security standards (e.g., baseline, restricted) cluster-wide.
- __2. Workload Identity and Access Control:__
- Use dedicated ServiceAccounts for Pods. Avoid the `default` ServiceAccount.
- Grant ServiceAccounts minimal RBAC permissions. Use Roles and RoleBindings for necessary API access.
- Disable auto-mounting ServiceAccount tokens. Set `automountServiceAccountToken: false` if Pods do not need Kubernetes API access.
- Mount secrets as volumes. Avoid using environment variables for sensitive data. Use Kubernetes Secret objects.
- Use cloud provider workload identity mechanisms. Allow Pods to securely access cloud APIs. Avoid storing long-lived cloud credentials in Pods.
- __3. Network Security for Pods:__
- Implement NetworkPolicies. Control Pod ingress and egress traffic. Start with a default-deny policy.
- Avoid `hostNetwork: true`. It bypasses NetworkPolicies and increases security risks significantly.
- Avoid `hostPort`. It can cause port conflicts and bypass standard service routing. Use Kubernetes Services instead.
- __4. Image and Runtime Security:__
- Use secure and minimal base images. Update base images regularly to include security patches.
- Scan images for vulnerabilities. Integrate image scanning into your CI/CD pipelines before deployment.
- Use immutable images. Do not patch running containers. Rebuild and redeploy new, patched images.
- Consider runtime sandboxing for sensitive workloads. Use gVisor or Kata Containers if available and appropriate.
- By implementing these best practices, you can significantly enhance the security of your Kubernetes Pod security and protect your applications and maintain cluster integrity.
- Please also review the full attack paths which will provide more detail about the actions an attacker is expected to attempt in order to abuse this Pod.
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
**Full Remediation Script** (bash):
```bash
#!/usr/bin/env bash
# ============================================================
# Finding:   16b7fdb284218c31adb03eadc382f844
# Category:  POD THAT EXPOSES MANY VALUED RESOURCES
# Priority:  CRITICAL (risk score 75.00)
# Resource:  reseller-wholesale-worker-deployment-84647999cf-ng6p4
# Project:   503291607878
# Generated: 2026-02-22
# ============================================================
set -euo pipefail

PROJECT="503291607878"
RESOURCE="reseller-wholesale-worker-deployment-84647999cf-ng6p4"
CATEGORY="POD THAT EXPOSES MANY VALUED RESOURCES"
DRY_RUN=${DRY_RUN:-false}

echo "[INFO] Inspecting ${CATEGORY} finding on resource: ${RESOURCE}"
echo "[INFO] Project: ${PROJECT}"
echo ""

echo "[INFO] Searching for the resource in Cloud Asset Inventory..."
gcloud asset search-all-resources \
  --project="${PROJECT}" \
  --query="name:${RESOURCE}" \
  --format="table(name,assetType,location,state)" || true

echo ""
echo "[INFO] View this finding in Security Command Center:"
echo "  https://console.cloud.google.com/security/command-center/findings?project=${PROJECT}"
echo ""
echo "[ACTION REQUIRED] Review the finding details and apply remediation steps from the report."

```
---
### USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BUCKET OBJECTS. — projects/wanaware-dev/serviceAccounts/45062729948-compute...

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: wanaware-dev
- **Finding Class**: TOXIC_COMBINATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BUCKET OBJECTS. finding on projects/wanaware-dev/serviceAccounts/45062729948-compute@developer.gserviceaccount.com

**Next Steps**:
- Service account keys can become [a security risk](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#credential-leakage) if not managed carefully. You should [choose a more secure alternative](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#alternatives) for authentication whenever possible.
- __Short-term mitigation__
- 1. Delete the key: [Delete the user-managed service account key](https://cloud.google.com/iam/docs/keys-create-delete#deleting) if they are not used by any application.
- 2. Rotate the key: Replace the key or, if automatic rotation of user-managed keys for service accounts is not enabled, [enable automatic rotation](https://cloud.google.com/iam/docs/key-rotation).
- 3. Revoke [excess permissions](https://cloud.google.com/policy-intelligence/docs/overview): Consider using more granular roles to remove any permissions that are not strictly required.
- __Long-term remediation__
- 1. Prefer short-lived credentials: Use user credentials, workload identity federation, or service account impersonation [instead of long-lived user-managed keys](https://cloud.google.com/docs/authentication#auth-decision-tree).
- 2. Use organization policy constraints: [Use organization policy constraints](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#orgpolicy-constraints) to prevent creating new service account keys or limit which projects that can create service account keys.
- 3. Enforce separation of duties: Prevent individuals and identities from having both key creation and IAM policy modification permissions and apply the principle of least privilege to all service accounts.
- 4. Monitor and alert: Set up alerts to [detect suspicious activity](https://cloud.google.com/security-command-center/docs/concepts-security-command-center-overview#manage-threats) related to IAM policies and service account keys.
- 5. Help [secure IAM with VPC Service Controls](https://cloud.google.com/iam/docs/secure-iam-vpc-sc): VPC Service Controls let you define a service perimeter for your resources. The service perimeter limits access to Google APIs and services within the defined perimeter.
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PATCH KUBERNETES NODES. — projects/wanaware-core-prod/serviceAccounts/tf-deploy@wan...

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: wanaware-core-prod
- **Finding Class**: TOXIC_COMBINATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PATCH KUBERNETES NODES. finding on projects/wanaware-core-prod/serviceAccounts/tf-deploy@wanaware-core-prod.iam.gserviceaccount.com

**Next Steps**:
- Service account keys can become [a security risk](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#credential-leakage) if not managed carefully. You should [choose a more secure alternative](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#alternatives) for authentication whenever possible.
- __Short-term mitigation__
- 1. Delete the key: [Delete the user-managed service account key](https://cloud.google.com/iam/docs/keys-create-delete#deleting) if they are not used by any application.
- 2. Rotate the key: Replace the key or, if automatic rotation of user-managed keys for service accounts is not enabled, [enable automatic rotation](https://cloud.google.com/iam/docs/key-rotation).
- 3. Revoke [excess permissions](https://cloud.google.com/policy-intelligence/docs/overview): Remove the service account's ability to [manage Kubernetes workloads](https://cloud.google.com/kubernetes-engine/docs/concepts/access-control#recommendations).
- __Long-term remediation__
- 1. Prefer short-lived credentials: Use user credentials, workload identity federation, or service account impersonation [instead of long-lived user-managed keys](https://cloud.google.com/docs/authentication#auth-decision-tree).
- 2. Use organization policy constraints: [Use organization policy constraints](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#orgpolicy-constraints) to prevent creating new service account keys or limit which projects that can create service account keys.
- 3. Enforce separation of duties: Prevent individuals and identities from having both key creation and Kubernetes management permissions and apply the principle of least privilege to all service accounts.
- 4. Monitor and alert: Set up alerts to [detect suspicious activity](https://cloud.google.com/security-command-center/docs/concepts-security-command-center-overview#manage-threats) related to IAM policies and service account keys.
- 5. Protecting resources with VPC Service Controls: VPC Service Controls let you [define a service perimeter](https://cloud.google.com/vpc-service-controls/docs/enable) for your resources. The service perimeter limits access to Google APIs and services within the defined perimeter.
- 6. Monitor your cluster configuration: Audit your cluster configurations for deviations from your defined settings. Many of the best practices for cluster security, as well as other common misconfigurations, can be [automatically checked using Security Health Analytics](https://cloud.google.com/security-command-center/docs/list-of-security-health-analytics-findings#container_vulnerability_findings).
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---
### USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO ASSUME OTHER SERVICE ACCOUNTS. — projects/wanaware-security-map-dev/serviceAccounts/tfdepl...

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: wanaware-security-map-dev
- **Finding Class**: TOXIC_COMBINATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

**Summary**: Remediate USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO ASSUME OTHER SERVICE ACCOUNTS. finding on projects/wanaware-security-map-dev/serviceAccounts/tfdeploy@wanaware-security-map-dev.iam.gserviceaccount.com

**Next Steps**:
- Service account keys can become [a security risk](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#credential-leakage) if not managed carefully. You should [choose a more secure alternative](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#alternatives) for authentication whenever possible.
- __Short-term mitigation__
- 1. Delete the key: [Delete the user-managed service account key](https://cloud.google.com/iam/docs/keys-create-delete#deleting) if they are not used by any application.
- 2. Rotate the key: Replace the key or, if automatic rotation of user-managed keys for service accounts is not enabled, [enable automatic rotation](https://cloud.google.com/iam/docs/key-rotation).
- 3. Revoke [excess permissions](https://cloud.google.com/policy-intelligence/docs/overview): Remove the service account's [permission to impersonate](https://cloud.google.com/iam/docs/service-account-impersonation#required-permissions) (assume the identity of) other service accounts.
- __Long-term remediation__
- 1. Prefer short-lived credentials: Use user credentials, workload identity federation, or service account impersonation [instead of long-lived user-managed keys](https://cloud.google.com/docs/authentication#auth-decision-tree).
- 2. Use organization policy constraints: [Use organization policy constraints](https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys#orgpolicy-constraints) to prevent creating new service account keys or limit which projects that can create service account keys.
- 3. Enforce separation of duties: Prevent individuals and identities from having both key creation and service account impersonation permissions and apply the principle of least privilege to all service accounts.
- 4. Monitor and alert: Set up alerts to [detect suspicious activity](https://cloud.google.com/security-command-center/docs/concepts-security-command-center-overview#manage-threats) related to IAM policies and service account keys.
- 5. Help [secure IAM with VPC Service Controls](https://cloud.google.com/iam/docs/secure-iam-vpc-sc): VPC Service Controls let you define a service perimeter for your resources. The service perimeter limits access to Google APIs and services within the defined perimeter.
**Estimated Effort**: Medium (1-2 hours)
**Automation Potential**: Medium
```bash
Refer to GCP documentation for CLI/Terraform automation options
```
---


*Report generated by [gcp-security-analyzer](https://github.com/wanaware/gcp-security-analyzer)*
