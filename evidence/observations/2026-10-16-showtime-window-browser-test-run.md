---
kind: observation
title: Showtime window browser test run
date: 2026-10-16
source: npm run e2e in showtime-window at commit a3d4392, against http://localhost:3123 and https://showtime-window.vercel.app, about 11:00 PM Eastern on 2026-10-16
projects: [showtime-window]
---

# Showtime window browser test run

`scripts/e2e.ts` drives the page in headless Chromium (Playwright, phone-sized viewport, browser clock set to 7:40 PM Eastern on the test day). Its expected shortlist is computed separately, from `/api/showtimes` with `lib/fit.ts`, so the run checks what the page renders.

Window 6:00–10:00 PM on 2026-10-17, travel 15, trailers 20, credits 0. Both runs, local dev server and production, gave the same result: 11 of 11 checks passed.

| Check | Result |
| --- | --- |
| Rendered shortlist equals the computed one | 9 shown, 9 expected |
| Ticket links go to `cinemark.com/TicketSeatMap/` | 9 of 9 |
| Trailers 0 lists the computed set | 11 |
| Settings persist across a reload | pass |
| Logging a start 22 minutes after "7:50 PM You Can See Everything" reports 22 minutes | pass |
| Median shows after one visit; "Use 22 min" sets trailers to 22 | pass |
| Visits persist across a reload; removing one clears the median | pass |
| No page errors | pass |

Screenshots of the new "At the theater" section in light and dark mode showed the time field's clock icon black on dark before `color-scheme` was declared, and light after.
