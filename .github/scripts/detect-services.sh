#!/bin/bash -eu
# Detects modified services and builds a matrix for GitHub Actions
# Usage: detect-services.sh <has_changes> <files> <marker_file>
#   has_changes: "true" if there are changes in services/, "false" otherwise
#   files: Space-separated list of changed files
#   marker_file: File that identifies the language (e.g., go.mod, poetry.lock, package.json)

readonly HAS_CHANGES="${1:-false}"
readonly FILES="${2:-}"
readonly MARKER_FILE="${3:-go.mod}"

# Skip if no services changes detected
if [[ "$HAS_CHANGES" != "true" ]]; then
  echo 'services_json={"include":[]}' >> "$GITHUB_OUTPUT"
  echo 'has_changes=false' >> "$GITHUB_OUTPUT"
  exit 0
fi

# Extract unique type/service pairs from services/<type>/<service>/ paths
# Filters out hidden directories (starting with .)
# Example: services/backend/reclamis/file.go -> backend/reclamis
type_services=$(echo "$FILES" \
  | grep -oP '^services/\K[^/]+/[^/]+(?=/)' \
  | grep -v '/\.' \
  | sort -u)

# Build JSON matrix for parallel jobs (only services with marker file)
# Format: {"include":[{"type":"backend","service":"reclamis"}]}
json='{"include":['
first=1
while IFS= read -r type_svc; do
  [[ -z "$type_svc" ]] && continue

  # Split type/service (e.g., "backend/reclamis" -> type="backend", svc="reclamis")
  type="${type_svc%%/*}"
  svc="${type_svc##*/}"

  # Only include services that have the marker file
  if [[ -f "services/$type/$svc/$MARKER_FILE" ]]; then
    if [[ $first -eq 1 ]]; then
      first=0
    else
      json+=','
    fi
    json+='{"type":"'"$type"'","service":"'"$svc"'"}'
  fi
done <<< "$type_services"
json+=']}'

echo "Matrix: $json"
echo "services_json=$json" >> "$GITHUB_OUTPUT"

# Check if we have any services
if [[ "$json" == '{"include":[]}' ]]; then
  echo 'has_changes=false' >> "$GITHUB_OUTPUT"
else
  echo 'has_changes=true' >> "$GITHUB_OUTPUT"
fi
