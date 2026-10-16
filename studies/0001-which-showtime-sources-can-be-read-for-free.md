---
kind: study
number: 0001
title: Which showtime sources can be read for free
status: done
owner: Michael Knepprath
date: 2026-10-16
projects: [showtime-window]
evidence:
  - evidence/observations/2026-10-16-theater-site-access-from-a-script.md
  - evidence/observations/2026-10-16-cinemark-screeningevent-data.md
  - evidence/observations/2026-10-16-alamo-drafthouse-schedule-json.md
  - evidence/observations/2026-10-16-free-showtime-data-options.md
---

# Study 0001: Which showtime sources can be read for free

## Question

Which US theater chains or showtime services can a scheduled script read at $0 and score well against [Spec 0001](../specs/0001-showtime-source-conformance.md)?

## Method

1. Fetched 11 chain sites with `curl` and recorded status, size and blocking vendor ([access probe](../evidence/observations/2026-10-16-theater-site-access-from-a-script.md)).
2. For the two that returned usable data, inspected the fields ([Cinemark](../evidence/observations/2026-10-16-cinemark-screeningevent-data.md), [Alamo](../evidence/observations/2026-10-16-alamo-drafthouse-schedule-json.md)).
3. Searched for free tiers of commercial feeds ([search results](../evidence/observations/2026-10-16-free-showtime-data-options.md)).
4. Scored each conforming source with the Spec 0001 rubric.

Each source was observed once, on one day, from a home connection.

## Findings

| Source | Conforms | Score /45 | Basis |
| --- | --- | --- | --- |
| Cinemark | Yes | 43 | Structured `ScreeningEvent` data with UTC offset, duration, format, availability and ticket URL; `?showDate=` selects the day ([observation](../evidence/observations/2026-10-16-cinemark-screeningevent-data.md)) |
| Alamo Drafthouse | Yes | 40 | Public JSON with 53 days of sessions, UTC times and sold-out status; no runtime ([observation](../evidence/observations/2026-10-16-alamo-drafthouse-schedule-json.md)) |
| AMC | No (R5) | – | Cloudflare 403 ([probe](../evidence/observations/2026-10-16-theater-site-access-from-a-script.md)) |
| Regal | No (R5) | – | Cloudflare 403 |
| Marcus | No (R5) | – | Imperva/Incapsula challenge |
| Studio Movie Grill | No (R5) | – | 403 |
| Fandango, Landmark, Showcase, Angelika | Inconclusive | – | Pages load, but showtime data was not found or traced |
| Harkins | Not assessed | – | Guessed URL was wrong |
| Paid feeds (Gracenote, MovieGlu, International Showtimes) | Out of budget | – | No confirmed free tier ([search results](../evidence/observations/2026-10-16-free-showtime-data-options.md)) |

Score breakdown, weights from Spec 0001:

- **Cinemark:** access 3, bot 3, completeness 3, coverage 3, terms 2, stability 2. Stability is unknown and gets a neutral 2.
- **Alamo:** access 3, bot 3, completeness 2, coverage 3, terms 2, stability 2.

## Open

- Probe Fandango, Landmark, Showcase, Angelika and Harkins theater pages and trace their data requests.
- Repeat the Cinemark probe on another day and from CI before treating it as stable.
