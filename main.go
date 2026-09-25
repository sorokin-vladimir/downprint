package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type config struct {
	input     string
	output    string
	css       string
	paper     string
	landscape bool
	chrome    string
	timeout   time.Duration
}

func main() {
	var cfg config
	flag.StringVar(&cfg.output, "o", "", "output PDF path (default: next to input, same name)")
	flag.StringVar(&cfg.css, "css", "", "extra stylesheet applied on top of the built-in one")
	flag.StringVar(&cfg.paper, "paper", "", "paper size: "+paperNames()+" or WIDTHxHEIGHT, e.g. 210mmx297mm (default: a4)")
	flag.BoolVar(&cfg.landscape, "landscape", false, "landscape orientation")
	flag.StringVar(&cfg.chrome, "chrome", os.Getenv("MD2PDF_CHROME"), "path to Chrome/Chromium binary (default: autodetect)")
	flag.DurationVar(&cfg.timeout, "timeout", 60*time.Second, "conversion timeout")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: md2pdf [flags] input.md\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}
	cfg.input = flag.Arg(0)

	if err := run(cfg); err != nil {
		fmt.Fprintln(os.Stderr, "md2pdf:", err)
		os.Exit(1)
	}
}

func run(cfg config) error {
	pageCSS, err := pageSizeCSS(cfg.paper, cfg.landscape)
	if err != nil {
		return err
	}

	input, err := filepath.Abs(cfg.input)
	if err != nil {
		return err
	}
	src, err := os.ReadFile(input)
	if err != nil {
		return err
	}

	output := cfg.output
	if output == "" {
		output = strings.TrimSuffix(input, filepath.Ext(input)) + ".pdf"
	}

	var userCSS string
	if cfg.css != "" {
		if userCSS, err = filepath.Abs(cfg.css); err != nil {
			return err
		}
		// Chrome silently ignores a missing stylesheet, so check it here.
		if _, err := os.Stat(userCSS); err != nil {
			return err
		}
	}

	html, err := renderHTML(src, pageOptions{
		Title:   strings.TrimSuffix(filepath.Base(input), filepath.Ext(input)),
		BaseDir: filepath.Dir(input),
		UserCSS: userCSS,
		PageCSS: pageCSS,
	})
	if err != nil {
		return fmt.Errorf("render markdown: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.timeout)
	defer cancel()

	pdf, err := printPDF(ctx, html, cfg.chrome)
	if err != nil {
		return fmt.Errorf("print pdf: %w", err)
	}

	if err := os.WriteFile(output, pdf, 0o644); err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
