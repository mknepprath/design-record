---
kind: adr
number: 0009
title: Back Duck Pond cards with Obsidian notes
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [duck-pond]
evidence:
  - evidence/observations/2026-10-16-claude-code-memory-notes-on-projects.md
---

# ADR 0009: Back Duck Pond cards with Obsidian notes

Recorded 2026-10-16 after the fact, from the sources cited. The decision itself was made earlier, when the code was written.

## Context

Duck Pond's dispatched prompts tell an agent to update the card when it's done, so the card's state must be something an agent can edit ([memory notes](../evidence/observations/2026-10-16-claude-code-memory-notes-on-projects.md)).

## Decision

Each card is a Markdown note in the vault, with status in YAML frontmatter. The app watches the folder and syncs live. It's a local-first native macOS app.

## Options considered

Alternatives weren't recorded.

## Consequences

- Any agent with file access can move a card by editing frontmatter.
- The app can't be sandboxed, since it needs access to an arbitrary folder.
