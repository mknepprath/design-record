---
kind: project
title: lily dex
status: active
owner: Michael Knepprath
date: 2026-10-16
---

# lily dex

A Pokémon GO companion app for iOS. It goes beyond a checklist: it says what to catch next, and gives one-tap search strings and widgets to act on that in the game.

## Goal

"Flighty for Pokémon GO": a polished, data-rich companion that brings up the right information at the right time. (Owner's vision, from Claude Code's memory notes.)

## What it covers

- Smart "next to catch" suggestions
- PvP rankings with one-tap search strings
- Regional filtering to what's obtainable nearby
- Quick Copy widgets for common in-game searches
- Live event integration and countdowns
- Form tracking across all variants

## Parts

| Part | Repo | Visibility |
| --- | --- | --- |
| iOS app (SwiftUI, widgets, Live Activity) | `lily-dex` | Private |
| Data API: a scheduled Node build that publishes static JSON to GitHub Pages every 6 hours | [lily-dex-api](https://github.com/mknepprath/lily-dex-api) | Public |
| Microsite | [lily-dex-web](https://github.com/mknepprath/lily-dex-web), live at [lilydex.com](https://lilydex.com) | Public |

## Public documents

- [lily-dex-api README](https://github.com/mknepprath/lily-dex-api/blob/main/README.md): data sources and endpoints
- [lily-dex-api invariants](https://github.com/mknepprath/lily-dex-api/blob/main/docs/INVARIANTS.md): rules the data pipeline keeps

## Private documents

The repo is private; these links work for collaborators only.

- [lily-dex repo](https://github.com/mknepprath/lily-dex)

## Status

Written 2026-10-16 from Claude Code's memory notes. Last non-bot commit: lily-dex 2026-10-03 (225 commits), lily-dex-api 2026-10-03 (260), lily-dex-web 2026-08-28 (1,620).

Status is set by rule, not confirmed: `active` if the last non-bot commit is on or after 2026-08-16 (60 days before 2026-10-16), otherwise `paused`. Commits made by Bishop jobs count as activity.
