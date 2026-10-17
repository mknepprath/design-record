---
kind: project
title: Memory Map
status: paused
owner: Michael Knepprath
date: 2026-10-16
---

# Memory Map

A 3D city built from an Obsidian vault, in a Wind Waker style: each note is a building. You walk or fly through it. Vite and Three.js; it started as an Obsidian plugin and was moved out after the plugin kept freezing.

## How it works

An export script turns the vault into `notes.json`. A deterministic layout places blocks, streets, plazas and parks, and buildings render instanced with spatial chunking (a 50×50 grid with a 150-unit active radius). (From Claude Code's memory notes.)

## Ideas not started

- Traffic along streets between notes connected by wikilinks
- Real street grids from OpenStreetMap as the base layer
- A minimap, click-to-open notes, search-to-fly, and a day/night cycle

## Code

Local only: `~/Developer/memory-map`, no remote.

## Status

Written 2026-10-16 from Claude Code's memory notes. Last non-bot commit: 2026-03-30 (12 commits).

Status `paused` confirmed by the owner on 2026-10-16 ([owner's confirmation](../evidence/feedback/2026-10-16-owner-confirms-project-statuses.md)). Commit history: [repository activity](../evidence/observations/2026-10-16-repository-activity.md).
