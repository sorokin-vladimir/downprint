package main

import (
	"context"
	"os"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// printPDF loads the HTML page in headless Chrome and prints it to PDF.
// chromePath may be empty to autodetect the browser.
func printPDF(ctx context.Context, html []byte, chromePath string) ([]byte, error) {
	// A temp file (rather than setDocumentContent) lets file:// images
	// referenced via <base href> load normally.
	f, err := os.CreateTemp("", "downprint-*.html")
	if err != nil {
		return nil, err
	}
	// Best-effort cleanup: the write error, if any, is the one worth reporting.
	defer func() { _ = os.Remove(f.Name()) }()
	if _, err := f.Write(html); err != nil {
		_ = f.Close()
		return nil, err
	}
	if err := f.Close(); err != nil {
		return nil, err
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("allow-file-access-from-files", true),
	)
	if chromePath == "" {
		chromePath = findBrowser()
	}
	// Still empty: fall back to chromedp's own lookup.
	if chromePath != "" {
		opts = append(opts, chromedp.ExecPath(chromePath))
	}

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(ctx, opts...)
	defer cancelAlloc()
	browserCtx, cancelBrowser := chromedp.NewContext(allocCtx)
	defer cancelBrowser()

	var pdf []byte
	err = chromedp.Run(browserCtx,
		chromedp.Navigate(string(fileURL(f.Name()))),
		// Collapsed <details> would hide their content on paper.
		chromedp.Evaluate(`document.querySelectorAll("details").forEach(d => d.open = true)`, nil),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			// Paper size and margins come from the @page rule in CSS.
			pdf, _, err = page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				WithDisplayHeaderFooter(true).
				WithHeaderTemplate(`<div></div>`).
				WithFooterTemplate(`<div style="width:100%;font-size:8px;color:#888;text-align:center;">` +
					`<span class="pageNumber"></span> / <span class="totalPages"></span></div>`).
				Do(ctx)
			return err
		}),
	)
	return pdf, err
}
