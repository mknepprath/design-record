---
kind: project
title: Duck Pond
status: dropped
owner: Michael Knepprath
date: 2026-10-16
---

# Duck Pond

A personal kanban board where the prompts are the product. Each card is backed by an Obsidian note. Dispatch compiles a self-contained prompt that tells an agent what to do and how to update the card when it's done.

## How it works

- Cards are Obsidian notes in a Duck Pond folder, with YAML frontmatter holding status.
- Backlog candidates are pulled live from Todoist and GitHub, ranked and deduplicated, and promoted to a note when chosen.
- Dispatch puts on the clipboard: the task instructions, the card's note path and status schema, and completion instructions.
- Native macOS app, local-first. Phase 1 (board, live file sync, drag and drop, templates) is done.

(From Claude Code's memory notes.)

## Private documents

The repo is private; these links work for collaborators only.

- [duck-pond repo](https://github.com/mknepprath/duck-pond)

## Status

Written 2026-10-16 from Claude Code's memory notes. Last non-bot commit: 2026-04-14 (1 commit).

Status `dropped` confirmed by the owner on 2026-10-16 ([owner's confirmation](../evidence/feedback/2026-10-16-owner-confirms-project-statuses.md)). Commit history: [repository activity](../evidence/observations/2026-10-16-repository-activity.md).
