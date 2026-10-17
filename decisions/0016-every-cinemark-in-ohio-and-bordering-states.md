---
kind: adr
number: 0016
title: Every Cinemark in Ohio and bordering states, generated from the sitemap
status: accepted
owner: Michael Knepprath
date: 2026-10-17
projects: [showtime-window]
evidence:
  - evidence/feedback/2026-10-17-owner-wants-as-many-theaters-as-possible.md
  - evidence/observations/2026-10-17-cinemark-sitemap-catalog.md
  - evidence/observations/2026-10-17-cinemark-rate-limits-vercel.md
---

# ADR 0016: Every Cinemark in Ohio and bordering states, generated from the sitemap

## Context

The owner wants as many theaters as possible ([feedback](../evidence/feedback/2026-10-17-owner-wants-as-many-theaters-as-possible.md)). The Cinemark adapter already reads any Cinemark theater, and the sitemap lists all of them ([catalog](../evidence/observations/2026-10-17-cinemark-sitemap-catalog.md)). [ADR 0015](0015-include-theaters-by-travel-time.md) keeps far theaters out of the list.

## Decision

- `scripts/cinemark-catalog.ts` builds `lib/theaters-cinemark.ts` from the sitemap and each theater page, one request a second. Each theater's timezone comes from its showtimes' UTC offset. A page without a `MovieTheater` node or showtimes is skipped.
- It runs for Ohio, Pennsylvania, Michigan, Indiana, Kentucky and West Virginia: 46 theaters, 53 in all with Atlas and Cleveland Cinemas.
- Travel estimates route only theaters whose straight-line guess is within 240 minutes, 50 per OSRM request. The rest keep the guess.
- One list covers at most the 20 nearest theaters in range. The page keeps at most 3 requests in flight and never repeats one in progress ([rate limits](../evidence/observations/2026-10-17-cinemark-rate-limits-vercel.md)).

## Options considered

| Option | For | Against |
| --- | --- | --- |
| Ohio and bordering states | Covers any drive from northeast Ohio; Eastern and Central only | Still regional |
| All 306 US theaters | Most coverage | Mountain and Pacific offsets, Arizona without DST; nobody drives from Cleveland to Texas |
| Hand-pick theaters | Smallest list | Doesn't scale |

## Consequences

- Cinemark rate-limits Vercel. A wide limit, or many visitors on a cold cache, can get 429s for some theaters, and the page names them as not loaded.
- The live conformance run can no longer cover every Cinemark theater from Vercel in one go. It takes `ONLY=` to run a subset.
- The offset-to-zone mapping assumes US zones that observe DST. Arizona and Hawaii need handling before those states are added.
- New or closed theaters appear only when the script is rerun.
