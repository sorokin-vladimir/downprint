package main

import (
	"fmt"
	"regexp"
	"strings"
)

// paperSizes maps flag values to CSS @page size keywords.
var paperSizes = map[string]string{
	"a3":      "A3",
	"a4":      "A4",
	"a5":      "A5",
	"b4":      "B4",
	"b5":      "B5",
	"letter":  "letter",
	"legal":   "legal",
	"ledger":  "ledger",
	"tabloid": "11in 17in",
}

// customSize matches WIDTHxHEIGHT with CSS units, e.g. 210mmx297mm or 8.5inx11in.
var customSize = regexp.MustCompile(`^(\d+(?:\.\d+)?(?:mm|cm|in|pt|px))x(\d+(?:\.\d+)?(?:mm|cm|in|pt|px))$`)

// pageSizeCSS returns an @page rule for the given paper name and orientation.
// It returns an empty string when neither is set, leaving the size to CSS.
func pageSizeCSS(paper string, landscape bool) (string, error) {
	paper = strings.ToLower(strings.TrimSpace(paper))
	if paper == "" && !landscape {
		return "", nil
	}
	if paper == "" {
		paper = "a4"
	}

	size, ok := paperSizes[paper]
	if !ok {
		m := customSize.FindStringSubmatch(paper)
		if m == nil {
			return "", fmt.Errorf("unknown paper size %q (use %s or WIDTHxHEIGHT like 210mmx297mm)", paper, paperNames())
		}
		size = m[1] + " " + m[2]
	}

	if landscape {
		if strings.Contains(size, " ") {
			// Explicit dimensions: swap width and height.
			w, h, _ := strings.Cut(size, " ")
			size = h + " " + w
		} else {
			size += " landscape"
		}
	}
	return "@page { size: " + size + "; }", nil
}

func paperNames() string {
	return "a3, a4, a5, b4, b5, letter, legal, ledger, tabloid"
}
