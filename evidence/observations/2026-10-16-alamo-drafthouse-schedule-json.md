---
kind: observation
title: Alamo Drafthouse serves a public schedule JSON per market
date: 2026-10-16
source: curl and Python inspection of an Alamo Drafthouse endpoint, 2026-10-16
projects: [showtime-window]
---

# Alamo Drafthouse serves a public schedule JSON per market

`GET https://drafthouse.com/s/mother/v2/schedule/market/austin` returned 200 and 1,301,907 bytes of JSON with no cookies or headers beyond a user agent. One run.

- Top level: `data` with `presentations` (76), `sessions` (1,407), `market`, `formats` (5), `presentationAttributes`, `sessionAttributes`, `agePolicies`, `relatedPresentations`, `queues`.
- `market[0].cinemas`: 5 cinemas, each with `id`, `slug`, `timeZoneName` (`America/Chicago`), coordinates and address.
- Session fields: `cinemaId`, `sessionId`, `presentationSlug`, `status`, `businessDateClt`, `showTimeClt`, `showTimeUtc`, `cinemaTimeZoneName`, `screenNumber`, `formatSlug`, `reservedSeating`.
- Session dates: 53 distinct days, 2026-10-16 to 2027-01-29.
- Status counts: `ONSALE` 1,404, `PAST` 2, `SOLDOUT` 1.
- No runtime or duration field was found in `presentations` or `show` (searched keys containing `run`, `duration`, `imdb`, `tmdb`).
