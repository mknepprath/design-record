---
kind: adr
number: 0014
title: Add Cleveland Cinemas through its box office API
status: accepted
owner: Michael Knepprath
date: 2026-10-17
projects: [showtime-window]
evidence:
  - evidence/observations/2026-10-16-cleveland-cinemas-schedule-api.md
  - evidence/observations/2026-10-17-cleveland-cinemas-api-boundaries.md
  - evidence/observations/2026-10-17-showtime-window-cleveland-cinemas-and-travel-limit-test-run.md
  - evidence/feedback/2026-10-17-owner-wants-as-many-theaters-as-possible.md
---

# ADR 0014: Add Cleveland Cinemas through its box office API

## Context

Cleveland Cinemas ranked second in [Study 0002](../studies/0002-which-cleveland-area-theaters-can-be-read-for-free.md) (35 of 45). The owner asked to build it and wants as many theaters as possible ([feedback](../evidence/feedback/2026-10-17-owner-wants-as-many-theaters-as-possible.md)). The site loads showtimes from its own unauthenticated JSON endpoints ([observation](../evidence/observations/2026-10-16-cleveland-cinemas-schedule-api.md)).

## Decision

- A third adapter, `lib/boxoffice.ts`, reads the `schedule` and `movies` endpoints of a Webedia theater site. Each theater stores the site origin and its theater code.
- Ranges run from 3 AM to 3 AM, the only boundary the API accepts, and movies are requested with repeated `ids=` ([boundaries](../evidence/observations/2026-10-17-cleveland-cinemas-api-boundaries.md)).
- Two weeks are fetched at a time, starting on a Monday, and cached for an hour. That is two requests per theater per hour at most. A failed request is retried once, 1.5 seconds later.
- Local times are converted to UTC in the theater's zone. Days are then assigned with the app's own 4 AM cutoff, as for every other source.
- Both theaters in the site's list are added: Cedar Lee and the Apollo in Oberlin.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| The site's JSON API | Structured; runtimes and ticket links; one adapter for any Webedia site | Undocumented, so it can change without notice; 500s in bursts |
| Scrape the rendered page | Follows what visitors see | Showtimes aren't in the HTML; needs a browser |
| Skip it | Less to maintain | Cedar Lee is the area's main art house |

## Consequences

- No sold-out status: occupancy is always `null`.
- No format is shown: the only format tag seen is the default `Digital`.
- "12 Hours of Terror" is listed with its real 12-hour runtime, so it fits only a 12-hour window.
- The Capitol Theatre isn't on this site and is still not covered.
