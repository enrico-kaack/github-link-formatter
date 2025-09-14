package template

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/enrico-kaack/github-link-formatter/pkg/github"
	url_parser "github.com/enrico-kaack/github-link-formatter/pkg/url_parser"
)

func TestTemplateEngine_Format(t *testing.T) {
	tempDir := t.TempDir()

	// create custom template files
	customIssue := "CUSTOM ISSUE {{.Number}}-{{.Org}}-{{.Repo}}-{{.Title}}-{{.State}}-{{.Url}}-{{.Type}}"
	customPR := "CUSTOM PR {{.Number}}-{{.Org}}-{{.Repo}}-{{.Title}}-{{.State}}-{{.Url}}-{{.Type}}"

	issuePath := filepath.Join(tempDir, "issue.tmpl")
	prPath := filepath.Join(tempDir, "pr.tmpl")

	if err := os.WriteFile(issuePath, []byte(customIssue), 0644); err != nil {
		t.Fatalf("failed to write custom issue template: %v", err)
	}
	if err := os.WriteFile(prPath, []byte(customPR), 0644); err != nil {
		t.Fatalf("failed to write custom pr template: %v", err)
	}

	tests := []struct {
		name          string
		engineFactory func() *TemplateEngine
		urlType       url_parser.UrlType
		wants         string
		expectError   bool
	}{
		{
			name: "default issue template",
			engineFactory: func() *TemplateEngine {
				e, _ := NewDefaultTemplateEngine()
				return e
			},
			urlType: url_parser.TypeIssue,
			wants:   "[Issue (test-org/test-repo): Test Title (open-#42)](https://github.com/test-org/test-repo/issues/42)",
		},
		{
			name: "default pr template",
			engineFactory: func() *TemplateEngine {
				e, _ := NewDefaultTemplateEngine()
				return e
			},
			urlType: url_parser.TypePR,
			wants:   "[PR (test-org/test-repo): Test Title (open-#42)](https://github.com/test-org/test-repo/pull/42)",
		},
		{
			name: "custom templates from dir - issue",
			engineFactory: func() *TemplateEngine {
				e, err := NewTemplateEngineFromDirFolderOrDefault(tempDir)
				if err != nil {
					t.Fatalf("failed to load from dir: %v", err)
				}
				return e
			},
			urlType: url_parser.TypeIssue,
			wants:   "CUSTOM ISSUE 42-test-org-test-repo-Test Title-open-https://github.com/test-org/test-repo/issues/42-issue",
		},
		{
			name: "custom templates from dir - pr",
			engineFactory: func() *TemplateEngine {
				e, err := NewTemplateEngineFromDirFolderOrDefault(tempDir)
				if err != nil {
					t.Fatalf("failed to load from dir: %v", err)
				}
				return e
			},
			urlType: url_parser.TypePR,
			wants:   "CUSTOM PR 42-test-org-test-repo-Test Title-open-https://github.com/test-org/test-repo/pull/42-pull",
		},
		{
			name: "unsupported type should error",
			engineFactory: func() *TemplateEngine {
				e, _ := NewDefaultTemplateEngine()
				return e
			},
			urlType:     "unsupported", // not in map
			expectError: true,
		},
	}

	// shared test data
	parsed := url_parser.UrlParsed{
		Host: "github.com",
		Type: url_parser.TypeIssue, // will be overridden in tests
		Org:  "test-org",
		Repo: "test-repo",
		Num:  "42",
	}
	resp := &github.GhResponse{
		Number: 42,
		Title:  "Test Title",
		State:  "open",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := tt.engineFactory()
			parsed.Type = tt.urlType
			got, err := engine.Format(&parsed, resp)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.wants {
				t.Errorf("expected output to be %q, got %q", tt.wants, got)
			}
		})
	}
}
