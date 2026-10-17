---
kind: adr
number: 0002
title: Prebuild showtimes as static JSON on a schedule
status: superseded
owner: Michael Knepprath
date: 2026-10-16
projects: [showtime-window]
evidence:
  - evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md
  - evidence/observations/2026-10-16-free-showtime-data-options.md
---

# ADR 0002: Prebuild showtimes as static JSON on a schedule

Superseded by [ADR 0011](0011-host-showtime-window-on-vercel.md) before it was built.

## Context

The budget is $0 beyond hosting ([owner's constraints](../evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md), quote 4). Browsers can't fetch theater pages directly because of CORS. Vercel Hobby cron runs at most once per day ([search results](../evidence/observations/2026-10-16-free-showtime-data-options.md), via summary). lily-dex-api already runs a scheduled build that publishes static JSON to GitHub Pages.

## Decision

A GitHub Actions workflow runs every 6 hours. It fetches each saved theater for today plus 6 days, writes `data/{theaterId}/{YYYY-MM-DD}.json` and `meta.json` (build time, per-source status), and commits. The site is static and does the fit math in the browser.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| GitHub Actions → static JSON | $0; same pattern as lily-dex-api; any schedule | Datacenter IP addresses may be blocked (Spec 0001, T5); Pages on a private repo needs a paid plan (owner's notes, not checked here) |
| Vercel serverless route with cache | Fresh on demand | Same IP concern; Hobby is non-commercial only (recalled, not checked) |
| Vercel cron → Blob | Integrated | Once per day at most on Hobby |
| Local launchd job pushes JSON | Home IP; known to work for Cinemark | Depends on a machine being on |

## Consequences

- Showtimes can be up to 6 hours stale. Sold-out flags are therefore advisory.
- If T5 fails in CI, switch to the local job without changing the site.
- GitHub disables scheduled workflows after 60 days without repo activity (recalled, not checked).
