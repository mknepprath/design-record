---
kind: observation
title: Who Goes There? deployment guide
date: 2026-10-16
source: mknepprath-next VERCEL-DEPLOYMENT.md at commit ff85ced (public repo), read 2026-10-16
projects: [who-goes-there]
---

# Who Goes There? deployment guide

From https://github.com/mknepprath/mknepprath-next/blob/master/VERCEL-DEPLOYMENT.md, quoted:

> Since Vercel doesn't support WebSockets, we split the deployment:
> - **Vercel**: Hosts the Next.js frontend
> - **Railway**: Hosts the Socket.IO game server

The guide then copies `lib/game-server.js` and a standalone server into a separate repo and deploys that repo to Railway. `wgt-server` is that repo.

Whether Vercel supports WebSockets today wasn't checked; this records what the guide says.
