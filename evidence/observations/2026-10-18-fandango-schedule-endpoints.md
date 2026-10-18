---
kind: observation
title: Fandango's schedule endpoints work from a script and from Vercel
date: 2026-10-18
source: Headless Chromium on https://www.fandango.com/atlas-cinemas-at-shaker-square-aafyb/theater-page, curl from a home connection, and a temporary route on a protected Vercel preview deployment (branch probe-fandango, since deleted), 2026-10-18 about 8:00–8:20 AM Eastern
projects: [showtime-window]
---

# Fandango's schedule endpoints work from a script and from Vercel

The theater page loads showtimes from `/napi/theaterMovieShowtimes/{id}?chainCode=…&startDate=…&isdesktop=true&partnerRestrictedTicketing=`, and lists nearby theaters from `/napi/nearbyTheaters?limit=…&zipCode=…`.

| Request | Result |
| --- | --- |
| Showtimes, browser user agent, full query string | 200 for Atlas Shaker Square (AAFYB) |
| Same for four other theaters, no `referer` | 403, `{"error":"FORBIDDEN","errorMessage":"Session expired or invalid token"}` |
| Same with `referer: https://www.fandango.com/` | 200 for AMC Ridge Park, Capitol, Silverspot Pinecrest, B&B Macedonia, Cinematheque, Cinemark Valley View, Regal Crocker Park |
| Showtimes without a user agent, short query string | 403 |
| Nearby theaters with a referer, four Northeast Ohio ZIP codes | 37 distinct theaters |
| From the Vercel preview, four theaters | 200 each: AMC Ridge Park 27, Capitol 12, Silverspot 44, Atlas Shaker Square 20 showtimes, the same counts as from home |

- Each showtime has a local `ticketingDate` (`2026-10-18+13:30`), `isSoldOut`, a format group (`Standard`, `Premium Format`, `3D`) and a ticket link. Each movie has `runtime` in minutes. Each theater has coordinates and `isOpen`.
- Cross-check for 2026-10-18: Cinemark Valley View had 120 showtimes on Fandango and 120 from cinemark.com, the same start times. Atlas Shaker Square had 20 and 20 from Veezi, the same start times.
- The 37 theaters include 8 Regal, 2 AMC, Phoenix Great Northern (North Olmsted), Silverspot Pinecrest (Beachwood), B&B Macedonia, Hickory Ridge (Brunswick), the Capitol, Kent Plaza, the Lyric (Wooster), the Nightlight (Akron), Lake 8 (Barberton) and Ravenna 7. Three carry chain code `VZ` (Amherst Cinema, Boardman Movies 8, The CUT); that `VZ` means Veezi is a guess.
- Fandango's terms of use prohibit this use ([terms observation](2026-10-18-theater-site-terms-on-automated-access.md)).
