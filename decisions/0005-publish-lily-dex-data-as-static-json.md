---
kind: adr
number: 0005
title: Publish lily dex data as static JSON on GitHub Pages every 6 hours
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [lily-dex]
evidence:
  - evidence/observations/2026-10-16-lily-dex-api-readme.md
---

# ADR 0005: Publish lily dex data as static JSON on GitHub Pages every 6 hours

Recorded 2026-10-16 after the fact, from the sources cited. The decision itself was made earlier, when the code was written.

## Context

The iOS app needs Pokémon GO data merged from several upstream sources ([README](../evidence/observations/2026-10-16-lily-dex-api-readme.md)).

## Decision

A GitHub Actions build runs on push, every 6 hours and on demand, writes JSON files and deploys them to GitHub Pages. Source fetches are cached in `cache/` and committed, so a build can fall back when an upstream is down. No runtime dependencies.

## Options considered

Alternatives weren't recorded.

## Consequences

- Data can be up to 6 hours old.
- No server to run or pay for.
- The repo must stay public for Pages on the current plan (owner's notes).
- The same pattern was proposed for Showtime window in [ADR 0002](0002-prebuild-showtimes-as-static-json.md).
