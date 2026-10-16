---
kind: adr
number: 0001
title: Build a web app first
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [showtime-window]
evidence:
  - evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md
---

# ADR 0001: Build a web app first

## Context

The first plan assumed an iOS app, because the owner's other app work is in SwiftUI. The owner said "a web app is fine to start" ([owner's constraints](../evidence/feedback/2026-10-16-owner-constraints-for-showtime-window.md), quote 3).

## Decision

Build the first version as a static web app. A native app is out of scope until the web version is in use.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| Web app | No App Store; free static hosting; works on any phone | No widgets or Live Activities |
| iOS app | Widgets and Live Activities ("leave by 6:40") | Slower to ship; an Apple developer account costs money |

## Consequences

- All showtime fetching happens before the page loads ([ADR 0002](0002-prebuild-showtimes-as-static-json.md)), since theater sites can't be called from the browser.
- Settings live in `localStorage` until there is a reason for accounts.
