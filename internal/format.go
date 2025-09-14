package internal

import (
	"github.com/enrico-kaack/github-link-formatter/pkg/github"
	"github.com/enrico-kaack/github-link-formatter/pkg/template"
	url_parser "github.com/enrico-kaack/github-link-formatter/pkg/url_parser"
)

func FormatGHLink(u string, templateFolder string) (string, error) {
	result, err := url_parser.ParseURL(u)
	if err != nil {
		return "", err
	}
	ghResponse, err := github.GetFromGHApi(*result)
	if err != nil {
		return "", err
	}
	templateEngine, err := template.NewTemplateEngineFromDirFolderOrDefault(templateFolder)
	if err != nil {
		return "", err
	}
	formatted, err := templateEngine.Format(result, ghResponse)
	if err != nil {
		return "", err
	}
	return formatted, nil
}
