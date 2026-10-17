---
kind: spec
number: 0001
title: Showtime source conformance
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [showtime-window]
evidence:
  - evidence/observations/2026-10-16-theater-site-access-from-a-script.md
  - evidence/observations/2026-10-16-cinemark-screeningevent-data.md
  - evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md
  - evidence/observations/2026-10-16-showtime-window-first-conformance-run.md
---

# Spec 0001: Showtime source conformance

## Scope

What a showtime source must provide before the [showtime window](../projects/showtime-window.md) app supports it, how sources are scored, and the tests every source adapter must pass. Only free access counts: public pages, public JSON a site already serves, and free API tiers ([owner's constraints](../evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md), quote 4).

Assessment of actual sources is in [Study 0001](../studies/0001-which-showtime-sources-can-be-read-for-free.md).

## Requirements

MUST, SHOULD and MAY as in RFC 2119. A source conforms when it meets every MUST.

| ID | Level | Requirement | Why |
| --- | --- | --- | --- |
| R1 | MUST | Film title for each showing | Show the film and match it to TMDB |
| R2 | MUST | Exact start time for each showing, to the minute | All fit math starts from this |
| R3 | MUST | Timezone stated, or inferable from the theater's location | Prevents off-by-hours errors |
| R4 | MUST | Showings for today; SHOULD cover 7 days | Plan tonight or this week |
| R5 | MUST | Readable with a plain HTTP request from a script: no login, CAPTCHA or JavaScript challenge | The build runs in CI with no browser |
| R6 | MUST | Stable theater identifier (ID or slug) | Keys the JSON files and settings |
| R7 | SHOULD | Runtime in minutes | Avoids a TMDB lookup |
| R8 | SHOULD | Format or auditorium type (IMAX, Dolby, 3D, D-BOX) | Padding may differ by format |
| R9 | SHOULD | Ticket or showtime deep link | One tap from the shortlist to buying |
| R10 | SHOULD | Structured data (JSON or schema.org) rather than visible HTML | Breaks less often |
| R11 | MAY | Sold-out or availability flag | Hide showings that can't be booked |
| R12 | MAY | Published preview length or "feature starts" time | Replaces estimated [padding](../glossary/padding.md) |
| R13 | MAY | Release ID (TMDB, IMDb) | Exact runtime join |

A source also fails if its terms forbid automated access *and* it actively blocks it. Terms alone are recorded as risk, because use is personal and low-volume.

## Scoring

Conforming sources only. Each criterion is scored 0–3 and multiplied by its weight; the maximum is 45.

| Criterion | Weight | 3 | 0 |
| --- | --- | --- | --- |
| Access path | 4 | Structured data in the page, or a public JSON endpoint | Rendered only by a JavaScript app behind an internal API |
| Bot protection | 4 | None seen from a plain `curl` | Cloudflare/Akamai/Imperva challenge or 403 |
| Data completeness (R7–R13) | 3 | Runtime, format and ticket link present | Title and time only |
| Coverage per adapter | 2 | One adapter covers a chain or platform | One theater only |
| Terms / legal risk | 1 | Data offered openly (API, schema.org) | Terms explicitly forbid scraping |
| Stability | 1 | URL and markup unchanged for years | Redesigned often |

## Adapter contract

```ts
export type Showing = {
  sourceId: "cinemark";
  theaterId: string;          // R6
  title: string;              // R1
  startUtc: string;           // R2 + R3, ISO 8601 with Z
  theaterTz: string;          // IANA, e.g. "America/Chicago"
  runtimeMin?: number;        // R7
  format?: string;            // R8
  ticketUrl?: string;         // R9
  soldOut?: boolean;          // R11
};

export interface SourceAdapter {
  id: Showing["sourceId"];
  fetchDay(theaterId: string, date: string /* YYYY-MM-DD, theater-local */): Promise<Showing[]>;
}
```

## Conformance tests

| Test | Pass condition | Requirement |
| --- | --- | --- |
| T1 Fixture parse | A saved response parses to at least one `Showing` without throwing | R1, R2 |
| T2 Required fields | Every showing has a non-empty `title`, a parseable `startUtc` and a `theaterId` | R1, R2, R6 |
| T3 Timezone round trip | `startUtc` shown in `theaterTz` equals the local time on the site, for 3 sampled showings | R3 |
| T4 Date window | `fetchDay` for today and today + 6 each return only showings on that local day | R4 |
| T5 Plain fetch from production | A live request from the production server (Vercel, [ADR 0011](../decisions/0011-host-showtime-window-on-vercel.md)) returns 200 with no browser | R5 |
| T6 Runtime sanity | When present, `runtimeMin` is between 40 and 300 | R7 |
| T7 Drift alarm | 0 showings for a theater that had more than 0 the day before fails the build and keeps the previous JSON | Stability |

T5 mattered most, because datacenter IP addresses may be treated differently from a home connection. It passed from Vercel on 2026-10-16, along with T1–T4 and T6 ([first conformance run](../evidence/observations/2026-10-16-showtime-window-first-conformance-run.md)). T7 isn't implemented: there is no stored previous day to compare against under ADR 0011.

### Cinemark notes for the adapter

From the [Cinemark observation](../evidence/observations/2026-10-16-cinemark-screeningevent-data.md):

- The schema.org `@graph` (MovieTheater, Movie, ScreeningEvent) is inside the React Server Components payload: the string chunks pushed to `self.__next_f`. Join and JSON-decode those chunks, then parse each `{"@context":"https://schema.org"…}` object. The film title comes from the Movie node that `workPresented` points to.
- `startDate` carries a UTC offset (`-05:00`); `duration` is ISO 8601 (`PT1H59M`).
- The day is selected with `?showDate=YYYY-MM-DD`; `?date=` is ignored.
- `theaterId` = the path after `/theatres/` (e.g. `tx-austin/cinemark-southpark-meadows`).
