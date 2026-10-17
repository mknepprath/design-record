---
kind: project
title: lilt
status: active
owner: Michael Knepprath
date: 2026-10-16
---

# lilt

A text adventure game, playable at [mknepprath.com/lilt](https://mknepprath.com/lilt) and on Mastodon as [@familiarlilt](https://mastodon.social/@familiarlilt). Started on Twitter in 2012.

## How it works

A stateless engine, `play(move, state) → {response, state}`, behind a Flask API on Railway. Clients pass the state on every request. The Mastodon adapter keeps each player's state. Game data (moves, items) is JSON.

## Code

- [mknepprath/lilt](https://github.com/mknepprath/lilt) (public): engine, API, Mastodon bot, CLI
- [mknepprath/lilt-web](https://github.com/mknepprath/lilt-web) (public): web client, deployed at lilt-web.vercel.app

## Status

Written 2026-10-16 from the repo's README and git log. First commit 2012-07-24, last non-bot commit 2026-10-03, 1,330 commits.

Status `active` confirmed by the owner on 2026-10-16 ([owner's confirmation](../evidence/feedback/2026-10-16-owner-confirms-project-statuses.md)). Commit history: [repository activity](../evidence/observations/2026-10-16-repository-activity.md).
