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
- One showtime source to start: Cinemark ([ADR 0003](../decisions/0003-cinemark-as-the-first-showtime-source.md)), Valley View, Ohio ([owner's answer](../evidence/feedback/2026-10-16-owner-picks-theater-and-hosting.md)).
- Hosted on Vercel, fetching on demand ([ADR 0011](../decisions/0011-host-showtime-window-on-vercel.md)).

## Scope

- In: one list across every theater within a max travel time, with travel per theater; Cinemark Valley View XD and IMAX, plus ([ADR 0012](../decisions/0012-add-atlas-cinemas-via-veezi.md)) Cinemark Strongsville and Cuyahoga Falls, five Atlas Cinemas theaters and ([ADR 0014](../decisions/0014-add-cleveland-cinemas-via-its-box-office-api.md)) Cleveland Cinemas' Cedar Lee and Apollo; any day the source lists, padding set per theater, travel minutes estimated from your location or entered by hand, a link to buy tickets.
- Direction: as many theaters as possible; none get removed once supported ([owner, 2026-10-17](../evidence/feedback/2026-10-17-owner-wants-as-many-theaters-as-possible.md)).
- Out: accounts, AMC and Regal (blocked by Cloudflare), seat maps, paid data feeds, a native app.

## Plan

| Phase | Deliverable | Done when |
| --- | --- | --- |
| 1 | `fit.ts` and its tests, on fixture data | Tests pass for midnight crossing, timezone, zero padding, the credits toggle |
| 2 | Cinemark adapter passing [Spec 0001](../specs/0001-showtime-source-conformance.md), cached API route on Vercel | T1–T6 pass, T5 from Vercel |
| 3 | Window picker and shortlist | Tonight's shortlist matches the Cinemark site when checked by hand |
| 4 | Measured [padding](../glossary/padding.md) per theater | After 3 or more visits, padding reflects measured starts |
| 5 | More sources | One broken source doesn't break the build |

## Status

Live at [showtime-window.vercel.app](https://showtime-window.vercel.app) since 2026-10-16.

| Phase | State |
| --- | --- |
| 1 Fit logic | Done: `lib/fit.ts` with tests |
| 2 Cinemark adapter and API | Done: T1–T6 pass, T5 from Vercel ([first conformance run](../evidence/observations/2026-10-16-showtime-window-first-conformance-run.md)) |
| 3 Window picker and shortlist | Done: the shortlist for 2026-10-17, 6–10 PM, matched the Cinemark page ([hand check](../evidence/observations/2026-10-16-showtime-window-shortlist-hand-check.md)). A [browser test](../evidence/observations/2026-10-16-showtime-window-browser-test-run.md) drives the form and matches the rendered list. |
| 4 Measured padding | Tool built: "At the theater" logs when a film starts and offers the median per theater as the trailer minutes. No visits logged yet. |
| 5 More sources | Veezi adapter live for five Atlas theaters, plus two more Cinemarks ([ADR 0012](../decisions/0012-add-atlas-cinemas-via-veezi.md), accepted; [conformance run](../evidence/observations/2026-10-16-showtime-window-veezi-conformance-run.md)). Cleveland Cinemas (Cedar Lee, Apollo) live through its box office API ([ADR 0014](../decisions/0014-add-cleveland-cinemas-via-its-box-office-api.md); [test run](../evidence/observations/2026-10-17-showtime-window-cleveland-cinemas-and-travel-limit-test-run.md)). Every Cinemark in Ohio and bordering states, from the sitemap ([ADR 0016](../decisions/0016-every-cinemark-in-ohio-and-bordering-states.md)): 53 theaters in all. Cinemark rate-limits requests from Vercel ([observation](../evidence/observations/2026-10-17-cinemark-rate-limits-vercel.md)). |

Since [ADR 0013](../decisions/0013-one-shortlist-across-theaters.md), one list covers several theaters, each with its own travel minutes, which can be estimated from your location. Since [ADR 0015](../decisions/0015-include-theaters-by-travel-time.md), it covers every theater within a max travel time, minus any turned off.

Open decision: Cinemark's and Cleveland Cinemas' terms prohibit automated collection, as do Fandango's, which would add Regal and AMC ([Study 0003](../studies/0003-more-northeast-ohio-theaters-and-what-their-terms-allow.md)).

Since 2026-10-18 the GitHub repo is connected to Vercel: pushes to `main` deploy to production. Commit 700dc90 built and went live within about a minute of the push.

## Private documents

The repo is private; these links work for collaborators only.

- [showtime-window repo](https://github.com/mknepprath/showtime-window): code, README, tests
