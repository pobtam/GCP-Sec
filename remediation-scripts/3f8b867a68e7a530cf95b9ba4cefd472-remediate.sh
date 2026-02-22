#!/usr/bin/env bash
# ============================================================
# Finding:   3f8b867a68e7a530cf95b9ba4cefd472
# Category:  SERVICE ACCOUNT KEY THAT EXPOSES MANY VALUED RESOURCES
# Priority:  CRITICAL (risk score 75.00)
# Resource:  projects/wanaware-security-map-dev/serviceAccounts/62518548529-compute@developer.gserviceaccount.com/keys/fbb56dc4621b79bb0969474221fa73777f4f4282
# Project:   62518548529
# Generated: 2026-02-22
# ============================================================
set -euo pipefail

PROJECT="62518548529"
RESOURCE="projects/wanaware-security-map-dev/serviceAccounts/62518548529-compute@developer.gserviceaccount.com/keys/fbb56dc4621b79bb0969474221fa73777f4f4282"
CATEGORY="SERVICE ACCOUNT KEY THAT EXPOSES MANY VALUED RESOURCES"
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
