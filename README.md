# md2pdf

Converts Markdown to PDF with GitHub-like styling. Markdown is rendered to HTML with [goldmark](https://github.com/yuin/goldmark) (GFM, syntax highlighting via chroma), then printed to PDF by headless Chrome through [chromedp](https://github.com/chromedp/chromedp).

Pure Go, no cgo. At runtime it needs Chrome or Chromium installed.

## Install

```sh
go install .
```

or build a binary:

```sh
CGO_ENABLED=0 go build -o md2pdf .
```

## Usage

```sh
md2pdf [flags] input.md
```

Without `-o`, the PDF is written next to the input with the same name (`notes.md` -> `notes.pdf`).

| Flag | Description |
|------|-------------|
| `-o path` | Output PDF path |
| `-paper size` | `a3`, `a4`, `a5`, `b4`, `b5`, `letter`, `legal`, `ledger`, `tabloid`, or `WIDTHxHEIGHT` such as `210mmx297mm` (units: `mm`, `cm`, `in`, `pt`, `px`). Default: A4 |
| `-landscape` | Landscape orientation |
| `-css path` | Extra stylesheet applied on top of the built-in one |
| `-chrome path` | Chrome/Chromium binary. Also read from `MD2PDF_CHROME`. Default: autodetect |
| `-timeout d` | Conversion timeout. Default: `1m` |

Examples:

```sh
md2pdf README.md
md2pdf -o out/report.pdf -paper letter -landscape report.md
md2pdf -css testdata/custom.css testdata/all-blocks.md
```

## Styling

The built-in stylesheet lives in `assets/style.css` and is embedded into the binary. A stylesheet passed with `-css` is loaded after it, so it only needs to override what you want to change. Relative `url()` references in it resolve against the CSS file's directory.

Page size and margins come from the CSS `@page` rule. `-paper` and `-landscape` take precedence over any `@page { size }` in the stylesheets.

Page numbers are printed in the footer. Relative images and links resolve against the Markdown file's directory. Collapsed `<details>` blocks are expanded before printing.

## Tests

```sh
go test ./...
```

`testdata/all-blocks.md` covers the main Markdown and GFM blocks and is useful for checking rendering by eye.
