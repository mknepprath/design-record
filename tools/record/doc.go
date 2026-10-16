package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type doc struct {
	Path     string // repo-relative, slash-separated
	Kind     kind
	Fields   map[string]string
	Lists    map[string][]string
	Body     string
	HasFront bool
}

func (d doc) Title() string      { return d.Fields["title"] }
func (d doc) Status() string     { return d.Fields["status"] }
func (d doc) Date() string       { return d.Fields["date"] }
func (d doc) Projects() []string { return d.Lists["projects"] }
func (d doc) Evidence() []string { return d.Lists["evidence"] }
func (d doc) Slug() string       { return strings.TrimSuffix(filepath.Base(d.Path), ".md") }

func (d doc) Number() int {
	n, _ := strconv.Atoi(d.Fields["number"])
	return n
}

// Ref is the short label used in the index: "ADR 0002", "Spec 0001".
func (d doc) Ref() string {
	if !d.Kind.Numbered {
		return ""
	}
	label := map[string]string{"adr": "ADR", "spec": "Spec", "study": "Study"}[d.Kind.Name]
	return fmt.Sprintf("%s %04d", label, d.Number())
}

// parseFrontmatter reads a YAML subset: `key: value`, `key: [a, b]` and
// block lists (`key:` followed by `  - item` lines). Enough for these
// documents without a dependency.
func parseFrontmatter(src string) (fields map[string]string, lists map[string][]string, body string, ok bool) {
	fields, lists = map[string]string{}, map[string][]string{}
	if !strings.HasPrefix(src, "---\n") {
		return fields, lists, src, false
	}
	end := strings.Index(src[4:], "\n---\n")
	if end < 0 {
		return fields, lists, src, false
	}
	front, body := src[4:4+end], src[4+end+5:]
	var listKey string
	sc := bufio.NewScanner(strings.NewReader(front))
	for sc.Scan() {
		line := sc.Text()
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if strings.HasPrefix(trimmed, "- ") && listKey != "" {
			lists[listKey] = append(lists[listKey], unquote(strings.TrimSpace(trimmed[2:])))
			continue
		}
		key, val, found := strings.Cut(line, ":")
		if !found {
			continue
		}
		key, val = strings.TrimSpace(key), strings.TrimSpace(val)
		listKey = ""
		switch {
		case val == "":
			listKey = key
			lists[key] = nil
		case strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]"):
			inner := strings.TrimSpace(val[1 : len(val)-1])
			lists[key] = nil
			if inner != "" {
				for _, item := range strings.Split(inner, ",") {
					lists[key] = append(lists[key], unquote(strings.TrimSpace(item)))
				}
			}
		default:
			fields[key] = unquote(val)
		}
	}
	return fields, lists, body, true
}

func unquote(s string) string {
	if len(s) >= 2 && (s[0] == '"' && s[len(s)-1] == '"' || s[0] == '\'' && s[len(s)-1] == '\'') {
		return s[1 : len(s)-1]
	}
	return s
}

// loadDocs reads every markdown file under the kind directories.
func loadDocs(root string) ([]doc, error) {
	var docs []doc
	for _, k := range kinds {
		dir := filepath.Join(root, filepath.FromSlash(k.Dir))
		err := filepath.WalkDir(dir, func(path string, e fs.DirEntry, err error) error {
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}
				return err
			}
			if e.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, ".md") || strings.EqualFold(e.Name(), "README.md") {
				return nil
			}
			raw, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			rel, _ := filepath.Rel(root, path)
			f, l, body, ok := parseFrontmatter(strings.ReplaceAll(string(raw), "\r\n", "\n"))
			docs = append(docs, doc{Path: filepath.ToSlash(rel), Kind: k, Fields: f, Lists: l, Body: body, HasFront: ok})
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.Slice(docs, func(i, j int) bool { return docs[i].Path < docs[j].Path })
	return docs, nil
}

var nonSlug = regexp.MustCompile(`[^a-z0-9]+`)

func slugify(title string) string {
	s := nonSlug.ReplaceAllString(strings.ToLower(title), "-")
	s = strings.Trim(s, "-")
	if len(s) > 60 {
		s = strings.TrimRight(s[:60], "-")
	}
	return s
}
