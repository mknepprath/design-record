// Command record validates, indexes and scaffolds the design record.
//
//	record new                       list the kinds and where each goes
//	record new KIND TITLE [PROJECT] [SOURCE]
//	record index                     regenerate INDEX.md
//	record check                     run the rules; non-zero exit on any problem
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	root, err := os.Getwd()
	if err != nil {
		fail(err)
	}
	switch os.Args[1] {
	case "check":
		problems, err := check(root)
		if err != nil {
			fail(err)
		}
		for _, p := range problems {
			fmt.Fprintln(os.Stderr, p)
		}
		if len(problems) > 0 {
			fmt.Fprintf(os.Stderr, "\n%d problem(s)\n", len(problems))
			os.Exit(1)
		}
		fmt.Println("ok")
	case "index":
		docs, err := loadDocs(root)
		if err != nil {
			fail(err)
		}
		out, err := renderIndex(docs)
		if err != nil {
			fail(err)
		}
		if err := os.WriteFile(filepath.Join(root, "INDEX.md"), []byte(out), 0o644); err != nil {
			fail(err)
		}
		fmt.Printf("INDEX.md: %d documents\n", len(docs))
	case "new":
		args := os.Args[2:]
		if len(args) == 0 || args[0] == "" {
			listKinds()
			return
		}
		get := func(i int) string {
			if i < len(args) {
				return args[i]
			}
			return ""
		}
		p, err := newDoc(root, get(0), get(1), get(2), get(3))
		if err != nil {
			fail(err)
		}
		fmt.Println(p)
	default:
		usage()
	}
}

func listKinds() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "KIND\tGOES IN\tLAYER\tWHAT IT IS")
	for _, k := range kinds {
		fmt.Fprintf(w, "%s\t%s/\t%s\t%s\n", k.Name, k.Dir, k.Layer, k.Purpose)
	}
	w.Flush()
	fmt.Println("\nmake new KIND=<kind> TITLE=\"...\" [PROJECT=<project-slug>] [SOURCE=\"who/what, YYYY-MM-DD\"]")
}

func newDoc(root, kindName, title, project, source string) (string, error) {
	k, ok := kindByName(kindName)
	if !ok {
		return "", fmt.Errorf("unknown kind %q; run make new to list kinds", kindName)
	}
	if title == "" {
		return "", fmt.Errorf("TITLE is required")
	}
	if k.NeedsProject && project == "" {
		return "", fmt.Errorf("PROJECT is required for %s", k.Name)
	}
	if project != "" {
		if _, err := os.Stat(filepath.Join(root, "projects", project+".md")); err != nil {
			return "", fmt.Errorf("project %q does not exist (expected projects/%s.md)", project, project)
		}
	}
	if k.NeedsSource && source == "" {
		return "", fmt.Errorf("SOURCE is required for %s", k.Name)
	}

	date := time.Now().Format("2006-01-02")
	slug := slugify(title)
	number := ""
	var name string
	switch {
	case k.Numbered:
		docs, err := loadDocs(root)
		if err != nil {
			return "", err
		}
		next := 1
		for _, d := range docs {
			if d.Kind.Name == k.Name && d.Number() >= next {
				next = d.Number() + 1
			}
		}
		number = fmt.Sprintf("%04d", next)
		name = number + "-" + slug + ".md"
	case k.Dated:
		name = date + "-" + slug + ".md"
	default:
		name = slug + ".md"
	}
	rel := filepath.ToSlash(filepath.Join(k.Dir, name))
	dst := filepath.Join(root, filepath.FromSlash(rel))
	if _, err := os.Stat(dst); err == nil {
		return "", fmt.Errorf("%s already exists", rel)
	}

	tmpl, err := os.ReadFile(filepath.Join(root, "templates", k.Name+".md"))
	if err != nil {
		return "", fmt.Errorf("no template for %s: %w", k.Name, err)
	}
	projects := "[]"
	if project != "" {
		projects = "[" + project + "]"
	}
	out := strings.NewReplacer(
		"{{title}}", title,
		"{{date}}", date,
		"{{number}}", number,
		"{{projects}}", projects,
		"{{source}}", source,
		"{{owner}}", owner(),
	).Replace(string(tmpl))
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return "", err
	}
	return rel, os.WriteFile(dst, []byte(out), 0o644)
}

func owner() string {
	if o := os.Getenv("OWNER"); o != "" {
		return o
	}
	if out, err := exec.Command("git", "config", "user.name").Output(); err == nil {
		if name := strings.TrimSpace(string(out)); name != "" {
			return name
		}
	}
	return "unknown"
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: record new|index|check")
	os.Exit(2)
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
