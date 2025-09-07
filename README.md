rosary
---

A program to say the rosary over telnet-ish.

Configured via env vars in the Dockerfile that are fairly self explanatory.

Runs by default on port 6724.

You can run it via docker:

```
docker run --rm -p 0.0.0.0:6724:6724 ghcr.io/packrat386/rosary:latest
```
