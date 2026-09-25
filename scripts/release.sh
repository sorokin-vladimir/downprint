#!/usr/bin/env bash
# Local release driver: validate, update CHANGELOG, commit, and tag.
# Usage: scripts/release.sh <version|patch|minor|major>
# Does NOT push; prints the push command to run after review.
set -euo pipefail

cd "$(git rev-parse --show-toplevel)"
# shellcheck source=scripts/changelog-lib.sh
source scripts/changelog-lib.sh
# shellcheck source=scripts/release-lib.sh
source scripts/release-lib.sh

CHANGELOG="CHANGELOG.md"
die() { echo "error: $*" >&2; exit 1; }

[ $# -eq 1 ] || die "usage: scripts/release.sh <version|patch|minor|major>"

# Preconditions.
[ -z "$(git status --porcelain)" ] || die "working tree is not clean"
[ "$(git rev-parse --abbrev-ref HEAD)" = "main" ] || die "not on main branch"

latest=$(git tag -l 'v*' --sort=-version:refname | head -n1)
[ -n "$latest" ] || latest="v0.0.0"

version=$(compute_target_version "$latest" "$1")
tag="v$version"

is_release_tag "$tag" || die "release tag must be vX.Y.Z (got $tag)"
git rev-parse "$tag" >/dev/null 2>&1 && die "tag $tag already exists"
changelog_unreleased_has_content "$CHANGELOG" \
  || die "nothing to release: [Unreleased] in $CHANGELOG is empty"

confirm_release "$latest" "$tag" || die "aborted"

changelog_promote "$CHANGELOG" "$version" "$(date +%F)"
body=$(changelog_extract_body "$CHANGELOG" "$version")

git add "$CHANGELOG"
git commit -m "chore: release $tag"
printf '%s\n\n%s\n' "$tag" "$body" | git tag -a "$tag" -F -

echo "Tagged $tag."
echo "Review the commit and tag, then run: git push origin main --follow-tags"
