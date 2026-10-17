---
kind: adr
number: 0012
title: Add Atlas Cinemas via Veezi and the nearby Cinemarks
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [showtime-window]
evidence:
  - evidence/observations/2026-10-16-atlas-cinemas-veezi-sessions-data.md
  - evidence/observations/2026-10-16-cleveland-area-theater-access.md
  - evidence/observations/2026-10-16-showtime-window-veezi-conformance-run.md
  - evidence/feedback/2026-10-17-owner-wants-as-many-theaters-as-possible.md
---

# ADR 0012: Add Atlas Cinemas via Veezi and the nearby Cinemarks

Accepted 2026-10-17: the owner keeps all eight and wants as many theaters as possible ([feedback](../evidence/feedback/2026-10-17-owner-wants-as-many-theaters-as-possible.md)).

## Context

[Study 0002](../studies/0002-which-cleveland-area-theaters-can-be-read-for-free.md) ranked Atlas Cinemas, read through Veezi sessions pages, as the best non-Cinemark source near Valley View (40 of 45). Two more Cinemark theaters near Valley View use the existing adapter ([access probe](../evidence/observations/2026-10-16-cleveland-area-theater-access.md)). The project's scope was one theater.

## Decision

- A second adapter, `lib/veezi.ts`, reads a Veezi sessions page by `siteToken` ([Atlas observation](../evidence/observations/2026-10-16-atlas-cinemas-veezi-sessions-data.md)). One fetch lists every day, so it's cached per theater for an hour, and each day is filtered from it.
- Shows starting before 4 AM belong to the previous day, for every source.
- The picker lists eight theaters: Cinemark Valley View, Strongsville and Cuyahoga Falls, and the five Atlas theaters. Valley View stays the default.
- Travel minutes are saved per theater, because the theaters are different distances away.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| Atlas via Veezi plus nearby Cinemarks | Highest score; one adapter covers any Veezi theater; Cinemarks cost nothing | More theaters in the picker than the owner may use |
| Cleveland Cinemas first | Cedar Lee is an art house | Undocumented internal API; scored 35 |
| Stay at one theater | Matches the original scope | Phase 5 not started |

## Consequences

- The shortlist covered one theater at a time; [ADR 0013](0013-one-shortlist-across-theaters.md) replaced that with one list across theaters.
- A Veezi failure returns a 502 for that theater only; other theaters are unaffected.
- T5 passed for Veezi from Vercel ([conformance run](../evidence/observations/2026-10-16-showtime-window-veezi-conformance-run.md)).
- Veezi pages give no format (IMAX, 3D), so Atlas showings show none.
