---
kind: observation
title: Fandango, Cinemark and Cleveland Cinemas terms prohibit automated collection
date: 2026-10-18
source: https://www.fandango.com/policies/terms-of-use, https://www.cinemark.com/terms-conditions, https://www.clevelandcinemas.com/terms/, the Atlas Cinemas home page and a Veezi sessions page, fetched with curl 2026-10-18 about 8:25 AM Eastern
projects: [showtime-window]
---

# Fandango, Cinemark and Cleveland Cinemas terms prohibit automated collection

| Site | Clause, quoted | Bears on |
| --- | --- | --- |
| Fandango | "you will not collect data from, or use any automated method (including, but not limited to robots, scripts, spiders, data extractors or other automated means) to collect data from, access or search, the Services, unless Fandango has separately given its express written authorization to do so" | The candidate source in [the endpoints observation](2026-10-18-fandango-schedule-endpoints.md) |
| Cinemark | "Systematic retrieval of content from the Service to create or compile, directly or indirectly, a collection, compilation, database or directory (whether through robots, spiders, automatic devices or manual processes) or creating links to the Service without written permission from Cinemark is strictly prohibited." | The live Cinemark adapter (46 theaters) |
| Cleveland Cinemas | Visitors agree not "To use any data mining, robots, or similar data gathering or extraction methods in connection with this Website" | The live box office adapter (Cedar Lee, Apollo) |
| Atlas Cinemas, Veezi | No terms link on atlascinemas.net or on the sessions page; `veezi.com/terms`, `/legal` and `/privacy-policy` returned 404 | The live Veezi adapter (five theaters). Terms weren't found, which isn't the same as permission. |

- Cinemark's `robots.txt` doesn't disallow `/theatres/`, and its pages set `robots: index, follow`. That governs crawlers' indexing, not the terms above.
- Study 0001 and Study 0002 scored "terms" without reading these pages: 2 (neutral) for Cinemark, Atlas and the Cinematheque, and 1 for Cleveland Cinemas. That was an error in method, corrected in [Study 0003](../../studies/0003-more-northeast-ohio-theaters-and-what-their-terms-allow.md).
- Reading the clauses' legal effect is outside this record. This is not legal advice.
