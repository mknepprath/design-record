---
kind: observation
title: Cleveland Cinemas API accepts only 3 AM ranges and repeated ids
date: 2026-10-17
source: curl calls to https://www.clevelandcinemas.com/api/gatsby-source-boxofficeapi/ and the site's theater list, about 9:30 AM Eastern on 2026-10-17
projects: [showtime-window]
---

# Cleveland Cinemas API accepts only 3 AM ranges and repeated ids

Follow-up to the [first look at the API](2026-10-16-cleveland-cinemas-schedule-api.md), for Cedar Lee (X04Q9):

| Request | Result |
| --- | --- |
| `schedule` from `2026-10-17T03:00:00` to `2026-10-18T03:00:00` | 21 showtimes, 12:15 PM to 10:25 PM |
| same, from `T04:00:00` to `T04:00:00` | 0 showtimes |
| from `2026-10-17T04:00:00` to `2026-10-31T04:00:00` | 93 showtimes, starting 2026-10-18: the 17th is dropped |
| from `2026-10-17T00:00:00` to `2026-10-31T00:00:00` | 114 showtimes, starting 2026-10-17 |
| `movies?ids=327174,1000010389` | `[]` |
| `movies?ids=327174&ids=1000010389` | both movies |
| `movies?ids[]=…` | `null` |

- Reading of the table (inference): a day is included only when the range covers its 3 AM start.
- Five quick calls in a row each returned HTTP 500 with `null`. The same URL, called three times two seconds apart, returned 200 each time. The cause, a rate limit or plain flakiness, isn't known.
- Over two weeks (2026-10-16 to 2026-10-30), every showtime had the tag `Format.Projection.Digital` and no other format. `occupancy.rate` was `null` on all of them, so sold-out status isn't available.
- "12 Hours of Terror" has a runtime of 43,200 seconds (12 hours); it is a marathon, not an error. The other 18 films run 85 to 209 minutes.
- The site's theater list (`page-data/sq/d/2506275789.json`) gives coordinates: Cedar Lee 41.5008, −81.5653 (2163 Lee Road); Apollo Theatre, Oberlin, 41.2916, −82.2168 (19 East College Street). Neither is marked closed.
