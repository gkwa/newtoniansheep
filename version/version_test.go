package version

import (
	"testing"
)

func TestBuildInfo_String(t *testing.T) {
	tests := []struct {
		name     string
		info     BuildInfo
		expected string
	}{
		{
			name: "Full build info",
			info: BuildInfo{
				Date:        "2023-05-01",
				FullGitSHA:  "abcdef1234567890",
				GoVersion:   "go1.16",
				ShortGitSHA: "abcdef1",
				Version:     "v1.0.0",
			},
			expected: `Version: v1.0.0, abcdef1234567890
Build Date: 2023-05-01
Go Version: go1.16`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.info.String() != tt.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tt.expected, tt.info.String())
			}
		})
	}
}

func TestGetBuildInfo(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		fullSHA  string
		goVer    string
		shortSHA string
		version  string
	}{
		{
			name:     "Default build info",
			date:     "2023-05-01",
			fullSHA:  "abcdef1234567890",
			goVer:    "go1.16",
			shortSHA: "abcdef1",
			version:  "v1.0.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Date = tt.date
			FullGitSHA = tt.fullSHA
			GoVersion = tt.goVer
			ShortGitSHA = tt.shortSHA
			Version = tt.version

			bi := GetBuildInfo()

			if bi.Date != tt.date {
				t.Errorf("Expected Date %s, got %s", tt.date, bi.Date)
			}
			if bi.FullGitSHA != tt.fullSHA {
				t.Errorf("Expected FullGitSHA %s, got %s", tt.fullSHA, bi.FullGitSHA)
			}
			if bi.GoVersion != tt.goVer {
				t.Errorf("Expected GoVersion %s, got %s", tt.goVer, bi.GoVersion)
			}
			if bi.ShortGitSHA != tt.shortSHA {
				t.Errorf("Expected ShortGitSHA %s, got %s", tt.shortSHA, bi.ShortGitSHA)
			}
			if bi.Version != tt.version {
				t.Errorf("Expected Version %s, got %s", tt.version, bi.Version)
			}
		})
	}
}
