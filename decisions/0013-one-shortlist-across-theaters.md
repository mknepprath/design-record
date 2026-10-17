---
kind: adr
number: 0013
title: One shortlist across theaters, with travel per theater
status: accepted
owner: Michael Knepprath
date: 2026-10-17
projects: [showtime-window]
evidence:
  - evidence/feedback/2026-10-17-owner-wants-one-list-across-theaters.md
  - evidence/observations/2026-10-17-showtime-window-combined-list-test-run.md
---

# ADR 0013: One shortlist across theaters, with travel per theater

## Context

With eight theaters ([ADR 0012](0012-add-atlas-cinemas-via-veezi.md)), the app showed one theater at a time. The owner said this "goes against the whole point" and that travel would differ by theater ([owner's reply](../evidence/feedback/2026-10-17-owner-wants-one-list-across-theaters.md)). Travel depends on the theater and the starting point, not on the film, so eight theaters need eight numbers per starting point.

## Decision

- The page ticks any set of theaters and shows one list, sorted by when the film starts. Each card names its theater.
- Each theater has its own travel minutes, entered by hand or filled by "Estimate travel from my location".
- The estimate asks the browser for its location once and sends it, with the theaters' coordinates, to the public OSRM routing server (router.project-osrm.org). Minutes are the driving time with no traffic plus 5 to park, rounded up to 5. If OSRM doesn't answer, a straight-line guess (1.3x the distance at 50 km/h) is used and labeled as such. The location isn't stored; only the minutes are.
- Trailer minutes measured at a theater apply to that theater only. The trailer field is the default for the rest.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| Travel per theater, typed or estimated once from location | Eight numbers; works offline once set; no key or cost | Estimates ignore traffic; stale if you start somewhere else |
| Route from the current location on every search | Always from where you are | A location prompt and a third-party call on every visit |
| Straight-line distance only | Nothing leaves the browser | Roads and highways make it rough |
| Paid routing with traffic (Google, Mapbox) | Traffic-aware | Needs a key and possibly money ($0 budget) |

## Consequences

- The OSRM demo server has no service guarantee and asks for light use (recalled, not checked). The app calls it only when the button is pressed.
- Starting from somewhere else means pressing the button again or editing the minutes.
- Coordinates for Atlas Barrington and Diamond Center are at the street, not the building.
- Estimates haven't been compared with real drives ([test run](../evidence/observations/2026-10-17-showtime-window-combined-list-test-run.md)).
