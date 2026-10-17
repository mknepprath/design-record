---
kind: observation
title: Cleveland-area theater sites from a script
date: 2026-10-16
source: curl from a home connection and web searches, about 10:50–11:15 PM Eastern on 2026-10-16
projects: [showtime-window]
---

# Cleveland-area theater sites from a script

## Probes

One `curl -L` each with a desktop Chrome user agent, as in the [first access probe](2026-10-16-theater-site-access-from-a-script.md).

| URL | Status | What came back |
| --- | --- | --- |
| https://www.atlascinemas.net/ | 200 | Static page linking five Veezi sessions pages ([Atlas observation](2026-10-16-atlas-cinemas-veezi-sessions-data.md)) |
| https://www.clevelandcinemas.com/showtimes/ | 200 | Gatsby app; showtimes loaded from JSON ([Cleveland Cinemas observation](2026-10-16-cleveland-cinemas-schedule-api.md)) |
| https://cinematheque.cia.edu/ | 200 | schema.org `ScreeningEvent` nodes for today only, with UTC offset and a purchase URL; `Movie` nodes with `duration` |
| https://cinematheque.cia.edu/2026-10-18 | 200 | Same `ld+json` (today's showings); the day's films and times only in visible HTML ("The Women (1939)… 2:00 pm") |
| https://www.regmovies.com/theatres/regal-willoughby-commons | 403 | Cloudflare challenge |
| https://www.amctheatres.com/movie-theatres/cleveland | 403 | Cloudflare challenge |
| https://www.atomtickets.com/theaters/atlas-cinemas-at-shaker-square/4556 | 403 | Cloudflare challenge |
| https://www.fandango.com/atlas-cinemas-at-shaker-square-AAFYB/theater-page | 200 | One `ld+json` block, no showtimes in the HTML (not traced further) |
| https://www.imdb.com/showtimes/cinema/US/ci0030592/US/44023 | 202 | 1,997 bytes, no schedule |
| https://www.clevelandart.org/events/film | 404 | Guessed URL; the museum's film listings weren't found |

## From web searches (secondary)

- The Capitol Theatre (Gordon Square) passed from Cleveland Cinemas to Atlanta-based Arthouse Management on July 31, 2026, per [Ideastream, 2026-05-16](https://www.ideastream.org/arts-culture/2026-05-16/changes-coming-to-gordon-squares-capitol-theatre). Its new showtime site wasn't found.
- Chagrin Cinemas closed in January 2023 ([WKYC](https://wkyc.com/article/entertainment/movies/chagrin-cinemas-closing-after-33-years/95-6ee503e2-ccab-4054-92de-9b3ca7da9b78)).
- Other Cinemark theaters near Valley View in Cinemark's sitemap: Strongsville at Southpark Mall, Cuyahoga Falls and XD. These read with the existing Cinemark adapter (same URL pattern; not fetched).
