---
title: About
---

# design-record

Michael Knepprath's design record across personal projects: what was heard and seen, what was concluded from it, what was decided, and what is being built. Projects are the entry point; everything else backs a project and links to it.

## What's here

[INDEX.md](INDEX.md) lists every document, and under each project, the documents that belong to it. The documents come in three layers:

- **Evidence:** interviews, observations, feedback and meeting notes, in `evidence/`. Dated and append-only.
- **Synthesis:** personas, journeys, glossary terms and studies. Studies, personas and journeys cite their evidence.
- **Decisions and specs:** ADRs in `decisions/` and conformance specs in `specs/`. Each moves through its statuses by commit.

## Adding or changing something

Needs Go 1.23 (standard library only, no network) and `make`.

```sh
make hooks                                   # once per clone: make check runs before every push
make new                                     # lists the kinds and where each goes
make new KIND=observation TITLE="Cinemark pages load from CI" \
         PROJECT=showtime-window SOURCE="GitHub Actions run, 2026-10-17"
make index                                   # regenerate INDEX.md
make check                                   # the rules, checked
```

`make new` picks the path, the number and today's date, and starts from the kind's template in `templates/`. [AGENTS.md](AGENTS.md) has the rest: which kind to use, how each kind changes, and what `make check` verifies.
