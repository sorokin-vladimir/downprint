#!/usr/bin/env bash
# Pure helpers for scripts/release.sh: version math and the confirmation gate.
# Sourced, not executed. No git side effects.

# compute_target_version <latest> <bump>
# bump is major|minor|patch or an explicit X.Y.Z / vX.Y.Z. Echoes X.Y.Z.
compute_target_version() {
  local latest=$1 bump=$2 major minor patch
  IFS=. read -r major minor patch <<<"${latest#v}"
  case "$bump" in
    major) echo "$((major + 1)).0.0" ;;
    minor) echo "${major}.$((minor + 1)).0" ;;
    patch) echo "${major}.${minor}.$((patch + 1))" ;;
    *)     echo "${bump#v}" ;;
  esac
}

# is_release_tag <tag>
# Single source of truth for tag validation: returns 0 for vX.Y.Z, else 1.
is_release_tag() {
  [[ $1 =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]
}

# confirm_release <current_tag> <new_tag>
# Prints a summary to stderr, reads one line from stdin. Returns 0 only on y/Y.
confirm_release() {
  local current=$1 newtag=$2 reply
  {
    echo "current version: ${current}"
    echo "new tag:         ${newtag}"
    printf 'proceed? [y/N] '
  } >&2
  read -r reply
  [[ $reply == y || $reply == Y ]]
}
