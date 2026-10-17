---
kind: adr
number: 0011
title: Host Showtime window on Vercel and fetch showtimes on demand
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [showtime-window]
supersedes: [decisions/0002-prebuild-showtimes-as-static-json.md]
evidence:
  - evidence/feedback/2026-10-16-owner-picks-theater-and-hosting.md
  - evidence/observations/2026-10-16-showtime-window-first-conformance-run.md
---

# ADR 0011: Host Showtime window on Vercel and fetch showtimes on demand

Supersedes [ADR 0002](0002-prebuild-showtimes-as-static-json.md).

## Context

ADR 0002 proposed a GitHub Actions build that writes static JSON to GitHub Pages. The owner chose Vercel for the app and kept GitHub Pages for the design record ([owner's answer](../evidence/feedback/2026-10-16-owner-picks-theater-and-hosting.md)). Vercel Hobby cron runs at most once a day ([search results](../evidence/observations/2026-10-16-free-showtime-data-options.md)), so a scheduled build there would be stale.

## Decision

A Next.js app on Vercel. The route `/api/showtimes?theater=…&date=…` fetches the Cinemark page for that theater-day when asked. The result is cached for an hour with `use cache` (`revalidate` 3600 s, `expire` 6 h) and at Vercel's edge for 15 minutes. Only theaters in the app's own list can be requested.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| Fetch on demand, cached | Fresh within an hour; no schedule to maintain; only days someone looks at get fetched | First request for a theater-day waits on Cinemark (about 1 s, not measured) |
| GitHub Actions → static JSON (ADR 0002) | Works with no server | Owner chose Vercel |
| Vercel cron → stored JSON | Pre-warmed | Once a day on Hobby |

## Consequences

- T5 now means "from Vercel's servers", and it passed on 2026-10-16 ([conformance run](../evidence/observations/2026-10-16-showtime-window-first-conformance-run.md)).
- If Cinemark starts blocking Vercel, a failed fetch returns a 502 and the page says so; nothing falls back yet.
- Vercel Hobby is for non-commercial use (recalled, not checked).
