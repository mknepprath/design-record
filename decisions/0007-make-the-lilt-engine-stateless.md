---
kind: adr
number: 0007
title: Make the lilt engine stateless, with state held by the client
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [lilt]
evidence:
  - evidence/observations/2026-10-16-lilt-readme.md
---

# ADR 0007: Make the lilt engine stateless, with state held by the client

Recorded 2026-10-16 after the fact, from the sources cited. The decision itself was made earlier, when the code was written.

## Context

lilt is played through several clients: a web page, Mastodon and a CLI ([README](../evidence/observations/2026-10-16-lilt-readme.md)).

## Decision

The engine is a pure function, `play(move, state) → {response, state}`, behind a Flask API on Railway. Each client stores the game state and sends it with every move. The Mastodon adapter persists state per user.

## Options considered

Alternatives weren't recorded.

## Consequences

- The API needs no database, and any client can be added without server changes.
- A client can send altered state; there's no server-side check.
