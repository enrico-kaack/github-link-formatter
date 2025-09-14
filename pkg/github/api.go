package github

import (
	"encoding/json"
	"fmt"
	"net/http"

	url_parser "github.com/enrico-kaack/github-link-formater/pkg/url_parser"
)

func GetFromGHApi(u url_parser.UrlParsed) (*GhResponse, error) {
	url, err := getApiUrl(u)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch issue: %s", resp.Status)
	}
	defer resp.Body.Close()

	var ghResp GhResponse
	err = json.NewDecoder(resp.Body).Decode(&ghResp)
	if err != nil {
		return nil, err
	}
	return &ghResp, nil
}

func getApiUrl(u url_parser.UrlParsed) (string, error) {
	url := ""
	switch u.Type {
	case url_parser.TypeIssue:
		url = issueUrl(u.Org, u.Repo, u.Num)
	case url_parser.TypePR:
		url = prUrl(u.Org, u.Repo, u.Num)
	default:
		return "", fmt.Errorf("unsupported type: %s. Only issues and pull requests are supported", u.Type)
	}
	return url, nil
}

func issueUrl(org, repo, num string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", org, repo, num)
}

func prUrl(org, repo, num string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls/%s", org, repo, num)
}
