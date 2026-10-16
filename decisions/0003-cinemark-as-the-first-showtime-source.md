---
kind: adr
number: 0003
title: Cinemark as the first showtime source
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [showtime-window]
evidence:
  - evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md
  - evidence/observations/2026-10-16-cinemark-screeningevent-data.md
  - evidence/observations/2026-10-16-theater-site-access-from-a-script.md
---

# ADR 0003: Cinemark as the first showtime source

## Context

The owner wants to start with one easy source ([owner's constraints](../evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md), quote 5). [Study 0001](../studies/0001-which-showtime-sources-can-be-read-for-free.md) found two conforming sources: Cinemark (43/45) and Alamo Drafthouse (40/45). The owner has a Cinemark nearby and no Alamo (quote 6).

## Decision

Build the Cinemark adapter first, against [Spec 0001](../specs/0001-showtime-source-conformance.md).

## Options considered

| Option | For | Against |
| --- | --- | --- |
| Cinemark | Nearby; runtime, format, availability and ticket link in one page ([observation](../evidence/observations/2026-10-16-cinemark-screeningevent-data.md)) | One request per theater per day |
| Alamo Drafthouse | 53 days per request | None nearby; no runtime |
| AMC, Regal | Large chains | Cloudflare 403 ([probe](../evidence/observations/2026-10-16-theater-site-access-from-a-script.md)) |

## Consequences

- No TMDB join is needed for the first version; `duration` comes from Cinemark.
- 7 days × N theaters = 7N requests per build, every 6 hours.
