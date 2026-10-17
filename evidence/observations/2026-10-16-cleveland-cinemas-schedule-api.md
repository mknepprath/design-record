---
kind: observation
title: Cleveland Cinemas serves schedules from a public JSON endpoint
date: 2026-10-16
source: Network requests recorded in headless Chromium on https://www.clevelandcinemas.com/our-locations/x04q9-cedar-lee-theatre/, then the same endpoints called with curl, about 11:05 PM Eastern on 2026-10-16
projects: [showtime-window]
---

# Cleveland Cinemas serves schedules from a public JSON endpoint

The site is a Gatsby app built on Webedia's theater website platform (images come from `cms-assets.webediamovies.pro`). Showtimes aren't in the HTML. The browser loads them from the site's own API, and the same requests work from `curl` with no cookies or tokens:

| Request | Returns |
| --- | --- |
| `/api/gatsby-source-boxofficeapi/schedule?from=2026-10-17T03:00:00&to=2026-10-18T03:00:00&theaters={"id":"X04Q9","timeZone":"America/New_York"}` | Per movie ID, per date, showtimes with `startsAt` (local time, no offset: `2026-10-17T21:40:00`), screen, format tags (`Format.Projection.Digital`) and ticket URLs (`booking.clevelandcinemas.com/launch/ticketing/…`) |
| `/api/gatsby-source-boxofficeapi/movies?ids=327174` | Title and `runtime` in seconds (`"Digger"`, 7740) |
| `/page-data/sq/d/2506275789.json` | Theaters, with ID and IANA `timeZone` |

- The range runs from 3:00 AM to 3:00 AM: the browser's own request for "today" used that window.
- Cedar Lee, 2026-10-17: 10 films, 21 showtimes. A two-week range returned dates 2026-10-17 through 2026-10-30.
- The theater list has two theaters: Apollo Theatre (X03GQ) and Cedar Lee Theatre (X04Q9). The Capitol isn't in it.
- The page loads Google reCAPTCHA, but the API calls above didn't require it.
- The first `schedule` call, copied from a truncated log, returned HTTP 500 with `null`; the full URL returned 200.
