# downprint

Converts Markdown to PDF with GitHub-like styling. Markdown is rendered to HTML with [goldmark](https://github.com/yuin/goldmark) (GFM, syntax highlighting via chroma), then printed to PDF by headless Chrome through [chromedp](https://github.com/chromedp/chromedp).

Pure Go, no cgo. At runtime it needs a Chromium-based browser (Chrome, Chromium, or Edge).

## Install

Homebrew (macOS, Linux). Homebrew only loads formulae from third-party taps you have trusted, so trust the tap first:

```sh
brew trust --tap sorokin-vladimir/tap
brew tap sorokin-vladimir/tap
brew install sorokin-vladimir/tap/downprint
```

Go:

```sh
go install github.com/sorokin-vladimir/downprint@latest
```

Prebuilt binaries, archives, and deb/rpm/apk packages are attached to each [release](https://github.com/sorokin-vladimir/downprint/releases).

downprint needs a Chromium-based browser at runtime: Google Chrome, Chromium, or Microsoft Edge (Windows and macOS). On Windows the preinstalled Edge is enough. On macOS: `brew install --cask google-chrome`. Browsers in standard locations are found automatically; otherwise use `-chrome` or `DOWNPRINT_CHROME`.

## Usage

```sh
downprint [flags] input.md
```

Without `-o`, the PDF is written next to the input with the same name (`notes.md` -> `notes.pdf`).

| Flag           | Description                                                                                                                                                      |
| -------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `-o path`      | Output PDF path                                                                                                                                                  |
| `-paper size`  | `a3`, `a4`, `a5`, `b4`, `b5`, `letter`, `legal`, `ledger`, `tabloid`, or `WIDTHxHEIGHT` such as `210mmx297mm` (units: `mm`, `cm`, `in`, `pt`, `px`). Default: A4 |
| `-landscape`   | Landscape orientation                                                                                                                                            |
| `-css path`    | Extra stylesheet applied on top of the built-in one                                                                                                              |
| `-chrome path` | Chrome/Chromium binary. Also read from `DOWNPRINT_CHROME`. Default: autodetect                                                                                   |
| `-timeout d`   | Conversion timeout. Default: `1m`                                                                                                                                |
| `-version`     | Print version and exit                                                                                                                                           |

Examples:

```sh
downprint README.md
downprint -o out/report.pdf -paper letter -landscape report.md
downprint -css testdata/custom.css testdata/all-blocks.md
```

## Styling

The built-in stylesheet lives in `assets/style.css` and is embedded into the binary. A stylesheet passed with `-css` is loaded after it, so it only needs to override what you want to change. Relative `url()` references in it resolve against the CSS file's directory.

Page size and margins come from the CSS `@page` rule. `-paper` and `-landscape` take precedence over any `@page { size }` in the stylesheets.

Page numbers are printed in the footer. Relative images and links resolve against the Markdown file's directory. Collapsed `<details>` blocks are expanded before printing.

## Development

```sh
mise run check      # vet + lint + tests
mise run sample     # render testdata/all-blocks.md to bin/all-blocks.pdf
```

`testdata/all-blocks.md` covers the main Markdown and GFM blocks and is useful for checking rendering by eye. See [CONTRIBUTING](.github/CONTRIBUTING.md) for setup, conventions, and the release process.

## License

[MIT](LICENSE)
