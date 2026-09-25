package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// findBrowser returns the first Chromium-based browser found in the standard
// locations, or "" to leave the lookup to chromedp. It differs from chromedp's
// own lookup by falling back to Microsoft Edge, which ships with Windows, so
// downprint works there without installing Chrome. Linux is left to chromedp,
// whose list of package names is broader.
func findBrowser() string {
	for _, p := range browserCandidates() {
		if found, err := exec.LookPath(p); err == nil {
			return found
		}
	}
	return ""
}

func browserCandidates() []string {
	switch runtime.GOOS {
	case "windows":
		local := filepath.Join(os.Getenv("USERPROFILE"), `AppData\Local`)
		return []string{
			"chrome.exe",
			`C:\Program Files\Google\Chrome\Application\chrome.exe`,
			`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
			filepath.Join(local, `Google\Chrome\Application\chrome.exe`),
			filepath.Join(local, `Chromium\Application\chrome.exe`),
			`C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe`,
			`C:\Program Files\Microsoft\Edge\Application\msedge.exe`,
		}
	case "darwin":
		return []string{
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
			"/Applications/Chromium.app/Contents/MacOS/Chromium",
			"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
		}
	}
	return nil
}
