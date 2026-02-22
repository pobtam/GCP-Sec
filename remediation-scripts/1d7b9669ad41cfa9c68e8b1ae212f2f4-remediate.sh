#!/usr/bin/env bash
# ============================================================
# Finding:   1d7b9669ad41cfa9c68e8b1ae212f2f4
# Category:  SOFTWARE_VULNERABILITY
# Priority:  CRITICAL (risk score 78.60)
# Resource:  cluster-1
# Project:   45062729948
# CVE:           CVE-2025-65945 (CVSS 7.5)
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
