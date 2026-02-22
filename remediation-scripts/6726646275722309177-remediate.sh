#!/usr/bin/env bash
# ============================================================
# Finding:   6726646275722309177
# Category:  VERTEX_1P_TUNED_MODEL_NOT_PROTECTED_BY_MODEL_ARMOR
# Priority:  CRITICAL (risk score 75.00)
# Resource:  //aiplatform.googleapis.com/projects/45062729948/locations/us-central1/models/new-excel-validator-model-enhanced
# Project:   
# Generated: 2026-02-22
# ============================================================
set -euo pipefail

PROJECT=""
RESOURCE="//aiplatform.googleapis.com/projects/45062729948/locations/us-central1/models/new-excel-validator-model-enhanced"
CATEGORY="VERTEX_1P_TUNED_MODEL_NOT_PROTECTED_BY_MODEL_ARMOR"
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
