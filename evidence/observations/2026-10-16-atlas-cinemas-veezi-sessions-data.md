---
kind: observation
title: Atlas Cinemas' Veezi pages carry schema.org sessions
date: 2026-10-16
source: curl of https://www.atlascinemas.net/ and its five Veezi sessions pages from a home connection, about 11:00 PM Eastern on 2026-10-16
projects: [showtime-window]
---

# Atlas Cinemas' Veezi pages carry schema.org sessions

One fetch of each page with a desktop Chrome user agent. Nothing was blocked.

The Atlas Cinemas home page links each theater's showtimes to Veezi, a ticketing platform: `https://ticketing.useast.veezi.com/sessions/?siteToken={token}`.

Each sessions page has two `application/ld+json` blocks:

- A `MovieTheater` with name and address.
- A list of `VisualArtsEvent` nodes, one per session, covering every listed day. Each has `name` (film title), `startDate` with a UTC offset (`2026-10-17T13:00:00-04:00`), `duration` (`PT1H52M`) and `url` (a `/purchase/{id}` ticket link). Example node, reformatted:

```json
{ "@type": "VisualArtsEvent", "startDate": "2026-10-17T13:00:00-04:00", "duration": "PT1H52M",
  "name": "Your Mother Your Mother Your Mother",
  "url": "https://ticketing.useast.veezi.com/purchase/317657?siteToken=3t9ejpjygsgfjz7s7kyfm5ja6m" }
```

No format field was seen in the data. The visible page marks some sessions with a screen note ("Upstairs").

| Theater (Veezi `MovieTheater.name`) | siteToken | Sessions | Days listed |
| --- | --- | --- | --- |
| Atlas Cinemas Shaker Square 6 | 3t9ejpjygsgfjz7s7kyfm5ja6m | 181 | 35 |
| Atlas Cinemas Barrington 10 | cajpmttby2nakgmxfetenxjy64 | 325 | 34 |
| Atlas Cinemas Diamond Center 16 | gdh71tz290ac87dc5r7astdjb0 | 351 | 12 |
| Atlas Cinemas Great Lakes Stadium 16 | tgy9knxrsty7403j245xzmasew | 769 | 53 |
| Atlas Cinemas Lakeshore 7 | wpqhgfd9mmxf5krjt61b4yega0 | 115 | 12 |

Runtimes for films also at Cinemark Valley View on 2026-10-17 matched Cinemark's to within 4 minutes, e.g. Whalefall 106 vs 106, Resident Evil 90 vs 89, Crawlers 82 vs 86.

A web search listed an Atlas Eastgate 10 in Mayfield Heights; the Atlas home page doesn't link it. Whether it's still open wasn't checked.
