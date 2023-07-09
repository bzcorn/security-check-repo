package checker

import (
	"testing"
)

func TestSplitRepoName(t *testing.T) {
	tests := []struct {
		repo string
		want []string
	}{
		{"owner/repo", []string{"owner", "repo"}},
		{"username/project", []string{"username", "project"}},
	}

	for _, tt := range tests {
		got := splitRepoName(tt.repo)
		for i, val := range got {
			if val != tt.want[i] {
				t.Errorf("splitRepoName() = %v, want %v", got, tt.want)
			}
		}
	}
}
