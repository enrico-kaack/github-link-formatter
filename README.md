# Github Link Formater

A small tool to format GitHub links for issues and PRs to repos into readable formats (like Markdown links).

## Features
- Fetch issue and PR details from GitHub API to enrich the formatted output (only public projects are supported).
- Format GitHub issue and PR links into Markdown format by default.
- Support custom templates for formatting.

## Installation
You can install the tool using `go install`:

```bash
go install github.com/enrico-kaack/github-link-formatter@latest
```

## Usage
Run the tool with a GitHub issue or PR URL as an argument:
```bash
github-link-formatter <url>
```

Combined with tools such as `pbpaste` and `pbcopy` on macOS, you can easily format links from your clipboard:
```bash
github-link-formatter $(pbpaste ) | pbcopy
```

## Configuration

You can customize the output format using `text/template` files. The tool will look for templates in the folder `HOME/.github-link-formatter/`. You can specify different templates for issues and PRs by naming them `issue.tmpl` and `pr.tmpl` respectively.

You can use the [text/template](https://pkg.go.dev/text/template) syntax to create your templates. The following fields are available in the template context:
- `Org`: The organization or user that owns the repository.
- `Repo`: The repository name.
- `Number`: The issue or PR number.
- `Title`: The title of the issue or PR.
- `State`: The state of the issue or PR (e.g., `open`, `closed`).
- `Url`: The URL of the issue or PR.
- `Type`: The type of the URL (`issue` or `pull`).

See the `example/` folder for sample templates.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.