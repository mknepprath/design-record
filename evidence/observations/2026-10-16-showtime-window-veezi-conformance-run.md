---
kind: observation
title: Showtime window's Veezi conformance run
date: 2026-10-16
source: npm test, npm run conformance:live and npm run e2e in showtime-window at commits 68c751d and 244d46e, locally and against https://showtime-window.vercel.app, about 11:40 PM to midnight Eastern on 2026-10-16
projects: [showtime-window]
---

# Showtime window's Veezi conformance run

Measured once, after adding a Veezi adapter for five Atlas Cinemas theaters and two more Cinemark theaters.

## Fixture tests (`npm test`)

26 tests, 26 passed. Against a saved Atlas Shaker Square sessions page: T1 (181 sessions), T2, T3 (every UTC start, shown in America/New_York, equals the listed local time), T6 (every runtime present and 40–300 minutes), ticket links to Veezi purchase pages, and 35 days covered.

## Live tests (`npm run conformance:live`)

The first local run used today, 2026-10-16, at about 11:40 PM: every theater returned 0 showings for today, because the day's shows had started. The script now uses tomorrow as "today" after 8 PM. Rerun, both direct from home and through Vercel (T5):

| Theater | 2026-10-17 | 2026-10-22 |
| --- | --- | --- |
| Cinemark Valley View XD and IMAX | 118 | 18 |
| Cinemark Strongsville at Southpark Mall | 77 | 10 |
| Cinemark Cuyahoga Falls and XD | 52 | 10 |
| Atlas Shaker Square 6 | 20 | 12 |
| Atlas Barrington 10 | 43 | 8 |
| Atlas Diamond Center 16 | 59 | 54 |
| Atlas Great Lakes Stadium 16 | 62 | 63 |
| Atlas Lakeshore 7 | 23 | 2 |

0 showings on another day in every case. The Vercel run is the first evidence that Veezi answers from a datacenter.

## Browser test (`npm run e2e`)

New checks: the Atlas Shaker Square shortlist (travel 25) matches `fit.ts` on the API data (3 shown, 3 expected), and travel minutes are kept per theater.

The first production run failed 2 checks: after switching theaters, the "At the theater" list still showed the previous theater's showings ("7:45 PM Other Mommy", an Atlas showing, while Valley View was selected), so the logged visit went to the wrong theater. Fixed in 244d46e. Afterwards, 13 of 13 checks passed locally and in production.

A screenshot of Atlas Shaker Square, 2026-10-17, 6–10 PM, listed Forgotten Island 6:45, Your Mother Your Mother Your Mother 7:00, Whalefall 7:15 and Other Mommy 7:45 PM, with 2 more too long. Those times match the Veezi page's visible list for that day ([Atlas observation](2026-10-16-atlas-cinemas-veezi-sessions-data.md)).
