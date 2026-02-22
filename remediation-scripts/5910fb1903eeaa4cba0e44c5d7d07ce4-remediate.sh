#!/usr/bin/env bash
# ============================================================
# Finding:   5910fb1903eeaa4cba0e44c5d7d07ce4
# Category:  USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BUCKET OBJECTS.
# Priority:  CRITICAL (risk score 75.00)
# Resource:  projects/wanaware-security-map-dev/serviceAccounts/62518548529-compute@developer.gserviceaccount.com
# Project:   62518548529
# Generated: 2026-02-22
# ============================================================
set -euo pipefail

PROJECT="62518548529"
RESOURCE="projects/wanaware-security-map-dev/serviceAccounts/62518548529-compute@developer.gserviceaccount.com"
CATEGORY="USER-MANAGED KEYS TO SERVICE ACCOUNT WITH PERMISSIONS TO PERFORM DATA OPERATIONS ON BUCKET OBJECTS."
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
