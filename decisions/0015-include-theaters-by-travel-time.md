---
kind: adr
number: 0015
title: Include theaters by travel time, not by ticking each one
status: accepted
owner: Michael Knepprath
date: 2026-10-17
projects: [showtime-window]
evidence:
  - evidence/feedback/2026-10-17-owner-wants-as-many-theaters-as-possible.md
  - evidence/observations/2026-10-17-showtime-window-cleveland-cinemas-and-travel-limit-test-run.md
---

# ADR 0015: Include theaters by travel time, not by ticking each one

## Context

[ADR 0013](0013-one-shortlist-across-theaters.md) put every ticked theater in one list. The owner wants as many theaters as possible ([feedback](../evidence/feedback/2026-10-17-owner-wants-as-many-theaters-as-possible.md)). Ticking each theater doesn't scale: a new theater stays out until someone finds and ticks it.

## Decision

- A "Max travel" setting, 30 minutes each way by default, blank for no limit.
- The list covers every theater whose travel time is known and within the limit, except theaters turned off. Turning one off is remembered, even if it later falls out of range and comes back in.
- A theater with no travel time stays out. The old 15-minute default would have put Oberlin as close as Valley View. On a first visit, the page asks for a location estimate or a typed travel time.
- The picker lists theaters in range, nearest first. The rest fold under "N more". The order changes only when the limit or the estimate changes, so a row doesn't move while its travel time is being typed.
- Settings saved by the hand-picked version are kept. Each theater picked there keeps its travel time, or gets the old default if it had none.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| Limit by travel time, opt out | New theaters show up by themselves; one number to set | Needs travel times, so a first visit needs a location or typing |
| Limit by distance (km) | No routing call | Distance isn't what fits a window; travel time is already computed |
| Keep ticking, opt in | Explicit | Doesn't scale; new theaters hidden |

## Consequences

- With the limit set, every theater in range costs one request per date shown, so a large catchment means many requests on a cold cache.
- The estimate still sends every theater to OSRM in one request. Past 100 or so theaters that may exceed the server's table limit. That limit is recalled, not checked.
