---
kind: observation
title: lily-dex-api README
date: 2026-10-16
source: lily-dex-api README.md at commit e3648c1 (public repo), read 2026-10-16
projects: [lily-dex]
---

# lily-dex-api README

From https://github.com/mknepprath/lily-dex-api/blob/main/README.md, quoted:

> Pokemon GO data API for lily dex. Aggregates data from multiple sources into a unified JSON API, deployed to GitHub Pages on a 6-hour build cycle.

> GitHub Actions runs `node src/index.js` on push to main, every 6 hours, or manual dispatch. Output files are deployed to GitHub Pages. Source fetch results are cached in `cache/` and committed back so builds can fall back to cached data if an upstream source is temporarily unavailable.

> Keeps timestamps as **naive strings** (no timezone) — Pokemon GO events happen at local time, so the app interprets them in the user's local timezone

> Requires Node.js >= 20. No runtime dependencies.
