# FogBugz JSON API for Go

This was written some time in 2017, but never got committed and pushed. This
is, therefore, presented as-is.

It's uh, missing tests. For shame.

See `main.go` to see how it's used. You'll want to set `FOGBOT_HOSTNAME` to
something like `fogbugz.yourdomain.com` for on-premise solutions, or
`mycompany` (or `mycompany.fogbugz.com`) for hosted.

You then need either a username-password pair (`FOGBOT_USERNAME` and
`FOGBOT_PASSWORD` env-vars) or a token (`FOGBOT_TOKEN`).


