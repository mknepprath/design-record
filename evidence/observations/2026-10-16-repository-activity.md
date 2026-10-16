---
kind: observation
title: Repository activity across projects
date: 2026-10-16
source: git log and gh repo view for each repo under ~/Developer, run 2026-10-16
projects: [bearings-scan, bishop, bout, duck-pond, every-pokemon-card, game-of-life, lilt, lily-dex, llm-style, memory-map, mknepprath-next, oracle, pokemon-facts, pokemon-prestige-red, robot-mk, stubpedia, tardy-critic, test-game-0, tiny-mystery-club, who-goes-there, ytcast]
---

# Repository activity across projects

Command per repo: the first commit date, the most recent commit not matching `dependabot`, `bump ` or `update … to <version>`, `git rev-list --count HEAD`, and visibility from `gh repo view`.

| Project | Repo | Visibility | First commit | Last non-bot commit | Commits |
| --- | --- | --- | --- | --- | --- |
| bearings-scan | `bearings-scan` | private | 2026-09-02 | 2026-09-24 | 99 |
| bishop | `bishop` | private | 2026-05-26 | 2026-10-04 | 112 |
| bishop | `BishopCockpit` | private | 2026-08-22 | 2026-09-10 | 2 |
| bout | `bout` | public | 2017-03-20 | 2026-10-03 | 78 |
| duck-pond | `duck-pond` | private | 2026-04-14 | 2026-04-14 | 1 |
| every-pokemon-card | `every-pokemon-card` | public | 2022-03-21 | 2026-10-04 | 9 |
| game-of-life | `gol-rpg` | no remote | 2026-08-15 | 2026-10-16 | 23 |
| lilt | `lilt` | public | 2012-07-24 | 2026-10-03 | 1330 |
| lilt | `lilt-web` | public | 2024-11-12 | 2024-11-12 | 3 |
| lily-dex | `lily-dex` | private | 2020-07-25 | 2026-10-03 | 225 |
| lily-dex | `lily-dex-api` | public | 2026-03-12 | 2026-10-03 | 260 |
| lily-dex | `lily-dex-web` | public | 2021-03-20 | 2026-08-28 | 1620 |
| llm-style | `llm-style` | private | 2026-09-24 | 2026-10-04 | 5 |
| memory-map | `memory-map` | no remote | 2026-03-29 | 2026-03-30 | 12 |
| mknepprath-next | `mknepprath-next` | public | 2018-11-07 | 2026-10-16 | 4504 |
| oracle | `oracle` | no remote | 2026-04-07 | 2026-04-07 | 4 |
| pokemon-facts | `pokemon-facts` | public | 2022-03-21 | 2026-09-15 | 17 |
| pokemon-prestige-red | `pokemon-prestige-red` | private | 2017-12-16 | 2026-10-16 | 3956 |
| robot-mk | `robot-mk` | public | 2013-08-09 | 2026-10-06 | 184 |
| stubpedia | `stubpedia` | no remote | 2026-08-04 | 2026-08-04 | 1 |
| tardy-critic | `tardy-critic` | public | 2023-04-03 | 2026-03-11 | 31 |
| test-game-0 | `test-game-0` | public | 2026-03-25 | 2026-03-26 | 9 |
| tiny-mystery-club | `tiny-mystery-club` | public | 2017-11-19 | 2026-08-25 | 2161 |
| who-goes-there | `wgt-server` | public | 2026-03-06 | 2026-09-19 | 18 |
| ytcast | `ytcast` | private | 2026-08-19 | 2026-09-16 | 6 |

Caveats:

- Commits dated 2026-10-16 in `gol-rpg`, `mknepprath-next` and `pokemon-prestige-red` were made by Claude Code while moving documents into this record. Before that, `gol-rpg`'s last commit was 2026-08-16.
- Commits made by Bishop jobs aren't filtered out, so a repo can look active because of automation.
