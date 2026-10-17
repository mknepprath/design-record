---
kind: observation
title: Showtime window's first conformance run
date: 2026-10-16
source: npm test and npm run conformance:live in showtime-window at commit d094286, 2026-10-16
projects: [showtime-window]
---

# Showtime window's first conformance run

Measured once, on 2026-10-16 between about 9:30 and 10:00 PM Eastern.

## Fixture tests (`npm test`)

Against a saved Cinemark Valley View page from 2026-10-16: 14 tests, 14 passed. Of the Spec 0001 tests, T1 (22 showings parsed), T2, T3 (all 22 UTC start times match the local time in the event name) and T6 passed. The rest are tests of the fit logic.

## Live tests (`npm run conformance:live`)

| Run from | T4 today (2026-10-16) | T4 today + 6 (2026-10-22) | T5 |
| --- | --- | --- | --- |
| Home connection, direct to cinemark.com | 22 showings, 0 on another day | 18 showings, 0 on another day | pass |
| Vercel production, through `/api/showtimes` (https://showtime-window.vercel.app) | 21 showings, 0 on another day | 18 showings, 0 on another day | pass |

The Vercel run is the first evidence that Cinemark answers requests from a datacenter. That was the open risk in Spec 0001.

The difference between 22 and 21 today is one showing that started between the two runs; the page drops showings once they've started (not checked).

A repeated request to `/api/showtimes` returned `x-vercel-cache: HIT`.
