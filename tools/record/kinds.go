package main

type layer string

const (
	layerProject   layer = "project"
	layerEvidence  layer = "evidence"
	layerSynthesis layer = "synthesis"
	layerDecision  layer = "decision"
)

type kind struct {
	Name     string
	Dir      string
	Layer    layer
	Statuses []string // empty: the kind has no status
	Numbered bool     // NNNN-slug.md, with a number field
	Dated    bool     // YYYY-MM-DD-slug.md, filename date matches the date field
	// Every document of this kind must cite at least one evidence record.
	NeedsEvidence bool
	NeedsSource   bool
	NeedsOwner    bool
	// Every document of this kind must name at least one project.
	NeedsProject bool
	Purpose      string
}

var kinds = []kind{
	{Name: "project", Dir: "projects", Layer: layerProject,
		Statuses: []string{"idea", "active", "paused", "done", "dropped"}, NeedsOwner: true,
		Purpose: "The entry point. What is being built and why; everything else backs a project."},

	{Name: "interview", Dir: "evidence/interviews", Layer: layerEvidence, Dated: true, NeedsSource: true, NeedsProject: true,
		Purpose: "A conversation with a person, as heard. Append-only."},
	{Name: "observation", Dir: "evidence/observations", Layer: layerEvidence, Dated: true, NeedsSource: true, NeedsProject: true,
		Purpose: "Something seen or measured: command output, a page, a log. Append-only."},
	{Name: "feedback", Dir: "evidence/feedback", Layer: layerEvidence, Dated: true, NeedsSource: true, NeedsProject: true,
		Purpose: "Something a person said unprompted: a message, a comment, a request. Append-only."},
	{Name: "meeting", Dir: "evidence/meetings", Layer: layerEvidence, Dated: true, NeedsSource: true, NeedsProject: true,
		Purpose: "Notes from a meeting. Append-only."},

	{Name: "persona", Dir: "personas", Layer: layerSynthesis, NeedsEvidence: true,
		Purpose: "A kind of user, built from evidence."},
	{Name: "journey", Dir: "journeys", Layer: layerSynthesis, NeedsEvidence: true, NeedsProject: true,
		Purpose: "The steps a persona takes to get something done, built from evidence."},
	{Name: "study", Dir: "studies", Layer: layerSynthesis, Numbered: true, NeedsEvidence: true, NeedsProject: true,
		Statuses: []string{"planned", "running", "done"}, NeedsOwner: true,
		Purpose: "A question, the method used to answer it, and what was found."},
	{Name: "glossary", Dir: "glossary", Layer: layerSynthesis,
		Purpose: "One term and what it means here."},

	{Name: "adr", Dir: "decisions", Layer: layerDecision, Numbered: true, NeedsEvidence: true, NeedsProject: true,
		Statuses: []string{"proposed", "accepted", "rejected", "superseded"}, NeedsOwner: true,
		Purpose: "One decision, its context, the options and the consequences."},
	{Name: "spec", Dir: "specs", Layer: layerDecision, Numbered: true, NeedsEvidence: true, NeedsProject: true,
		Statuses: []string{"draft", "review", "accepted", "superseded"}, NeedsOwner: true,
		Purpose: "A conformance spec: MUST/SHOULD/MAY requirements and the tests that check them."},
}

func kindByName(name string) (kind, bool) {
	for _, k := range kinds {
		if k.Name == name {
			return k, true
		}
	}
	return kind{}, false
}
