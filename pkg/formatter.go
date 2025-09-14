package pkg

import (
	"fmt"

	"github.com/enrico-kaack/github-link-formater/pkg/github"
	url_parser "github.com/enrico-kaack/github-link-formater/pkg/url_parser"
)

func Format(u url_parser.UrlParsed, ghResp *github.GhResponse) string {
	switch u.Type {
	case url_parser.TypeIssue:
		return formatIssue(u, ghResp)
	case url_parser.TypePR:
		return formatPR(u, ghResp)
	default:
		return "Unsupported type"
	}
}

func formatIssue(u url_parser.UrlParsed, ghResp *github.GhResponse) string {
	return fmt.Sprintf("[Issue (%s/%s): %s (%s-#%d)](%s)", u.Org, u.Repo, ghResp.Title, ghResp.State, ghResp.Number, u.ToUrl())
}

func formatPR(u url_parser.UrlParsed, ghResp *github.GhResponse) string {
	return fmt.Sprintf("[PR (%s/%s): %s (%s-#%d)](%s)", u.Org, u.Repo, ghResp.Title, ghResp.State, ghResp.Number, u.ToUrl())
}
