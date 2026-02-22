# GCP Security Findings Analysis Report

Generated: 2026-02-22 22:49:37 UTC
Input File: findings.csv

---

## Executive Summary

- **Total Active Findings**: 3479
- **Critical Priority**: 326 (9.4%)
- **High Priority**: 919 (26.4%)
- **Medium Priority**: 49 (1.4%)
- **Low Priority**: 2185 (62.8%)

**Risk Score Statistics:**
- Mean: 39.52
- Median: 25.6
- Range: 16 - 94.8
- Std Dev: 25

**Top Risk Categories:**
1. CONTAINER_IMAGE_VULNERABILITY (492 findings)
2. FLOW_LOGS_DISABLED (484 findings)
3. PRIVATE_GOOGLE_ACCESS_DISABLED (474 findings)
4. SOFTWARE_VULNERABILITY (378 findings)
5. PERSISTENCE: NEW API METHOD (112 findings)
6. INITIAL ACCESS: LOG4J COMPROMISE ATTEMPT (103 findings)
7. OS_VULNERABILITY (91 findings)
8. FIREWALL_RULE_LOGGING_DISABLED (87 findings)
9. USER_MANAGED_SERVICE_ACCOUNT_KEY (83 findings)
10. SERVICE_ACCOUNT_KEY_NOT_ROTATED (72 findings)


---

## Priority Breakdown

| Priority | Count | Percentage | Avg Risk Score | Remediation SLA |
|----------|------:|----------:|---------------:|-----------------|
| CRITICAL | 326 | 9.4% | 80.7 | 24-48 hours |
| HIGH | 919 | 26.4% | 68.4 | 1 week |
| MEDIUM | 49 | 1.4% | 36.0 | 30 days |
| LOW | 2185 | 62.8% | 21.3 | 90 days |

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

### CRITICAL Priority (326 findings)

| # | Category | Resource | Risk Score | CVE | Project |
|--:|----------|----------|----------:|-----|---------|
| 1 | VERTEX_AI_MODEL_NOT_PROTECTED_BY_MODEL_ARMOR |  | 75.00 | - |  |
| 2 | VERTEX_AI_MODEL_NOT_PROTECTED_BY_MODEL_ARMOR |  | 75.00 | - |  |
| 3 | VERTEX_1P_TUNED_MODEL_NOT_PROTECTED_BY_MODEL_ARMOR |  | 75.00 | - |  |
| 4 | VERTEX_1P_TUNED_MODEL_NOT_PROTECTED_BY_MODEL_ARMOR |  | 75.00 | - |  |
| 5 | VERTEX_1P_TUNED_MODEL_NOT_PROTECTED_BY_MODEL_ARMOR |  | 75.00 | - |  |
| 6 | POD THAT EXPOSES MANY VALUED RESOURCES | konnectivity-agent-autoscaler-d765574... | 75.00 | - |  |
| 7 | POD THAT EXPOSES MANY VALUED RESOURCES | fleet-manager-deployment-84f6db59df-t... | 75.00 | - |  |
| 8 | PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY. | owasp-test | 75.00 | - |  |
| 9 | POD THAT EXPOSES MANY VALUED RESOURCES | chisel-6fb5b877c6-4krgq | 75.00 | - |  |
| 10 | POD THAT EXPOSES MANY VALUED RESOURCES | reseller-wholesale-worker-deployment-... | 75.00 | - |  |
| 11 | USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO ASSUME OTHER SERVICE ACCOUNTS. | projects/wanaware-security-map-dev/se... | 75.00 | - |  |
| 12 | POD THAT EXPOSES MANY VALUED RESOURCES | l7-default-backend-78858cccc9-7mzlq | 75.00 | - |  |
| 13 | SERVICE ACCOUNT THAT EXPOSES MANY VALUED RESOURCES | projects/wanaware-core-prod/serviceAc... | 75.00 | - |  |
| 14 | FIREWALL THAT EXPOSES MANY VALUED RESOURCES | wanaware-vpc-allow-https | 75.00 | - |  |
| 15 | FIREWALL THAT EXPOSES MANY VALUED RESOURCES | default-allow-rdp | 75.00 | - |  |
| 16 | POD THAT EXPOSES MANY VALUED RESOURCES | task-scheduler-deployment-8477b65488-... | 75.00 | - |  |
| 17 | POD THAT EXPOSES MANY VALUED RESOURCES | postgres-client-578d77d76c-h62mx | 75.00 | - |  |
| 18 | FIREWALL THAT EXPOSES MANY VALUED RESOURCES | default-allow-rdp | 75.00 | - |  |
| 19 | SERVICE ACCOUNT KEY THAT EXPOSES MANY VALUED RESOURCES | projects/wanaware-core-prod/serviceAc... | 75.00 | - |  |
| 20 | SERVICE ACCOUNT THAT EXPOSES MANY VALUED RESOURCES | projects/wa-gcp-test/serviceAccounts/... | 75.00 | - |  |

### HIGH Priority (919 findings)

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
| 19 | SERVICE_AGENT_GRANTED_BASIC_ROLE | wanaware-game-dev | 55.00 | - |  |
| 20 | SERVICE_AGENT_GRANTED_BASIC_ROLE | wa-gcp-test | 55.00 | - |  |

### MEDIUM Priority (49 findings)

| # | Category | Resource | Risk Score | CVE | Project |
|--:|----------|----------|----------:|-----|---------|
| 1 | SQL_PUBLIC_IP | cosmicwar | 36.00 | - |  |
| 2 | SQL_PUBLIC_IP | wanaware-postgressql | 36.00 | - |  |
| 3 | WEAK_SSL_POLICY | cloud-resource-integrator | 36.00 | - |  |
| 4 | WEAK_SSL_POLICY | k8s2-ts-757m3iab-tunnel-chisel-ingres... | 36.00 | - |  |
| 5 | WEAK_SSL_POLICY | k8s2-ts-0n5j5rr1-fleet-manager-fleet-... | 36.00 | - |  |
| 6 | WEAK_SSL_POLICY | k8s2-ts-0n5j5rr1-function-integr-func... | 36.00 | - |  |
| 7 | WEAK_SSL_POLICY | wanaware-docs-https-proxy | 36.00 | - |  |
| 8 | WEAK_SSL_POLICY | worker-registration-service | 36.00 | - |  |
| 9 | SQL_PUBLIC_IP | wanaware-database | 36.00 | - |  |
| 10 | SQL_PUBLIC_IP | wanaware-postgressql-clone | 36.00 | - |  |
| 11 | SQL_PUBLIC_IP | wanaware-database | 36.00 | - |  |
| 12 | WEAK_SSL_POLICY | k8s2-ts-0jk8byvv-worker-service-worke... | 36.00 | - |  |
| 13 | WEAK_SSL_POLICY | azure-integrator | 36.00 | - |  |
| 14 | WEAK_SSL_POLICY | wanaware-docs-https-proxy | 36.00 | - |  |
| 15 | SQL_PUBLIC_IP | wanaware-database | 36.00 | - |  |
| 16 | WEAK_SSL_POLICY | k8s2-ts-757m3iab-frontend-securityapp... | 36.00 | - |  |
| 17 | WEAK_SSL_POLICY | k8s2-ts-0n5j5rr1-aws-integrator-aws-i... | 36.00 | - |  |
| 18 | WEAK_SSL_POLICY | k8s2-ts-p2f3ru6o-frontend-securityapp... | 36.00 | - |  |
| 19 | WEAK_SSL_POLICY | k8s2-ts-suomt3qi-function-integr-func... | 36.00 | - |  |
| 20 | WEAK_SSL_POLICY | k8s2-ts-0jk8byvv-function-integr-func... | 36.00 | - |  |

### LOW Priority (2185 findings)

| # | Category | Resource | Risk Score | CVE | Project |
|--:|----------|----------|----------:|-----|---------|
| 1 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 2 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 3 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 4 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 5 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 6 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 7 | VERTEX_AI_METADATA_STORE_CMEK_DISABLED | projects/358972189709/locations/us-ce... | 22.40 | - |  |
| 8 | VERTEX_AI_MODEL_CMEK_DISABLED | projects/45062729948/locations/us-cen... | 22.40 | - |  |
| 9 | VERTEX_AI_METADATA_STORE_CMEK_DISABLED | projects/62518548529/locations/us-cen... | 22.40 | - |  |
| 10 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 11 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 12 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 13 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 14 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 15 | VERTEX_AI_METADATA_STORE_CMEK_DISABLED | projects/45062729948/locations/us-cen... | 22.40 | - |  |
| 16 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 17 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 18 | VERTEX_AI_MODEL_CMEK_DISABLED | projects/358972189709/locations/us-ce... | 22.40 | - |  |
| 19 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |
| 20 | VERTEX_AI_ENDPOINT_CMEK_DISABLED |  | 22.40 | - |  |


---

## Category Breakdown

| Category | Total | Critical | High | Medium | Low | Avg Score |
|----------|------:|---------:|-----:|-------:|----:|----------:|
| CONTAINER_IMAGE_VULNERABILITY | 492 | 141 | 351 | 0 | 0 | 74.0 |
| FLOW_LOGS_DISABLED | 484 | 0 | 0 | 0 | 484 | 22.0 |
| PRIVATE_GOOGLE_ACCESS_DISABLED | 474 | 0 | 0 | 0 | 474 | 16.0 |
| SOFTWARE_VULNERABILITY | 378 | 81 | 297 | 0 | 0 | 75.8 |
| PERSISTENCE: NEW API METHOD | 112 | 0 | 0 | 0 | 112 | 16.0 |
| INITIAL ACCESS: LOG4J COMPROMISE ATTEMPT | 103 | 0 | 0 | 0 | 103 | 17.6 |
| OS_VULNERABILITY | 91 | 27 | 64 | 0 | 0 | 77.1 |
| FIREWALL_RULE_LOGGING_DISABLED | 87 | 0 | 0 | 0 | 87 | 32.0 |
| USER_MANAGED_SERVICE_ACCOUNT_KEY | 83 | 0 | 0 | 0 | 83 | 30.0 |
| SERVICE_ACCOUNT_KEY_NOT_ROTATED | 72 | 0 | 0 | 0 | 72 | 30.0 |
| BUCKET_LOGGING_DISABLED | 65 | 0 | 0 | 0 | 65 | 24.0 |
| INITIAL ACCESS: GKE NODEPORT SERVICE CREATED | 41 | 0 | 0 | 0 | 41 | 24.0 |
| VERTEX_AI_ENDPOINT_CMEK_DISABLED | 41 | 0 | 0 | 0 | 41 | 22.4 |
| WEAK_SSL_POLICY | 40 | 0 | 0 | 40 | 0 | 36.0 |
| POD THAT EXPOSES MANY VALUED RESOURCES | 33 | 33 | 0 | 0 | 0 | 75.0 |
| HTTP_LOAD_BALANCER | 31 | 0 | 0 | 0 | 31 | 26.4 |
| PERSISTENCE: SERVICE ACCOUNT KEY CREATED | 26 | 0 | 0 | 0 | 26 | 16.0 |
| COMPUTE_SECURE_BOOT_DISABLED | 21 | 0 | 0 | 0 | 21 | 20.0 |
| COMPUTE_PROJECT_WIDE_SSH_KEYS_ALLOWED | 21 | 0 | 0 | 0 | 21 | 24.0 |
| PUBLIC_IP_ADDRESS | 21 | 0 | 21 | 0 | 0 | 55.0 |
| INSTANCE_OS_LOGIN_DISABLED | 21 | 0 | 0 | 0 | 21 | 25.6 |
| AUDIT_LOGGING_DISABLED | 20 | 0 | 0 | 0 | 20 | 22.0 |
| BUCKET_POLICY_ONLY_DISABLED | 20 | 0 | 0 | 0 | 20 | 25.6 |
| DNS_LOGGING_DISABLED | 19 | 0 | 0 | 0 | 19 | 25.6 |
| VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED | 18 | 0 | 18 | 0 | 0 | 55.0 |
| PRIMITIVE_ROLES_USED | 18 | 0 | 0 | 0 | 18 | 24.0 |
| LOG_NOT_EXPORTED | 17 | 0 | 0 | 0 | 17 | 17.6 |
| OPEN_SSH_PORT | 17 | 0 | 17 | 0 | 0 | 55.0 |
| CUSTOM_ROLE_NOT_MONITORED | 16 | 0 | 0 | 0 | 16 | 16.0 |
| NETWORK_NOT_MONITORED | 16 | 0 | 0 | 0 | 16 | 20.0 |
| PERSISTENCE: IAM ANOMALOUS GRANT | 16 | 0 | 16 | 0 | 0 | 55.0 |
| OWNER_NOT_MONITORED | 16 | 0 | 0 | 0 | 16 | 16.0 |
| ROUTE_NOT_MONITORED | 16 | 0 | 0 | 0 | 16 | 16.0 |
| FIREWALL_NOT_MONITORED | 16 | 0 | 0 | 0 | 16 | 20.0 |
| BUCKET_IAM_NOT_MONITORED | 16 | 0 | 0 | 0 | 16 | 22.0 |
| AUDIT_CONFIG_NOT_MONITORED | 16 | 0 | 0 | 0 | 16 | 17.6 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BUCKET OBJECTS. | 14 | 0 | 14 | 0 | 0 | 55.0 |
| UNUSED_IAM_ROLE | 13 | 0 | 11 | 0 | 2 | 49.1 |
| OPEN_FIREWALL | 13 | 0 | 13 | 0 | 0 | 55.0 |
| SQL_LOG_STATEMENT | 12 | 0 | 0 | 0 | 12 | 17.6 |
| SQL_LOG_CHECKPOINTS_DISABLED | 12 | 0 | 0 | 0 | 12 | 25.6 |
| FIREWALL THAT EXPOSES MANY VALUED RESOURCES | 12 | 12 | 0 | 0 | 0 | 75.0 |
| SQL_LOG_TEMP_FILES | 12 | 0 | 0 | 0 | 12 | 25.6 |
| SQL_LOG_CONNECTIONS_DISABLED | 12 | 0 | 0 | 0 | 12 | 25.6 |
| SQL_LOG_LOCK_WAITS_DISABLED | 12 | 0 | 0 | 0 | 12 | 25.6 |
| SQL_LOG_DURATION_DISABLED | 12 | 0 | 0 | 0 | 12 | 25.6 |
| OPEN_RDP_PORT | 12 | 0 | 12 | 0 | 0 | 55.0 |
| SQL_LOG_DISCONNECTIONS_DISABLED | 12 | 0 | 0 | 0 | 12 | 25.6 |
| EGRESS_DENY_RULE_NOT_SET | 11 | 0 | 0 | 0 | 11 | 16.0 |
| SSL_NOT_ENFORCED | 11 | 0 | 11 | 0 | 0 | 55.0 |
| GKE_SECURITY_BULLETIN | 11 | 0 | 11 | 0 | 0 | 55.0 |
| DEFAULT_NETWORK | 11 | 0 | 0 | 0 | 11 | 30.0 |
| OS_LOGIN_DISABLED | 11 | 0 | 0 | 0 | 11 | 25.6 |
| SERVICE ACCOUNT THAT EXPOSES MANY VALUED RESOURCES | 10 | 10 | 0 | 0 | 0 | 75.0 |
| PERSISTENCE: GCE ADMIN ADDED SSH KEY | 10 | 0 | 0 | 0 | 10 | 16.0 |
| OPEN_HTTP_PORT | 9 | 0 | 9 | 0 | 0 | 55.0 |
| SQL_PUBLIC_IP | 9 | 0 | 0 | 9 | 0 | 36.0 |
| API_KEY_EXISTS | 8 | 0 | 0 | 0 | 8 | 24.0 |
| BINARY_AUTHORIZATION_DISABLED | 8 | 0 | 0 | 0 | 8 | 20.0 |
| CLUSTER_SECRETS_ENCRYPTION_DISABLED | 8 | 0 | 0 | 0 | 8 | 24.0 |
| SQL_INSTANCE_NOT_MONITORED | 8 | 0 | 0 | 0 | 8 | 16.0 |
| GKE_RUN_AS_NONROOT | 8 | 0 | 0 | 0 | 8 | 20.0 |
| GKE_PRIVILEGE_ESCALATION | 8 | 0 | 0 | 0 | 8 | 30.0 |
| PRIVATE_CLUSTER_DISABLED | 8 | 0 | 0 | 0 | 8 | 24.0 |
| IAM_ROLE_HAS_EXCESSIVE_PERMISSIONS | 7 | 0 | 4 | 0 | 3 | 38.7 |
| API_KEY_APPS_UNRESTRICTED | 7 | 0 | 0 | 0 | 7 | 24.0 |
| OVER_PRIVILEGED_ACCOUNT | 7 | 0 | 0 | 0 | 7 | 24.0 |
| PERSISTENCE: SERVICE ACCOUNT CREATED IN SENSITIVE NAMESPACE | 7 | 0 | 0 | 0 | 7 | 16.0 |
| API_KEY_NOT_ROTATED | 7 | 0 | 0 | 0 | 7 | 24.0 |
| LOAD_BALANCER_LOGGING_DISABLED | 6 | 0 | 0 | 0 | 6 | 28.0 |
| ADMIN_SERVICE_ACCOUNT | 6 | 0 | 0 | 0 | 6 | 24.0 |
| PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY. | 6 | 2 | 4 | 0 | 0 | 61.7 |
| PUBLIC_BUCKET_ACL | 5 | 0 | 5 | 0 | 0 | 55.0 |
| OVER_PRIVILEGED_SERVICE_ACCOUNT_USER | 5 | 0 | 0 | 0 | 5 | 24.0 |
| SERVICE ACCOUNT KEY THAT EXPOSES MANY VALUED RESOURCES | 5 | 5 | 0 | 0 | 0 | 75.0 |
| VERTEX_AI_MODEL_CMEK_DISABLED | 5 | 0 | 0 | 0 | 5 | 22.4 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO ASSUME OTHER SERVICE ACCOUNTS. | 5 | 1 | 4 | 0 | 0 | 59.0 |
| INSTANCE THAT EXPOSES MANY VALUED RESOURCES | 5 | 5 | 0 | 0 | 0 | 75.0 |
| MASTER_AUTHORIZED_NETWORKS_DISABLED | 5 | 0 | 0 | 0 | 5 | 30.0 |
| DEFAULT_SERVICE_ACCOUNT_USED | 5 | 0 | 0 | 0 | 5 | 24.0 |
| SERVICE_AGENT_GRANTED_BASIC_ROLE | 4 | 0 | 4 | 0 | 0 | 55.0 |
| WORKLOAD_IDENTITY_DISABLED | 4 | 0 | 0 | 0 | 4 | 24.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PATCH KUBERNETES NODES. | 4 | 0 | 4 | 0 | 0 | 55.0 |
| NETWORK_POLICY_DISABLED | 4 | 0 | 0 | 0 | 4 | 30.0 |
| OVER_PRIVILEGED_SCOPES | 4 | 0 | 0 | 0 | 4 | 24.0 |
| PERSISTENCE: NEW USER AGENT | 4 | 0 | 0 | 0 | 4 | 16.0 |
| INTRANODE_VISIBILITY_DISABLED | 4 | 0 | 0 | 0 | 4 | 24.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO MODIFY THE METADATA INFORMATION OF AN INSTANCE. | 4 | 1 | 3 | 0 | 0 | 60.0 |
| VERTEX_1P_TUNED_MODEL_NOT_PROTECTED_BY_MODEL_ARMOR | 3 | 3 | 0 | 0 | 0 | 75.0 |
| VERTEX_AI_METADATA_STORE_CMEK_DISABLED | 3 | 0 | 0 | 0 | 3 | 22.4 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO ASSUME OTHER SERVICE ACCOUNTS LEADING TO MANAGEMENT OPERATIONS ON VERTEX AI MODEL. | 3 | 0 | 3 | 0 | 0 | 55.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BIGQUERY DATASETS. | 3 | 0 | 3 | 0 | 0 | 55.0 |
| API_KEY_APIS_UNRESTRICTED | 3 | 0 | 0 | 0 | 3 | 24.0 |
| VERTEX_AI_TENSORBOARD_CMEK_DISABLED | 2 | 0 | 0 | 0 | 2 | 22.4 |
| PUBLICLY ACCESSIBLE INSTANCE WITH PROJECT-WIDE SSH KEY AND THE ABILITY TO ASSUME SERVICE ACCOUNTS. | 2 | 2 | 0 | 0 | 0 | 75.0 |
| VERTEX_AI_MODEL_NOT_PROTECTED_BY_MODEL_ARMOR | 2 | 2 | 0 | 0 | 0 | 75.0 |
| PERSISTENCE: GCE ADMIN ADDED STARTUP SCRIPT | 2 | 0 | 0 | 0 | 2 | 16.0 |
| AUTO_BACKUP_DISABLED | 2 | 0 | 0 | 0 | 2 | 24.0 |
| VERTEX_AI_DATASET_CMEK_DISABLED | 2 | 0 | 0 | 0 | 2 | 22.4 |
| INITIAL ACCESS: EXCESSIVE PERMISSION DENIED ACTIONS | 2 | 0 | 0 | 0 | 2 | 24.0 |
| PERSISTENCE: UNMANAGED ACCOUNT GRANTED SENSITIVE ROLE | 2 | 0 | 2 | 0 | 0 | 55.0 |
| PERSISTENCE: NEW GEOGRAPHY | 2 | 0 | 0 | 0 | 2 | 16.0 |
| USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PATCH KUBERNETES WORKLOADS. | 1 | 1 | 0 | 0 | 0 | 75.0 |
| OPEN_LDAP_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| NON_ORG_IAM_MEMBER | 1 | 0 | 1 | 0 | 0 | 55.0 |
| MALWARE: CRYPTOMINING BAD IP | 1 | 0 | 0 | 0 | 1 | 16.0 |
| OPEN_MONGODB_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_MYSQL_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_TELNET_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| FULL_API_ACCESS | 1 | 0 | 0 | 0 | 1 | 24.0 |
| MALWARE: BAD IP | 1 | 0 | 0 | 0 | 1 | 16.0 |
| INCREASING DENY RATIO | 1 | 0 | 0 | 0 | 1 | 24.0 |
| OPEN_NETBIOS_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_DNS_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_CASSANDRA_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_POP3_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| ADDED BINARY EXECUTED | 1 | 0 | 0 | 0 | 1 | 16.0 |
| OPEN_CISCOSECURE_WEBSM_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| ACCESS_TRANSPARENCY_DISABLED | 1 | 0 | 0 | 0 | 1 | 24.0 |
| DNSSEC_DISABLED | 1 | 0 | 0 | 0 | 1 | 24.0 |
| KMS_PROJECT_HAS_OWNER | 1 | 0 | 0 | 0 | 1 | 24.0 |
| ESSENTIAL_CONTACTS_NOT_CONFIGURED | 1 | 0 | 0 | 0 | 1 | 16.0 |
| VERTEX_AI_TRAINING_PIPELINE_CMEK_DISABLED | 1 | 0 | 0 | 0 | 1 | 22.4 |
| OPEN_ORACLEDB_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_DIRECTORY_SERVICES_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| NODEPOOL_SECURE_BOOT_DISABLED | 1 | 0 | 0 | 0 | 1 | 24.0 |
| OPEN_SMTP_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_POSTGRESQL_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_ELASTICSEARCH_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_FTP_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_REDIS_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |
| OPEN_MEMCACHED_PORT | 1 | 0 | 1 | 0 | 0 | 55.0 |


---

## Project Breakdown

| Project | Total | Critical | High | Medium | Low | Avg Score |
|---------|------:|---------:|-----:|-------:|----:|----------:|


---

## Remediation Actions
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### VERTEX_AI_JUPYTERLAB_FILE_DOWNLOADING_ENABLED — 

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: MISCONFIGURATION

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### SERVICE_AGENT_GRANTED_BASIC_ROLE — wanaware-game-dev

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### SERVICE_AGENT_GRANTED_BASIC_ROLE — wa-gcp-test

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### UNUSED_IAM_ROLE — wanaware-core-prod

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 1.0x; GCP severity floor applied (HIGH → 55); = 55.00

---
### UNUSED_IAM_ROLE — wanaware-core-prod

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 1.0x; GCP severity floor applied (HIGH → 55); = 55.00

---
### IAM_ROLE_HAS_EXCESSIVE_PERMISSIONS — wanaware-core-stage

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 1.0x; GCP severity floor applied (HIGH → 55); = 55.00

---
### UNUSED_IAM_ROLE — wanaware-core-stage

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 1.0x; GCP severity floor applied (HIGH → 55); = 55.00

---
### UNUSED_IAM_ROLE — wa-gcp-test

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 1.0x; GCP severity floor applied (HIGH → 55); = 55.00

---
### SERVICE_AGENT_GRANTED_BASIC_ROLE — wanaware-prod

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---
### UNUSED_IAM_ROLE — wanaware-game-dev

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 1.0x; GCP severity floor applied (HIGH → 55); = 55.00

---
### UNUSED_IAM_ROLE — wanaware.com

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 1.0x; GCP severity floor applied (HIGH → 55); = 55.00

---
### IAM_ROLE_HAS_EXCESSIVE_PERMISSIONS — wanaware-core-prod

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 1.0x; GCP severity floor applied (HIGH → 55); = 55.00

---
### SERVICE_AGENT_GRANTED_BASIC_ROLE — wanaware-stage

- **Priority**: HIGH
- **Risk Score**: 55.00
- **Project**: 
- **Finding Class**: VULNERABILITY

- **Rationale**: HIGH severity (30 pts); category weight 0.8x; GCP severity floor applied (HIGH → 55); = 55.00

---


*Report generated by [gcp-security-analyzer](https://github.com/wanaware/gcp-security-analyzer)*
