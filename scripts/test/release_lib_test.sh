#!/usr/bin/env bash
# Tests for scripts/release-lib.sh. Run: bash scripts/test/release_lib_test.sh
set -uo pipefail
cd "$(git rev-parse --show-toplevel)" || exit 1
source scripts/release-lib.sh

fail=0
check() { # check <description> <expected> <actual>
  if [ "$2" = "$3" ]; then
    echo "ok: $1"
  else
    echo "FAIL: $1"
    echo "  expected: [$2]"
    echo "  actual:   [$3]"
    fail=1
  fi
}

# --- compute_target_version ---
check "target: patch"    "1.7.1" "$(compute_target_version v1.7.0 patch)"
check "target: minor"    "1.8.0" "$(compute_target_version v1.7.0 minor)"
check "target: major"    "2.0.0" "$(compute_target_version v1.7.0 major)"
check "target: explicit" "3.2.1" "$(compute_target_version v1.7.0 v3.2.1)"
check "target: first"    "0.1.0" "$(compute_target_version v0.0.0 minor)"

# --- is_release_tag ---
tag_ok() { if is_release_tag "$1"; then echo 1; else echo 0; fi; }
check "tag: stable"        "1" "$(tag_ok v1.8.0)"
check "tag: multidigit"    "1" "$(tag_ok v10.20.30)"
check "tag: beta invalid"  "0" "$(tag_ok v1.8.0-beta.1)"
check "tag: no v invalid"  "0" "$(tag_ok 1.8.0)"
check "tag: two-part"      "0" "$(tag_ok v1.8)"

# --- confirm_release (reads stdin) ---
if echo y    | confirm_release v1.7.0 v1.8.0 >/dev/null 2>&1; then got=1; else got=0; fi
check "confirm: y accepts"     "1" "$got"
if echo n    | confirm_release v1.7.0 v1.8.0 >/dev/null 2>&1; then got=1; else got=0; fi
check "confirm: n rejects"     "0" "$got"
if printf '' | confirm_release v1.7.0 v1.8.0 >/dev/null 2>&1; then got=1; else got=0; fi
check "confirm: empty rejects" "0" "$got"

exit $fail
