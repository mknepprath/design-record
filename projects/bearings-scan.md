---
kind: project
title: Bearings Scan
status: active
owner: Michael Knepprath
date: 2026-10-16
---

# Bearings Scan

An iOS app: point the phone at a page of a book, and it identifies the book and your place in it, then opens a map of the book's settings at that point.

It reads a `bearings.json` file produced by an upstream extractor that belongs to a collaborator. At read time it uses no model, network or API key.

## Private documents

The repo is private; these links work for collaborators only.

- [README](https://github.com/mknepprath/bearings-scan/blob/master/README.md): how scanning finds your place, and measured accuracy
- [Reader plan](https://github.com/mknepprath/bearings-scan/blob/master/docs/READER_PLAN.md): milestones, invariants and conformance results
- [AGENTS.md](https://github.com/mknepprath/bearings-scan/blob/master/AGENTS.md): rules the code keeps

## Status

From the repo's git log, read 2026-10-16: 99 commits, 2026-09-02 to 2026-09-24. Xcode Cloud delivers every push to `master` to TestFlight internal testers.

Status is set by rule ([repository activity](../evidence/observations/2026-10-16-repository-activity.md)), not confirmed: `active` if the last non-bot commit is on or after 2026-08-16 (60 days before 2026-10-16), otherwise `paused`.
