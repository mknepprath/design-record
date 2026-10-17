---
kind: study
number: 0002
title: Which Cleveland-area theaters can be read for free
status: done
owner: Michael Knepprath
date: 2026-10-16
projects: [showtime-window]
evidence:
  - evidence/observations/2026-10-16-cleveland-area-theater-access.md
  - evidence/observations/2026-10-16-atlas-cinemas-veezi-sessions-data.md
  - evidence/observations/2026-10-16-cleveland-cinemas-schedule-api.md
---

# Study 0002: Which Cleveland-area theaters can be read for free

## Question

Which theaters near Valley View, Ohio, other than Cinemark, can the [Showtime window](../projects/showtime-window.md) read at $0, and how do they score against [Spec 0001](../specs/0001-showtime-source-conformance.md)? This is phase 5 of the project, following [Study 0001](0001-which-showtime-sources-can-be-read-for-free.md), which looked at national chains.

## Method

1. Listed nearby theaters from web searches and from the theater sites' own location lists.
2. Fetched each operator's site with `curl` and recorded status and blocking ([access probe](../evidence/observations/2026-10-16-cleveland-area-theater-access.md)).
3. Where the HTML had no showtimes, loaded the page in headless Chromium, recorded its JSON requests, and repeated them with `curl` ([Cleveland Cinemas](../evidence/observations/2026-10-16-cleveland-cinemas-schedule-api.md)).
4. Inspected the fields of each usable source ([Atlas Cinemas](../evidence/observations/2026-10-16-atlas-cinemas-veezi-sessions-data.md)) and scored it with the Spec 0001 rubric.

Each source was observed once, from a home connection, on the night of 2026-10-16. Nothing was checked from Vercel (Spec 0001 T5). Drive times from Valley View weren't measured.

## Findings

| Source | Theaters | Conforms | Score /45 | Basis |
| --- | --- | --- | --- | --- |
| Atlas Cinemas via Veezi | Shaker Square 6, Barrington 10 (Aurora), Diamond Center 16 and Great Lakes Stadium 16 (Mentor), Lakeshore 7 (Euclid) | Yes | 40 | schema.org sessions with UTC offset, runtime and ticket link, 12 to 53 days in one page; no format ([observation](../evidence/observations/2026-10-16-atlas-cinemas-veezi-sessions-data.md)) |
| Cleveland Cinemas (Webedia platform API) | Cedar Lee (Cleveland Heights), Apollo (Oberlin) | Yes | 35 | Public JSON the site's own app calls: local times plus IANA zone, format tags, ticket links; runtime from a second call ([observation](../evidence/observations/2026-10-16-cleveland-cinemas-schedule-api.md)) |
| Cleveland Cinematheque (Nightjar site) | Cinematheque | Yes, with HTML for later days | 32 | schema.org `ScreeningEvent` for today only; other days only in visible HTML ([probe](../evidence/observations/2026-10-16-cleveland-area-theater-access.md)) |
| Other Cinemarks | Strongsville at Southpark Mall, Cuyahoga Falls and XD | Expected yes | – | Same adapter as Valley View; not fetched ([probe](../evidence/observations/2026-10-16-cleveland-area-theater-access.md)) |
| Regal | Willoughby Commons and others | No (R5) | – | Cloudflare 403 |
| AMC | Cleveland area | No (R5) | – | Cloudflare 403 |
| Atom Tickets | Aggregator | No (R5) | – | Cloudflare 403 |
| Fandango | Aggregator | Inconclusive | – | No showtimes in the HTML; not traced |
| IMDb showtimes | Aggregator | Inconclusive | – | 202 with no schedule |
| Capitol Theatre | Gordon Square | Not assessed | – | New operator since 2026-07-31 per a news report (secondary); its showtime site wasn't found |
| Cleveland Museum of Art | Museum film program | Not assessed | – | Guessed URL returned 404 |

Score breakdown, weights from Spec 0001:

- **Atlas via Veezi:** access 3, bot 3, completeness 2 (no format), coverage 3 (one adapter reads any Veezi theater by its siteToken), terms 2, stability 2.
- **Cleveland Cinemas:** access 2 (a JavaScript app's API, though callable directly), bot 3, completeness 3, coverage 2 (Webedia runs other theater sites; whether they share this API wasn't checked), terms 1 (undocumented internal API), stability 1 (internal API).
- **Cinematheque:** access 2, bot 3, completeness 2, coverage 1, terms 2, stability 2.

Stability is unknown for every source and gets a neutral 2, except the internal API.

## Implications

- Atlas via Veezi is the next source to build. One adapter covers five Atlas theaters, and one fetch returns weeks of sessions, so caching can be per theater, not per theater-day.
- Cleveland Cinemas is the second. Its local times need the theater's zone to make UTC, and the 3:00 AM day boundary means late shows belong to the previous date.
- The `Showing.sourceId` type and the API route assume Cinemark only. Adding a source changes them; phase 5's "one broken source doesn't break the build" applies once there are two.
- Which theaters to add is for the owner to choose; how near each is wasn't measured.

## Open

- Run Spec 0001 T5 for Veezi and Cleveland Cinemas from Vercel.
- Find the Capitol Theatre's new showtime site.
- Check whether Atlas Eastgate 10 is still open.
