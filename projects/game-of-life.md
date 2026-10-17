---
kind: project
title: Game of Life
status: done
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

Status `done` confirmed by the owner on 2026-10-16 ([owner's confirmation](../evidence/feedback/2026-10-16-owner-confirms-project-statuses.md)). Commit history: [repository activity](../evidence/observations/2026-10-16-repository-activity.md).
