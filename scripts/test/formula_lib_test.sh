#!/usr/bin/env bash
# Tests for scripts/formula-lib.sh. Run: bash scripts/test/formula_lib_test.sh
set -uo pipefail
cd "$(git rev-parse --show-toplevel)" || exit 1
source scripts/formula-lib.sh

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

f=$(render_formula sorokin-vladimir 1.2.0 https://example/base AAA BBB CCC DDD)
check "class name"      "1" "$(grep -c '^class Downprint < Formula' <<<"$f")"
check "version"         "1" "$(grep -c 'version "1.2.0"' <<<"$f")"
check "license"         "1" "$(grep -c 'license "MIT"' <<<"$f")"
check "darwin_amd url"  "1" "$(grep -c 'url "https://example/base/downprint_darwin_amd64.tar.gz"' <<<"$f")"
check "darwin_amd sha"  "1" "$(grep -c 'sha256 "AAA"' <<<"$f")"
check "darwin_arm sha"  "1" "$(grep -c 'sha256 "BBB"' <<<"$f")"
check "linux_amd sha"   "1" "$(grep -c 'sha256 "CCC"' <<<"$f")"
check "linux_arm sha"   "1" "$(grep -c 'sha256 "DDD"' <<<"$f")"
check "install"         "1" "$(grep -c 'bin.install "downprint"$' <<<"$f")"
check "chrome caveat"   "1" "$(grep -c 'brew install --cask google-chrome' <<<"$f")"
check "version test"    "1" "$(grep -c 'shell_output("#{bin}/downprint -version")' <<<"$f")"

exit $fail
