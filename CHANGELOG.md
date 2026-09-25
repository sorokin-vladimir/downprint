# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/).

A human title for a release is written as an em-dash suffix on its heading,
e.g. `## [1.2.0] - 2026-06-11 — Custom fonts`.

## [Unreleased]

## [1.0.0] - 2026-09-25 — First release

### Added

- Convert a Markdown file to PDF: CommonMark + GFM (tables, task lists,
  strikethrough, autolinks), syntax highlighting, raw HTML, and GitHub-like
  styling, printed by headless Chrome.
- Without `-o`, the PDF is written next to the input with the same name.
- `-paper` for the page size (named formats or `WIDTHxHEIGHT`) and
  `-landscape` for the orientation.
- `-css` to apply an extra stylesheet on top of the built-in one.
- `-chrome` and `DOWNPRINT_CHROME` to point at a specific browser binary. Without
  them, Chrome or Chromium is found automatically, with Microsoft Edge as a
  fallback on Windows and macOS.
- Page numbers in the footer; relative images resolve against the Markdown
  file; collapsed `<details>` blocks are expanded before printing.
