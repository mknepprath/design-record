---
kind: adr
number: 0008
title: Move Memory Map out of Obsidian into a standalone web app
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [memory-map]
evidence:
  - evidence/observations/2026-10-16-claude-code-memory-notes-on-projects.md
---

# ADR 0008: Move Memory Map out of Obsidian into a standalone web app

Recorded 2026-10-16 after the fact, from the sources cited. The decision itself was made earlier, when the code was written.

## Context

Memory Map started as an Obsidian plugin, which kept freezing ([memory notes](../evidence/observations/2026-10-16-claude-code-memory-notes-on-projects.md)).

## Decision

Rebuild it as a standalone Vite and Three.js app that reads an exported `notes.json` instead of the live vault.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| Standalone app | No longer freezes Obsidian; full control of rendering | Notes must be exported first; no live updates |
| Keep the plugin | Live vault access | Froze |

## Consequences

- An export script has to run before the map reflects new notes.
- Evidence is a secondary summary; the freeze itself wasn't recorded in detail.
