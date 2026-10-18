---
kind: observation
title: Cinemark answers Vercel with 429 after a few dozen requests
date: 2026-10-17
source: scripts/conformance-live.ts against https://showtime-window.vercel.app (private showtime-window repo, commits f77f8ef and its follow-up), vercel logs for production, and curl from a home connection, 2026-10-17 between about 4:40 and 5:00 PM Eastern
projects: [showtime-window]
---

# Cinemark answers Vercel with 429 after a few dozen requests

| Run | Uncached Cinemark requests from Vercel | Result |
| --- | --- | --- |
| Unpaced, 43 Cinemark theaters × 2 dates, back to back | about 70 succeeded | Every later request returned 502 from the app. The production log shows `Error: cinemark 429 for pa-erie/cinemark-tinseltown-erie 2026-10-17` and the same for the others. |
| About 2 minutes later, after a new deployment: one request | 1 | 200 |
| Paced at 1.5 s per request, every theater again | about 25 that weren't already cached | 429 from `oh-milford` onward, 42 of 106 checks failed |
| Same theater pages from a home connection, during the 429s | 2 | 200 |

- The 429 applies to requests from Vercel's servers, not to the pages: the same URLs loaded from home.
- Pacing at 1.5 seconds didn't avoid it. The limit's window and threshold aren't known. Vercel's outbound addresses are shared, so other traffic may count toward it (inference).
- Earlier runs with 3 to 10 Cinemark theaters passed from Vercel ([2026-10-16](2026-10-16-showtime-window-first-conformance-run.md), [2026-10-17](2026-10-17-showtime-window-cleveland-cinemas-and-travel-limit-test-run.md)).
- The page's fetch loop re-requested every theater-day still loading each time one finished, so a page with many theaters sent many duplicate requests. Fixed in the same follow-up.

## Follow-up, 2026-10-18

The 21 theaters whose checks were refused were rerun through production in four groups of four or five theaters, 3 seconds per request and 2 minutes between groups, from 8:12 to 8:22 AM Eastern. Result: 42 of 42 passed, with no 429s. Each theater had showings on 2026-10-18 and 2026-10-22, and none was on another day.
