---
kind: observation
title: Cinemark theater pages embed ScreeningEvent data
date: 2026-10-16
source: curl of a Cinemark theater page, 2026-10-16
projects: [showtime-window]
---

# Cinemark theater pages embed ScreeningEvent data

Page: https://www.cinemark.com/theatres/tx-austin/cinemark-southpark-meadows (200, 367,601 bytes). One theater, one run.

The HTML holds schema.org `ScreeningEvent` objects as escaped JSON inside the page, not in a `<script type="application/ld+json">` tag: a parser looking only for that tag finds nothing. One event, as it appeared (unescaped, description cut):

```json
{
  "@type": "ScreeningEvent",
  "eventStatus": "https://schema.org/EventScheduled",
  "location": {"@id": "https://www.cinemark.com/theatres/tx-austin/cinemark-southpark-meadows#theater"},
  "name": "Street Fighter 11:05 AM D-BOX Screening",
  "offers": {
    "@type": "Offer",
    "availability": "https://schema.org/InStock",
    "url": "https://www.cinemark.com/TicketSeatMap/?showtimeId=491800"
  },
  "startDate": "2026-10-16T11:05:00-05:00",
  "workPresented": {"@id": "https://www.cinemark.com/movies/street-fighter#movie"},
  "videoFormat": "D-BOX",
  "duration": "PT1H59M",
  "inLanguage": "en"
}
```

| Request | ScreeningEvents | Dates in `startDate` |
| --- | --- | --- |
| Theater page, no query | 65 | 2026-10-16 only |
| `?showDate=2026-10-18` | 73 | 2026-10-18 only |
| `?date=2026-10-18` | 65 | 2026-10-16 only (parameter ignored) |

Durations seen include `PT1H16M`, `PT1H29M`, `PT1H33M`, `PT1H59M`. How far ahead `showDate` returns data was not tested beyond +2 days.
