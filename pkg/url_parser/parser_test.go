package url_test

import (
	"testing"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantHost  string
		wantOrg   string
		wantRepo  string
		wantType  UrlType
		wantNum   string
		wantError bool
	}{
		{
			name:     "valid issue URL",
			input:    "https://github.com/orgname/reponame/issues/123",
			wantHost: "github.com",
			wantOrg:  "orgname",
			wantRepo: "reponame",
			wantType: TypeIssue,
			wantNum:  "123",
		},
		{
			name:     "valid pull URL",
			input:    "https://github.com/orgname/reponame/pull/42",
			wantHost: "github.com",
			wantOrg:  "orgname",
			wantRepo: "reponame",
			wantType: TypePR,
			wantNum:  "42",
		},
		{
			name:      "unsupported host",
			input:     "https://gitlab.com/org/repo/issues/123",
			wantError: true,
		},
		{
			name:      "invalid format - too few segments",
			input:     "https://github.com/org/repo/issues",
			wantError: true,
		},
		{
			name:      "missing components",
			input:     "https://github.com///issues/123",
			wantError: true,
		},
		{
			name:      "unsupported type",
			input:     "https://github.com/org/repo/discussions/123",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.input)

			if tt.wantError {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.Host != tt.wantHost {
				t.Errorf("Host: got %q, want %q", got.Host, tt.wantHost)
			}
			if got.Org != tt.wantOrg {
				t.Errorf("Org: got %q, want %q", got.Org, tt.wantOrg)
			}
			if got.Repo != tt.wantRepo {
				t.Errorf("Repo: got %q, want %q", got.Repo, tt.wantRepo)
			}
			if got.Type != tt.wantType {
				t.Errorf("Type: got %q, want %q", got.Type, tt.wantType)
			}
			if got.Num != tt.wantNum {
				t.Errorf("Num: got %q, want %q", got.Num, tt.wantNum)
			}

			if !tt.wantError && got.ToUrl() != tt.input {
				t.Errorf("ToUrl: got %q, want %q", got.ToUrl(), tt.input)
			}
		})
	}
}
