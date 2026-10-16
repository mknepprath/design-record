---
kind: observation
title: Theater site access from a script
date: 2026-10-16
source: curl from a home connection (macOS), 2026-10-16
projects: [showtime-window]
---

# Theater site access from a script

One run each, not repeated. Command, with a desktop Chrome user agent:

```sh
UA='Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0 Safari/537.36'
curl -s -L -o /dev/null -w '%{http_code} %{size_download} %{url_effective}' -A "$UA" --max-time 20 "$URL"
```

| URL | Status | Bytes | What came back |
| --- | --- | --- | --- |
| https://www.amctheatres.com/movies | 403 | 5,488 | `server: cloudflare`, `__cf_bm` cookie, `cf-ray` header |
| https://www.regmovies.com/ | 403 | 5,512 | `server: cloudflare`, `__cf_bm` cookie |
| https://www.studiomoviegrill.com/ | 403 | 5,681 | Vendor not identified |
| https://www.marcustheatres.com/ | 200 | 212 | Imperva/Incapsula challenge: `<META NAME="ROBOTS" CONTENT="NOINDEX, NOFOLLOW">` and a script from `/_Incapsula_Resource` |
| https://www.cinemark.com/ | 200 | 195,664 | Full page |
| https://drafthouse.com/ | 200 | 4,498 | JavaScript app shell |
| https://www.harkins.com/ | 200 | 32,653 | Page (not inspected) |
| https://www.landmarktheatres.com/ | 200 | 173,509 | Page, no `ld+json`, no `ScreeningEvent` |
| https://www.showcasecinemas.com/ | 200 | 199,622 | Page, no `ld+json`, no `ScreeningEvent` |
| https://www.fandango.com/ | 200 | 260,022 | Page |
| https://www.fandango.com/austin_tx_movietimes | 200 | 220,864 | One `ld+json` block, no `ScreeningEvent` |
| https://angelikafilmcenter.com/ | 200 | 5,793 | JavaScript app shell |
| https://www.harkins.com/theatres/tempe-marketplace-16 | 404 | 10,237 | Guessed URL; wrong |

robots.txt, fetched the same day:

- `https://www.cinemark.com/robots.txt` disallows `/employee-forms/`, `/media/`, `/tickets/`, `/ticketseatmap/`, `/TicketSeatMap`, `/shoppingcart`, `/Tickets/`. Theater pages are not disallowed.
- `https://drafthouse.com/robots.txt` is `User-agent: *` / `Allow: /`.
