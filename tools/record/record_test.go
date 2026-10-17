package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func write(t *testing.T, root, rel, content string) {
	t.Helper()
	p := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

const project = "---\nkind: project\ntitle: P\nstatus: active\nowner: O\ndate: 2026-10-16\n---\n# P\n"
const evidence = "---\nkind: observation\ntitle: E\ndate: 2026-10-16\nsource: s\nprojects: [p]\n---\n# E\n"

func validRepo(t *testing.T) string {
	root := t.TempDir()
	write(t, root, "projects/p.md", project)
	write(t, root, "evidence/observations/2026-10-16-e.md", evidence)
	write(t, root, "decisions/0001-d.md", "---\nkind: adr\nnumber: 0001\ntitle: D\nstatus: proposed\nowner: O\ndate: 2026-10-16\nprojects: [p]\nevidence:\n  - evidence/observations/2026-10-16-e.md\n---\n# D\n\nSee [P](../projects/p.md).\n")
	docs, err := loadDocs(root)
	if err != nil {
		t.Fatal(err)
	}
	idx, _ := renderIndex(docs)
	write(t, root, "INDEX.md", idx)
	data, _ := renderData(docs)
	write(t, root, dataPath, data)
	return root
}

func problemsFor(t *testing.T, root string) string {
	t.Helper()
	p, err := check(root)
	if err != nil {
		t.Fatal(err)
	}
	return strings.Join(p, "\n")
}

func TestValidRepoPasses(t *testing.T) {
	if got := problemsFor(t, validRepo(t)); got != "" {
		t.Fatalf("unexpected problems:\n%s", got)
	}
}

func TestRules(t *testing.T) {
	cases := []struct {
		name, rel, content, want string
	}{
		{"wrong kind", "decisions/0002-x.md", "---\nkind: spec\nnumber: 0002\ntitle: X\nstatus: draft\nowner: O\ndate: 2026-10-16\nprojects: [p]\nevidence: [evidence/observations/2026-10-16-e.md]\n---\n", `must be kind "adr"`},
		{"bad status", "decisions/0002-x.md", "---\nkind: adr\nnumber: 0002\ntitle: X\nstatus: maybe\nowner: O\ndate: 2026-10-16\nprojects: [p]\nevidence: [evidence/observations/2026-10-16-e.md]\n---\n", `status "maybe"`},
		{"duplicate number", "decisions/0001-y.md", "---\nkind: adr\nnumber: 0001\ntitle: Y\nstatus: proposed\nowner: O\ndate: 2026-10-16\nprojects: [p]\nevidence: [evidence/observations/2026-10-16-e.md]\n---\n", "also used by"},
		{"no evidence", "decisions/0002-x.md", "---\nkind: adr\nnumber: 0002\ntitle: X\nstatus: proposed\nowner: O\ndate: 2026-10-16\nprojects: [p]\nevidence: []\n---\n", "must cite at least one evidence"},
		{"unknown project", "decisions/0002-x.md", "---\nkind: adr\nnumber: 0002\ntitle: X\nstatus: proposed\nowner: O\ndate: 2026-10-16\nprojects: [nope]\nevidence: [evidence/observations/2026-10-16-e.md]\n---\n", `project "nope" does not exist`},
		{"evidence date mismatch", "evidence/feedback/2026-10-15-f.md", "---\nkind: feedback\ntitle: F\ndate: 2026-10-16\nsource: s\nprojects: [p]\n---\n[P](../../projects/p.md)\n", "does not match date field"},
		{"uncited evidence", "evidence/feedback/2026-10-16-f.md", "---\nkind: feedback\ntitle: F\ndate: 2026-10-16\nsource: s\nprojects: [p]\n---\n", "not cited by any document"},
		{"broken link", "glossary/term.md", "---\nkind: glossary\ntitle: T\ndate: 2026-10-16\n---\n[x](missing.md)\n", "does not resolve"},
		{"no frontmatter", "glossary/term.md", "# T\n", "missing frontmatter"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			root := validRepo(t)
			write(t, root, c.rel, c.content)
			if got := problemsFor(t, root); !strings.Contains(got, c.want) {
				t.Fatalf("want a problem containing %q, got:\n%s", c.want, got)
			}
		})
	}
}

func TestStaleIndex(t *testing.T) {
	root := validRepo(t)
	write(t, root, "glossary/term.md", "---\nkind: glossary\ntitle: T\ndate: 2026-10-16\n---\n")
	if got := problemsFor(t, root); !strings.Contains(got, "INDEX.md: out of date") {
		t.Fatalf("stale index not reported:\n%s", got)
	}
}

func TestNewNumbersAndFills(t *testing.T) {
	root := validRepo(t)
	write(t, root, "templates/adr.md", "---\nkind: adr\nnumber: {{number}}\ntitle: {{title}}\nprojects: {{projects}}\n---\n")
	rel, err := newDoc(root, "adr", "Use Go", "p", "")
	if err != nil {
		t.Fatal(err)
	}
	if rel != "decisions/0002-use-go.md" {
		t.Fatalf("path = %s", rel)
	}
	raw, _ := os.ReadFile(filepath.Join(root, rel))
	if !strings.Contains(string(raw), "number: 0002") || !strings.Contains(string(raw), "projects: [p]") {
		t.Fatalf("template not filled:\n%s", raw)
	}
	if _, err := newDoc(root, "adr", "X", "missing", ""); err == nil {
		t.Fatal("expected error for unknown project")
	}
}

func TestParseFrontmatterLists(t *testing.T) {
	f, l, body, ok := parseFrontmatter("---\ntitle: \"A: b\"\nprojects: [x, y]\nevidence:\n  - one.md\n  - two.md\n---\nbody\n")
	if !ok || f["title"] != "A: b" || len(l["projects"]) != 2 || len(l["evidence"]) != 2 || body != "body\n" {
		t.Fatalf("got %v %v %q %v", f, l, body, ok)
	}
}

func TestBacklinks(t *testing.T) {
	root := validRepo(t)
	docs, _ := loadDocs(root)
	data, _ := renderData(docs)
	if !strings.Contains(data, `"backlinks": [
      "decisions/0001-d.md"
    ]`) {
		t.Fatalf("project should be linked from the ADR:\n%s", data)
	}
}
