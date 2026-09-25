package main

import "testing"

func TestPageSizeCSS(t *testing.T) {
	tests := []struct {
		paper     string
		landscape bool
		want      string
		wantErr   bool
	}{
		{"", false, "", false},
		{"", true, "@page { size: A4 landscape; }", false},
		{"A5", false, "@page { size: A5; }", false},
		{"letter", true, "@page { size: letter landscape; }", false},
		{"tabloid", true, "@page { size: 17in 11in; }", false},
		{"210mmx297mm", false, "@page { size: 210mm 297mm; }", false},
		{"8.5inx11in", true, "@page { size: 11in 8.5in; }", false},
		{"a7", false, "", true},
		{"210x297", false, "", true},
	}
	for _, tt := range tests {
		got, err := pageSizeCSS(tt.paper, tt.landscape)
		if (err != nil) != tt.wantErr {
			t.Errorf("pageSizeCSS(%q, %v) error = %v, wantErr %v", tt.paper, tt.landscape, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("pageSizeCSS(%q, %v) = %q, want %q", tt.paper, tt.landscape, got, tt.want)
		}
	}
}
