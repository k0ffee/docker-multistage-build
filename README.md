Examples for small Go containers
================================

These example Docker containers separate building and running
a Go application into two containers of which only the second
will be deployed afterwards.

Dockerfiles
-----------

`Dockerfile.v1-scratch-multi`:
standalone binary file without additional GNU/Linux userland.

`Dockerfile.v2-alpine-multi`:
binary file embedded into minimal Alpine distribution.

`Dockerfile.v3-busybox-multi`:
standalone binary file embedded alongside the busybox tool.

`Dockerfile.v4-scratch-multi`:
standalone binary file without additional GNU/Linux userland.
Statically linked using Cgo.

`Dockerfile.v5-busybox-multi`:
standalone binary file embedded alongside the busybox tool.
Statically linked using Cgo.

`Dockerfile`:
using a Makefile for flexible build options.

Design
------

- Use "make" everywhere.
- Have Go cache in a separate directory to avoid name conflicts.
- Run as non-root user everywhere.
- Prefer the smaller scratch image, scratch containers can be
  debugged via fully featured sidecar container sharing its
  namespace.
