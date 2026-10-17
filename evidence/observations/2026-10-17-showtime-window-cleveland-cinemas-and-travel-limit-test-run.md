---
kind: observation
title: Cleveland Cinemas and the travel limit pass from Vercel
date: 2026-10-17
source: npm test, scripts/conformance-live.ts and npm run e2e in the private showtime-window repo at commit 6a08c03, run locally and against https://showtime-window.vercel.app, 2026-10-17 about 9:45 AM Eastern
projects: [showtime-window]
---

# Cleveland Cinemas and the travel limit pass from Vercel

- Unit tests: 43 of 43 pass. The new ones cover the box office adapter (T1, T2, T3, T6 on a saved two-week Cedar Lee response of 114 showtimes) and the travel limit.
- Live T4 and T5, through the deployed `/api/showtimes`: 20 of 20, ten theaters × two dates. Cedar Lee: 21 showings on 2026-10-17, 10 on 2026-10-22. Apollo: 6 and 3.
- Browser test against production: 17 of 17. From a location near Valley View with a 30-minute limit, the estimate put three theaters in range: Cinemark Valley View, Atlas Shaker Square and Cedar Lee. The list read "20 fit between 6:00 PM and 10:00 PM across 3 theaters", and every card named one of those three.
- The first local run of the live check failed for Atlas Diamond Center and Lakeshore on Friday 2026-10-23, with 0 showings at each. Both list every day through Thursday 2026-10-22, then only scattered dates (2026-10-29, 2026-11-05, December). The other three Atlas theaters list 2026-10-23. Reading (inference): those two hadn't posted the release week starting that Friday. The check now uses the coming Thursday as its far date.
