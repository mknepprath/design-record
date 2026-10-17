---
kind: observation
title: Cinemark's sitemap lists every theater, and each page gives its location
date: 2026-10-17
source: https://www.cinemark.com/sitemap.xml and theater pages fetched one per second by scripts/cinemark-catalog.ts (private showtime-window repo, commit f77f8ef), 2026-10-17 about 4:30 PM Eastern
projects: [showtime-window]
---

# Cinemark's sitemap lists every theater, and each page gives its location

- The sitemap lists 306 US theater pages (`/theatres/{state}-{city}/{slug}`). Texas has 81, California 54, Ohio 23. `robots.txt` doesn't disallow `/theatres/`.
- Each open theater's page carries a schema.org `MovieTheater` node with name, street address and `geo` coordinates. Its `ScreeningEvent` start times carry a UTC offset: `2026-10-17T16:10:00-05:00` at Paducah, Kentucky.
- For Ohio, Pennsylvania, Michigan, Indiana, Kentucky and West Virginia: 47 pages. 46 had a `MovieTheater` node and showtimes. The other, `oh-wooster/cinemark-wooster-movies-10-now-closed`, had neither.
- Two of the 46 showed −05:00 in daylight time, which is Central: Valparaiso, Indiana and Paducah, Kentucky. The rest showed −04:00, which is Eastern.
- Drive times from a point near Valley View, routed by OSRM with 5 minutes to park: 12 Cinemark and non-Cinemark theaters within 70 minutes, 15 within 120. Boardman was 95 routed against 145 from the straight-line formula.
