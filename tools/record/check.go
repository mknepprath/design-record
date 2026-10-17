package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

var (
	datePattern     = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	numberedName    = regexp.MustCompile(`^(\d{4})-[a-z0-9-]+\.md$`)
	datedName       = regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})-[a-z0-9-]+\.md$`)
	plainName       = regexp.MustCompile(`^[a-z0-9-]+\.md$`)
	mdLink          = regexp.MustCompile(`\]\(([^)\s]+)\)`)
	fencedCodeBlock = regexp.MustCompile("(?s)```.*?```")
)

// check returns one message per broken rule, each prefixed with the file.
func check(root string) ([]string, error) {
	docs, err := loadDocs(root)
	if err != nil {
		return nil, err
	}
	byPath := map[string]doc{}
	for _, d := range docs {
		byPath[d.Path] = d
	}
	var problems []string
	add := func(d doc, format string, args ...any) {
		problems = append(problems, d.Path+": "+fmt.Sprintf(format, args...))
	}

	numbers := map[string]map[int]string{}
	for _, d := range docs {
		k := d.Kind
		name := filepath.Base(d.Path)
		if !d.HasFront {
			add(d, "missing frontmatter (--- block at the top)")
			continue
		}
		if got := d.Fields["kind"]; got != k.Name {
			add(d, "kind is %q, but files in %s/ must be kind %q", got, k.Dir, k.Name)
		}
		if d.Title() == "" {
			add(d, "missing title")
		}
		if !datePattern.MatchString(d.Date()) {
			add(d, "date %q is not YYYY-MM-DD", d.Date())
		}
		if k.NeedsOwner && d.Fields["owner"] == "" {
			add(d, "missing owner")
		}
		if len(k.Statuses) > 0 && !slices.Contains(k.Statuses, d.Status()) {
			add(d, "status %q is not one of %s", d.Status(), strings.Join(k.Statuses, ", "))
		}
		if len(k.Statuses) == 0 && d.Status() != "" {
			add(d, "kind %s has no status; remove the status field", k.Name)
		}
		if k.NeedsSource && d.Fields["source"] == "" {
			add(d, "missing source (who or what, and when)")
		}

		switch {
		case k.Numbered:
			m := numberedName.FindStringSubmatch(name)
			if m == nil {
				add(d, "file name must be NNNN-slug.md")
				break
			}
			if fmt.Sprintf("%04d", d.Number()) != m[1] {
				add(d, "number field %q does not match file name %s", d.Fields["number"], m[1])
			}
			if numbers[k.Name] == nil {
				numbers[k.Name] = map[int]string{}
			}
			if other, dup := numbers[k.Name][d.Number()]; dup {
				add(d, "number %04d is also used by %s", d.Number(), other)
			}
			numbers[k.Name][d.Number()] = d.Path
		case k.Dated:
			m := datedName.FindStringSubmatch(name)
			if m == nil {
				add(d, "file name must be YYYY-MM-DD-slug.md")
			} else if m[1] != d.Date() {
				add(d, "file name date %s does not match date field %s", m[1], d.Date())
			}
		default:
			if !plainName.MatchString(name) {
				add(d, "file name must be lowercase-slug.md")
			}
		}

		if k.NeedsProject && len(d.Projects()) == 0 {
			add(d, "must name at least one project in projects")
		}
		for _, p := range d.Projects() {
			if pd, ok := byPath["projects/"+p+".md"]; !ok || pd.Kind.Name != "project" {
				add(d, "project %q does not exist (expected projects/%s.md)", p, p)
			}
		}
		if k.NeedsEvidence && len(d.Evidence()) == 0 {
			add(d, "must cite at least one evidence record in evidence")
		}
		for _, e := range d.Evidence() {
			ed, ok := byPath[e]
			if !ok {
				add(d, "cited evidence %s does not exist", e)
			} else if ed.Kind.Layer != layerEvidence {
				add(d, "cited %s is a %s, not an evidence record", e, ed.Kind.Name)
			}
		}
		for _, s := range d.Lists["supersedes"] {
			if _, ok := byPath[s]; !ok {
				add(d, "superseded document %s does not exist", s)
			}
		}

		problems = append(problems, checkLinks(root, d.Path, d.Body)...)
	}

	// Evidence nobody cites or links to is not backing anything.
	cited := map[string]bool{}
	for _, d := range docs {
		for _, e := range d.Evidence() {
			cited[e] = true
		}
		for _, target := range linkTargets(d.Path, d.Body) {
			cited[target] = true
		}
	}
	for _, d := range docs {
		if d.Kind.Layer == layerEvidence && !cited[d.Path] {
			add(d, "evidence is not cited by any document; cite it or link to it from a project")
		}
	}

	for _, extra := range []string{"README.md", "AGENTS.md"} {
		if raw, err := os.ReadFile(filepath.Join(root, extra)); err == nil {
			problems = append(problems, checkLinks(root, extra, string(raw))...)
		}
	}

	want, err := renderIndex(docs)
	if err != nil {
		return nil, err
	}
	have, _ := os.ReadFile(filepath.Join(root, "INDEX.md"))
	if string(have) != want {
		problems = append(problems, "INDEX.md: out of date; run make index")
	}
	wantData, err := renderData(docs)
	if err != nil {
		return nil, err
	}
	haveData, _ := os.ReadFile(filepath.Join(root, filepath.FromSlash(dataPath)))
	if string(haveData) != wantData {
		problems = append(problems, dataPath+": out of date; run make index")
	}
	return problems, nil
}

// linkTargets returns the repo-relative paths of local links in a body.
func linkTargets(from, body string) []string {
	body = fencedCodeBlock.ReplaceAllString(body, "")
	var out []string
	for _, m := range mdLink.FindAllStringSubmatch(body, -1) {
		target := m[1]
		if strings.Contains(target, "://") || strings.HasPrefix(target, "#") || strings.HasPrefix(target, "mailto:") {
			continue
		}
		target, _, _ = strings.Cut(target, "#")
		out = append(out, path.Clean(path.Join(path.Dir(from), target)))
	}
	return out
}

func checkLinks(root, from, body string) []string {
	var problems []string
	for _, target := range linkTargets(from, body) {
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(target))); err != nil {
			problems = append(problems, fmt.Sprintf("%s: link to %s does not resolve", from, target))
		}
	}
	return problems
}
