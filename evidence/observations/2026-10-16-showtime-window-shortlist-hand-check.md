---
kind: observation
title: Showtime window shortlist checked against the Cinemark site
date: 2026-10-16
source: Rendered Cinemark Valley View page for 2026-10-17 (headless Chrome, visible text) compared with https://showtime-window.vercel.app/api/showtimes and lib/fit.ts, about 10:40 PM Eastern on 2026-10-16
projects: [showtime-window]
---

# Showtime window shortlist checked against the Cinemark site

Done once, for Saturday 2026-10-17, by script rather than by eye. The theater page was rendered in a headless browser and its visible text read: film title, rating, runtime ("1 hr 59 min") and the listed times. This path doesn't use the schema.org data the adapter reads.

## Showings

| | Cinemark page (visible) | `/api/showtimes` |
| --- | --- | --- |
| Films | 28 | 28 |
| Showings | 118 | 118 |
| Runtimes that differ | | 0 |

Every title and local start time matched, except that two Movie names in the schema data had a trailing space ("The Social Reckoning ", "Spirited Away (English Dubbed) Studio Ghibli Fest 2026 "). The adapter now trims them (showtime-window commit ee4a49c), and the deployed API returned 0 untrimmed titles for 2026-10-17 and 2026-10-18 afterward.

## Shortlist

Window 6:00–10:00 PM, 20 minutes of padding, 15 minutes of travel each way, no credits skipped. The expected list was worked out from the visible page data with a separate script: leave (listed − 15) at or after 6:00, home (listed + 20 + runtime + 15) at or before 10:00. `shortlist()` run on the API data gave the same 9 showings:

| Listed | Film |
| --- | --- |
| 6:20 PM | Your Mother Your Mother Your Mother |
| 6:25 PM | You Can See Everything |
| 6:50 PM | The Devils |
| 6:55 PM | Whalefall |
| 7:00 PM | Street Fighter |
| 7:00 PM | The Devils |
| 7:05 PM | Sense and Sensibility |
| 7:15 PM | Verity |
| 7:45 PM | You Can't Leave |

20 more showings started in time but ran too long.

## Limits

- The web page's form and result rendering weren't driven in this check; it compared data and fit logic only.
- Whether the padding of 20 minutes is right isn't tested here: that is phase 4.
