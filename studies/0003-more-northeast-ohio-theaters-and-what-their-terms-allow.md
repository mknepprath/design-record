---
kind: study
number: 0003
title: More Northeast Ohio theaters, and what their terms allow
status: running
owner: Michael Knepprath
date: 2026-10-18
projects: [showtime-window]
evidence:
  - evidence/observations/2026-10-18-fandango-schedule-endpoints.md
  - evidence/observations/2026-10-18-theater-site-terms-on-automated-access.md
  - evidence/observations/2026-10-16-cleveland-area-theater-access.md
---

# Study 0003: More Northeast Ohio theaters, and what their terms allow

Blocked: the owner has to decide how the project treats sites whose terms prohibit automated collection.

## Question

Which theaters near Valley View aren't covered yet, which sources could read them at $0, and what do those sources' terms allow? This continues [Study 0002](0002-which-cleveland-area-theaters-can-be-read-for-free.md).

## Method

1. Web searches for the Capitol Theatre, Atlas Eastgate, independent theaters and Veezi theaters in Ohio.
2. curl probes of independent theater sites; headless Chromium to record the JSON requests of Fandango's and the Nightlight's pages.
3. Fandango's schedule endpoints called with curl and from a temporary route on a Vercel preview ([endpoints](../evidence/observations/2026-10-18-fandango-schedule-endpoints.md)).
4. Read the terms of Fandango, Cinemark and Cleveland Cinemas, and looked for terms on the Atlas and Veezi pages ([terms](../evidence/observations/2026-10-18-theater-site-terms-on-automated-access.md)).

One pass, 2026-10-18 morning. Fandango was called about 25 times, then no more once its terms were read.

## Findings

- **Fandango** covers 37 Northeast Ohio theaters, 22 of them not in the app, including 8 Regal and 2 AMC. Its data matched cinemark.com and Veezi time for time, and it answered from Vercel. On access, completeness and coverage it is the best source found. Its terms prohibit automated collection without written authorization.
- **The live sources:** Cinemark's and Cleveland Cinemas' terms also prohibit robots and data gathering. No terms were found for Atlas or Veezi.
- **Atlas Eastgate 10** closed on 2024-07-14 ([News 5 Cleveland](https://news5cleveland.com/news/local-news/atlas-cinemas-at-eastgate-shopping-center-closing-down), [Cleveland Jewish News](https://www.clevelandjewishnews.com/news/local_news/atlas-cinemas-to-close-in-mayfield-heights-dave-and-buster-s-set-to-fill-space/article_efe5c3a2-3fa2-11ef-bb7e-6f3b38c7dd1b.html); secondary).
- **Capitol Theatre:** its own showtime site still wasn't found. It is on Fandango as an independent (`INDY`).
- **Independents with their own sites:** the Nightlight (Akron) loads showtimes from its own `/graphql`. Mayfield Road Drive-In and Midway Twin Drive-In sell through SimpleTix. Canton Palace links Eventbrite. None of these sites' terms was read, and none was assessed further.
- **Official APIs:** AMC runs a developer API that needs an approved vendor key (recalled, not checked). Nothing comparable was found for Regal.

## Options for the owner

| Option | For | Against |
| --- | --- | --- |
| Keep the current sources and stay off Fandango | No change; personal, low-volume use with an hour's cache | The current Cinemark and Cleveland Cinemas use is already contrary to their terms; the app is public, so it isn't only personal use |
| Ask for written permission (Cinemark, Cleveland Cinemas, Fandango) | The terms allow it with permission | Answers may never come; Fandango likely licenses data commercially (guess) |
| Only sources whose terms allow it: Atlas via Veezi, official APIs (AMC, if a key is granted), theaters that agree | Clean | Drops 48 of 53 theaters today |
| Keep the sources but make the app private to the owner | Closer to personal use | The terms make no exception for personal use |

Strongest case for keeping the sources: the data is published as schema.org markup and public JSON for any visitor, the request volume is close to one person browsing, and `robots.txt` allows the pages. Strongest case against: the terms are explicit, and two of the three operators have already said no in writing.

## Open

- The owner's decision on the options above.
- Read the terms of every further source before building it.
