package template

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/enrico-kaack/github-link-formater/pkg/github"
	url_parser "github.com/enrico-kaack/github-link-formater/pkg/url_parser"
)

type TemplateEngine struct {
	templates map[url_parser.UrlType]*template.Template
}

type templateData struct {
	Number int
	Org    string
	Repo   string
	Title  string
	State  string
	Url    string
	Type   string
}

const (
	defaultIssueTemplate = "[Issue ({{.Org}}/{{.Repo}}): {{.Title}} ({{.State}}-#{{.Number}})]({{.Url}})"
	defaultPRTemplate    = "[PR ({{.Org}}/{{.Repo}}): {{.Title}} ({{.State}}-#{{.Number}})]({{.Url}})"
)

func parseTemplate(name, content string) (*template.Template, error) {
	return template.New(name).Parse(content)
}

func NewDefaultTemplateEngine() (*TemplateEngine, error) {
	issue, err := parseTemplate("issue", defaultIssueTemplate)
	if err != nil {
		return nil, err
	}
	pr, err := parseTemplate("pr", defaultPRTemplate)
	if err != nil {
		return nil, err
	}
	return &TemplateEngine{
		templates: map[url_parser.UrlType]*template.Template{
			url_parser.TypeIssue: issue,
			url_parser.TypePR:    pr,
		},
	}, nil
}

func NewTemplateEngineFromDirFolderOrDefault(dir string) (*TemplateEngine, error) {
	engine, err := NewDefaultTemplateEngine()
	if err != nil {
		return nil, err
	}

	loadTemplate := func(filename string, tType url_parser.UrlType) error {
		path := filepath.Join(dir, filename)
		if _, err := os.Stat(path); err == nil {
			t, err := template.ParseFiles(path)
			if err != nil {
				return fmt.Errorf("failed to parse %s: %w", path, err)
			}
			engine.templates[tType] = t
		}
		return nil
	}

	if err := loadTemplate("issue.tmpl", url_parser.TypeIssue); err != nil {
		return nil, err
	}
	if err := loadTemplate("pr.tmpl", url_parser.TypePR); err != nil {
		return nil, err
	}

	return engine, nil
}

func NewTemplateEngine(issueTemplate, prTemplate string) (*TemplateEngine, error) {
	iT, err := template.New("issue").Parse(issueTemplate)
	if err != nil {
		return nil, err
	}
	pT, err := template.New("pr").Parse(prTemplate)
	if err != nil {
		return nil, err
	}
	return &TemplateEngine{
		templates: map[url_parser.UrlType]*template.Template{
			url_parser.TypeIssue: iT,
			url_parser.TypePR:    pT,
		},
	}, nil
}

func (te *TemplateEngine) Format(u *url_parser.UrlParsed, ghResp *github.GhResponse) (string, error) {
	tmpl, ok := te.templates[u.Type]
	if !ok {
		return "", fmt.Errorf("no template found for type: %s", u.Type)
	}

	data := templateData{
		Number: ghResp.Number,
		Title:  ghResp.Title,
		State:  ghResp.State,
		Url:    u.ToUrl(),
		Type:   string(u.Type),
		Org:    u.Org,
		Repo:   u.Repo,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("error rendering %s template: %w", u.Type, err)
	}
	return buf.String(), nil
}
