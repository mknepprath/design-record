---
kind: observation
title: Showtime window combined-list test run
date: 2026-10-17
source: npm test and npm run e2e in showtime-window at commit 29cfcc2, locally and against https://showtime-window.vercel.app, 2026-10-17
projects: [showtime-window]
---

# Showtime window combined-list test run

Measured once. `npm test`: 31 of 31 passed, including the cross-theater shortlist and the travel estimate helpers.

`npm run e2e`, phone-sized headless Chromium, browser clock at 7:40 PM on 2026-10-18, location set to 41.388, −81.604 (near Valley View). Window 6:00–10:00 PM. Local and production gave the same results, 13 of 13:

| Check | Result |
| --- | --- |
| Valley View alone, travel 15, trailers 20, matches `fit.ts` on the API data | 14 shown, 14 expected |
| Trailers 0 | 18, matching |
| Valley View (travel 15) and Atlas Shaker Square (travel 25) in one list, in film-start order, each card naming its theater | 17 shown, 17 expected |
| Ticked theaters and travel minutes persist across a reload | pass |
| "Estimate travel from my location" | Valley View 15 min, Atlas Diamond Center 55 min, "by road with no traffic, plus 5 min to park" |
| A measured 22-minute start at Valley View changes trailers for Valley View only | pass |
| Visit logger, persistence and removal; no page errors | pass |

The estimates came from router.project-osrm.org once per run. Neither was compared with a real drive.
