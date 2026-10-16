---
kind: adr
number: 0010
title: Give each content type on mknepprath.com its own physical object
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [mknepprath-next]
evidence:
  - evidence/observations/2026-10-16-claude-code-memory-notes-on-projects.md
---

# ADR 0010: Give each content type on mknepprath.com its own physical object

Recorded 2026-10-16 after the fact, from the sources cited. The decision itself was made earlier, when the code was written.

## Context

The activity feed mixes runs, films, commits, music, posts, books and photos ([memory notes](../evidence/observations/2026-10-16-claude-code-memory-notes-on-projects.md)).

## Decision

Draw each content type as the object it is (a movie ticket for a film, a Polaroid for a photo), not as one card in different colors. They're tied together by a consistent stroke weight and slight random rotation.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| One object per type | Content is recognizable at a glance | More components to build and maintain |
| One card, recolored | Less code | Generic |

## Consequences

- Each new content type needs its own design.
- Rule from the notes: "if a detail needs explaining, cut it".
