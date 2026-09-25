package main

import "testing"

func TestFileURL(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/Users/me/docs/", "file:///Users/me/docs/"},
		{"/tmp/a b/c.css", "file:///tmp/a%20b/c.css"},
		// Windows drive paths, as filepath.ToSlash leaves them.
		{"C:/Users/me/docs/", "file:///C:/Users/me/docs/"},
	}
	for _, tt := range tests {
		if got := string(fileURL(tt.path)); got != tt.want {
			t.Errorf("fileURL(%q) = %q, want %q", tt.path, got, tt.want)
		}
	}
}
