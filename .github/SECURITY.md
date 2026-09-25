# Security Policy

## Supported versions

Security fixes are applied to the latest released version. Please make sure you
are on the most recent [release](https://github.com/sorokin-vladimir/downprint/releases)
before reporting.

## Reporting a vulnerability

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, report them privately through GitHub's
[Security Advisories](https://github.com/sorokin-vladimir/downprint/security/advisories/new)
("Report a vulnerability"). If that is not possible, email
**v.sorokin@hey.com** with the details.

Please include:

- a description of the issue and its impact,
- steps to reproduce or a proof of concept,
- the version of `downprint`, your OS, and the Chrome/Chromium version.

You can expect an initial response within a few days.

## Untrusted input

`downprint` renders raw HTML from the Markdown source, as GitHub does, and loads
the result in a local headless Chrome with access to local files. Only convert
documents you trust: a hostile document can reference remote resources or read
local files into the PDF.
