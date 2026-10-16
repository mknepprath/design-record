---
kind: adr
number: 0006
title: Keep Pokémon GO event times as naive local times
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [lily-dex]
evidence:
  - evidence/observations/2026-10-16-lily-dex-api-readme.md
---

# ADR 0006: Keep Pokémon GO event times as naive local times

Recorded 2026-10-16 after the fact, from the sources cited. The decision itself was made earlier, when the code was written.

## Context

Pokémon GO events happen at the same local time in every timezone ([README](../evidence/observations/2026-10-16-lily-dex-api-readme.md)).

## Decision

Event timestamps in `events.json` carry no timezone. The app reads them as the user's local time.

## Options considered

Alternatives weren't recorded.

## Consequences

- A UTC conversion anywhere in the pipeline would shift every event by the user's offset.
- Events that are global (one moment worldwide) can't be represented exactly with this format.
