---
kind: adr
number: 0004
title: Run the Who Goes There? game server on Railway, apart from the site
status: accepted
owner: Michael Knepprath
date: 2026-10-16
projects: [who-goes-there]
evidence:
  - evidence/observations/2026-10-16-who-goes-there-deployment-guide.md
---

# ADR 0004: Run the Who Goes There? game server on Railway, apart from the site

Recorded 2026-10-16 after the fact, from the sources cited. The decision itself was made earlier, when the code was written.

## Context

The game needs a long-lived Socket.IO connection per player. The site runs on Vercel, and the deployment guide states "Vercel doesn't support WebSockets" ([guide](../evidence/observations/2026-10-16-who-goes-there-deployment-guide.md)).

## Decision

Host the frontend in mknepprath-next on Vercel and the Socket.IO game server as a separate repo, `wgt-server`, on Railway. The frontend finds the server through `NEXT_PUBLIC_GAME_SERVER_URL`.

## Options considered

| Option | For | Against |
| --- | --- | --- |
| Separate server on Railway | WebSockets work; site deploys are unaffected | Two deployments; cross-origin connection |
| Everything on Vercel | One deployment | Ruled out by the guide's WebSocket statement |

Other options weren't recorded.

## Consequences

- Game logic was copied from mknepprath-next into `wgt-server`; the two copies can drift.
- The game depends on Railway staying up and on its pricing.
