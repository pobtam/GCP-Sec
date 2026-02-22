#!/usr/bin/env bash
# ============================================================
# Finding:   528e727b82b344eed4624870e94fc80d
# Category:  SOFTWARE_VULNERABILITY
# Priority:  CRITICAL (risk score 79.68)
# Resource:  cluster-1
# Project:   45062729948
# CVE:           CVE-2023-30533 (CVSS 7.8)
# Generated: 2026-02-22
# ============================================================
set -euo pipefail

PROJECT="45062729948"
RESOURCE="cluster-1"
CATEGORY="SOFTWARE_VULNERABILITY"
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
