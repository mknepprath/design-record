---
kind: project
title: Showtime window
status: active
owner: Michael Knepprath
date: 2026-10-16
---

# Showtime window

A web app where you enter the time you have free and get back the movies you can watch inside it, counting trailers and travel.

## Problem

Theater apps start from the movie: pick a film, then a theater, then a time. Someone who starts from "I'm free 6–9pm" has to read every film's times and add up runtimes and trailers by hand ([owner's constraints](../evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md), quote 1). Whether every major app works this way has been checked only from recall, not in the apps.

## Goal

For a chosen window and a saved set of theaters, list every showing where leaving, trailers, the film and getting home all fit inside the window. Sort by start time and show when the film itself starts and ends.

## Constraints

- $0 beyond hosting ([quote 4](../evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md)).
- Web first ([ADR 0001](../decisions/0001-build-a-web-app-first.md)).
- One showtime source to start: Cinemark ([ADR 0003](../decisions/0003-cinemark-as-the-first-showtime-source.md)).

## Scope

- In: one Cinemark theater, today plus 6 days, padding set per theater, travel minutes entered by hand, a link to buy tickets.
- Out: accounts, AMC and Regal (blocked by Cloudflare), seat maps, paid data feeds, a native app.

## Plan

| Phase | Deliverable | Done when |
| --- | --- | --- |
| 1 | `fit.ts` and its tests, on fixture data | Tests pass for midnight crossing, timezone, zero padding, the credits toggle |
| 2 | Cinemark adapter passing [Spec 0001](../specs/0001-showtime-source-conformance.md), scheduled build, `meta.json` | JSON refreshes on its own; T5 passes from CI |
| 3 | Window picker and shortlist | Tonight's shortlist matches the Cinemark site when checked by hand |
| 4 | Measured [padding](../glossary/padding.md) per theater | After 3 or more visits, padding reflects measured starts |
| 5 | More sources | One broken source doesn't break the build |

## Status

Nothing built. Planning documents only: [Study 0001](../studies/0001-which-showtime-sources-can-be-read-for-free.md), [Spec 0001](../specs/0001-showtime-source-conformance.md), ADRs [0001](../decisions/0001-build-a-web-app-first.md), [0002](../decisions/0002-prebuild-showtimes-as-static-json.md), [0003](../decisions/0003-cinemark-as-the-first-showtime-source.md).

Not checked yet: whether the Cinemark page loads from GitHub Actions IP addresses; how many days ahead `showDate` reaches; Cinemark's terms of use.
