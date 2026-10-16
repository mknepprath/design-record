---
kind: project
title: Game of Life
status: active
owner: Michael Knepprath
date: 2026-10-16
---

# Game of Life

A turn-based roguelike where your balance sheet is your character sheet. An emergency fund is a shield, debt is poison, income is mana regen and insurance blocks damage. Fully simulated: no real financial data.

## How it works

One turn is a month, and a run is 120 turns unless you go broke first. Each turn runs income, upkeep, your action, an encounter, interest, the market, then settlement. You only make decisions in the action step. Full mechanics: [Spec 0002](../specs/0002-game-of-life-design.md).

## Code

Local only: `~/Developer/gol-rpg`, no remote. `npm run play` to play, `npm test` for the engine tests.

## Status

Written 2026-10-16 from the repo's README and git log. First commit 2026-08-15, last non-bot commit 2026-08-16, 22 commits.

Status is set by rule, not confirmed: `active` if the last non-bot commit is on or after 2026-08-16 (60 days before 2026-10-16), otherwise `paused`. Commits made by Bishop jobs count as activity.
