# Contributing to downprint

Thanks for your interest in improving `downprint`. This document explains how to set
up a development environment, the conventions the project follows, and how to
get your changes merged.

## Code of Conduct

This project follows the [Contributor Covenant](../CODE_OF_CONDUCT.md). By
participating, you are expected to uphold it.

## Ways to contribute

- **Report a bug**: open an issue using the [Bug report](ISSUE_TEMPLATE/bug_report.md) template.
- **Request a feature**: open an issue using the [Feature request](ISSUE_TEMPLATE/feature_request.md) template.
- **Submit a fix or feature**: see the workflow below.

For anything non-trivial, please open an issue first so we can agree on the
approach before you spend time on a pull request.

## Development setup

Requires **Go 1.26+** and Google Chrome or Chromium.

This repo uses [mise](https://mise.jdx.dev/) for tooling and tasks, and
[lefthook](https://github.com/evilmartians/lefthook) for git hooks.

```sh
git clone https://github.com/sorokin-vladimir/downprint
cd downprint

mise install        # installs the pinned Go toolchain
mise run hooks      # installs git hooks via lefthook
```

Render the sample document that covers every block type:

```sh
mise run sample     # writes bin/all-blocks.pdf
```

## Before you open a PR

Run the full check locally. This is also what the git hooks enforce on push:

```sh
mise run check      # go vet ./... + golangci-lint run ./... + go test ./...
```

| Task                   | Command                              |
| ---------------------- | ------------------------------------ |
| `mise run test`        | `go test ./...`                      |
| `mise run test-scripts`| release script tests in `scripts/test` |
| `mise run lint`        | `golangci-lint run ./...`            |
| `mise run vet`         | `go vet ./...`                       |
| `mise run tidy`        | `go mod tidy`                        |

Code is formatted with `gofmt` (run automatically on commit via lefthook).

Rendering changes are hard to cover with unit tests: check them by eye on
`testdata/all-blocks.md`, and extend that file when you add support for a new
kind of block.

## Commit messages

Commits follow a [Conventional Commits](https://www.conventionalcommits.org/)
style, referencing the related issue where one exists:

```
fix: #12 Keep table rows on one page
feat: #15 Table of contents
docs: Update changelog
chore: release v0.2.0
```

Common types: `feat`, `fix`, `docs`, `refactor`, `test`, `chore`.

## Pull request process

1. Fork the repo and create a branch from `main`.
2. Make your change, including tests and doc updates where relevant.
3. Run `mise run check` and make sure it passes.
4. Update [`CHANGELOG.md`](../CHANGELOG.md) under the unreleased section if the
   change is user-facing.
5. Open the PR and fill in the template. Link the issue it resolves.

## Releases

Maintainers cut releases with `mise run release <version|patch|minor|major>`.
It promotes the `[Unreleased]` changelog section, commits, and tags locally.
Pushing the tag runs the release workflow: checks, GoReleaser (archives,
deb/rpm/apk), and the Homebrew formula update.

## License

By contributing, you agree that your contributions will be licensed under the
[MIT](../LICENSE) license that covers the project.
