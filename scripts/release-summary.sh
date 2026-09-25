#!/usr/bin/env bash
# Render the release job summary: release notes + a per-target publish status
# table, appended to $GITHUB_STEP_SUMMARY. Never fails.
#
# Reads step outcomes from GORELEASER_OUTCOME / HOMEBREW_OUTCOME and the release
# title/notes from RELEASE_NAME / RELEASE_HEADER (set by changelog-notes.sh).
set -uo pipefail

# GitHub step outcome (success|failure|skipped|"") -> table status.
outcome_to_status() {
  case "$1" in
  success) echo "published" ;;
  failure) echo "failed" ;;
  skipped) echo "skipped" ;;
  *) echo "not run" ;;
  esac
}

{
  printf '## %s\n\n' "${RELEASE_NAME:-Release}"
  if [ -n "${RELEASE_HEADER:-}" ]; then
    printf '%s\n\n' "$RELEASE_HEADER"
  fi
  printf '| Target                    | Status    |\n'
  printf '| ------------------------- | --------- |\n'
  printf '| GitHub release + packages | %s |\n' "$(outcome_to_status "${GORELEASER_OUTCOME:-}")"
  printf '| Homebrew                  | %s |\n' "$(outcome_to_status "${HOMEBREW_OUTCOME:-}")"
} >>"${GITHUB_STEP_SUMMARY:-/dev/stdout}"
