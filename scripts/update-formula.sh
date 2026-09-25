#!/usr/bin/env bash
#
# Render the Homebrew formula for a release tag and push it to the tap repo
# (sorokin-vladimir/homebrew-tap). Replaces GoReleaser's deprecated `brews`
# publisher.
#
# Tarball checksums are read from dist/checksums.txt (produced by
# `goreleaser release`), so this MUST run after goreleaser in the same job.
#
# Usage: scripts/update-formula.sh <tag>
# Requires: HOMEBREW_TAP_TOKEN in the environment (push access to the tap).
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"
# shellcheck source=scripts/formula-lib.sh
source scripts/formula-lib.sh
# shellcheck source=scripts/release-lib.sh
source scripts/release-lib.sh

tag=${1:?usage: scripts/update-formula.sh <tag>}
version=${tag#v}
: "${HOMEBREW_TAP_TOKEN:?HOMEBREW_TAP_TOKEN is not set}"

is_release_tag "$tag" \
  || { echo "error: unexpected release tag $tag (want vX.Y.Z)" >&2; exit 1; }

owner="sorokin-vladimir"
tap="homebrew-tap"
base="https://github.com/${owner}/downprint/releases/download/${tag}"
checksums="dist/checksums.txt"

# sha256 for a release artifact, looked up by file name in checksums.txt.
sha() {
  local name=$1 line
  line=$(grep -E "  ${name}\$" "$checksums") ||
    {
      echo "checksum for $name not found in $checksums" >&2
      exit 1
    }
  echo "${line%% *}"
}

formula=$(render_formula "$owner" "$version" "$base" \
  "$(sha downprint_darwin_amd64.tar.gz)" \
  "$(sha downprint_darwin_arm64.tar.gz)" \
  "$(sha downprint_linux_amd64.tar.gz)" \
  "$(sha downprint_linux_arm64.tar.gz)")

workdir=$(mktemp -d)
trap 'rm -rf "$workdir"' EXIT

git clone --depth 1 \
  "https://x-access-token:${HOMEBREW_TAP_TOKEN}@github.com/${owner}/${tap}.git" \
  "$workdir"

mkdir -p "$workdir/Formula"
printf '%s\n' "$formula" >"$workdir/Formula/downprint.rb"

git -C "$workdir" config user.name "github-actions[bot]"
git -C "$workdir" config user.email "github-actions[bot]@users.noreply.github.com"

# Stage first, then compare the index against HEAD: `git diff --quiet` alone
# does not see untracked files, so the first publish would look unchanged.
git -C "$workdir" add Formula/downprint.rb
if git -C "$workdir" diff --cached --quiet; then
  echo "${tap}: downprint.rb already up to date for ${tag}"
else
  git -C "$workdir" commit -m "Brew formula update for downprint (${tag})"
  git -C "$workdir" push
  echo "${tap}: pushed downprint.rb for ${tag}"
fi
