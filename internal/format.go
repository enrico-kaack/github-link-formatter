package internal

import (
	"github.com/enrico-kaack/github-link-formater/pkg"
	"github.com/enrico-kaack/github-link-formater/pkg/github"
	url_parser "github.com/enrico-kaack/github-link-formater/pkg/url_parser"
)

func FormatGHLink(u string) (string, error) {
	result, err := url_parser.ParseURL(u)
	if err != nil {
		return "", err
	}
	ghResponse, err := github.GetFromGHApi(*result)
	if err != nil {
		return "", err
	}
	formatted := pkg.Format(*result, ghResponse)
	return formatted, nil
}
