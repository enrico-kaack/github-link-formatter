package url_test

import (
	"fmt"
	"net/url"
	"strings"
)

type UrlParsed struct {
	Host string
	Type UrlType
	Org  string
	Repo string
	Num  string
}

func (u *UrlParsed) ToUrl() string {
	return fmt.Sprintf("https://github.com/%s/%s/%s/%s", u.Org, u.Repo, u.Type, u.Num)
}

type UrlType string

const (
	TypeIssue UrlType = "issues"
	TypePR    UrlType = "pull"
)

func ParseURL(urlStr string) (*UrlParsed, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	if u.Host != "github.com" {
		return nil, fmt.Errorf("unsupported host: %s. Only github.com is supported", u.Host)
	}
	s := strings.Split(u.Path, "/")
	if len(s) < 5 {
		return nil, fmt.Errorf("invalid URL format")
	}
	org := s[1]
	repo := s[2]
	t := s[3]
	num := s[4]
	if org == "" || repo == "" || t == "" || num == "" {
		return nil, fmt.Errorf("missing required URL components")
	}

	var urlType UrlType
	switch t {
	case "issues":
		urlType = TypeIssue
	case "pull":
		urlType = TypePR
	default:
		return nil, fmt.Errorf("unsupported type: %s. Only issues and pull are supported", t)
	}
	return &UrlParsed{
		Host: u.Host,
		Type: urlType,
		Org:  org,
		Repo: repo,
		Num:  num,
	}, nil

}
