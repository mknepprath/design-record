# AGENTS.md

Instructions for anyone, human or agent, adding to this record.

## Where things go

Run `make new` for the live list. In short:

| You have | Kind | Goes in |
| --- | --- | --- |
| Something being built, and why | project | `projects/` |
| A conversation, as heard | interview | `evidence/interviews/` |
| Something seen or measured: command output, a page, a log | observation | `evidence/observations/` |
| Something a person said or asked for, unprompted | feedback | `evidence/feedback/` |
| Meeting notes | meeting | `evidence/meetings/` |
| A question, how it was answered, what was found | study | `studies/` |
| A kind of user | persona | `personas/` |
| The steps a user takes, and where they break | journey | `journeys/` |
| One term and what it means here | glossary | `glossary/` |
| One decision, with options and consequences | adr | `decisions/` |
| Requirements a thing must meet, and the tests for them | spec | `specs/` |

Rule of thumb: if it happened, it's evidence. If it interprets evidence, it's synthesis. If it commits to something, it's a decision or a spec.

## Writing rules

- Evidence records what was heard or seen, quoted or pasted where possible. No interpretation.
- Every claim in synthesis, ADRs and specs links the evidence it rests on. Mark anything that comes from recall rather than evidence: "(recalled, not checked)".
- State evidence strength where it matters: single observation, repeated, independently corroborated, primary document.
- This repo is public. Material from private repos, or about other people, is not copied in: link to it under a project's **Private documents** heading, with a line saying only collaborators can follow the links. Anything else a design needs is written here, not linked.
- Name people only with their consent. Refer to collaborators by role otherwise.

## Changing things

| Kind | How it changes |
| --- | --- |
| Evidence | Append-only. Add a dated section; never rewrite what was recorded. A correction is a new record. |
| Project | Edit freely; keep the Status section current. Statuses: idea → active → paused / done / dropped. |
| Study | planned → running → done. Once done, findings change only through a new study. |
| ADR | proposed → accepted / rejected. An accepted ADR is not edited; a new ADR supersedes it (`supersedes: [decisions/NNNN-….md]`) and the old one's status becomes `superseded`. |
| Spec | draft → review → accepted → superseded. Requirement IDs (R1, T1…) are never reused. |
| Persona, journey, glossary | Edit freely; keep citations current. |

## Frontmatter

```yaml
---
kind: adr                 # must match the directory
number: 0003              # numbered kinds: adr, spec, study; matches the file name
title: Cinemark as the first showtime source
status: accepted          # kinds with statuses only
owner: Michael Knepprath  # project, adr, spec, study
date: 2026-10-16          # YYYY-MM-DD; evidence file names start with it
source: "..."             # evidence only: who or what, and when
projects: [showtime-window]
evidence:                 # adr, spec, study, persona, journey: at least one
  - evidence/observations/2026-10-16-cinemark-screeningevent-data.md
---
```

## What `make check` verifies

- Every document has frontmatter, its `kind` matches its directory, and it has a title and a valid date.
- Status is one the kind allows; owner and source are present where required.
- File names: `NNNN-slug.md` for numbered kinds (number matches, no duplicates), `YYYY-MM-DD-slug.md` for evidence (date matches), `slug.md` otherwise.
- Every named project exists; every cited evidence path exists and is evidence.
- ADRs, specs, studies, personas and journeys cite at least one evidence record.
- Every evidence record is cited or linked by some document.
- Every relative link in every document, README.md and AGENTS.md resolves.
- INDEX.md matches what `make index` would generate.

Not verified: append-only evidence, and that an accepted ADR is unchanged. Those are conventions.
