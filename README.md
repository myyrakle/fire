# Fire: Automated Initialization in Go

- (fork or [google/wire](https://github.com/google/wire))

[![Build Status](https://github.com/google/fire/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/google/fire/actions)
[![godoc](https://godoc.org/github.com/google/fire?status.svg)][godoc]
[![Coverage](https://codecov.io/gh/google/fire/branch/master/graph/badge.svg)](https://codecov.io/gh/google/fire)

Fire is a code generation tool that automates connecting components using
[dependency injection][]. Dependencies between components are represented in
Fire as function parameters, encouraging explicit initialization instead of
global variables. Because Fire operates without runtime state or reflection,
code written to be used with Fire is useful even for hand-written
initialization.

For an overview, see the [introductory blog post][].

[dependency injection]: https://en.wikipedia.org/wiki/Dependency_injection
[introductory blog post]: https://blog.golang.org/fire
[godoc]: https://godoc.org/github.com/google/fire
[travis]: https://travis-ci.com/google/fire

## Installing

Install Fire by running:

```shell
go install github.com/google/fire/cmd/fire@latest
```

and ensuring that `$GOPATH/bin` is added to your `$PATH`.

## Documentation

- [Tutorial][]
- [User Guide][]
- [Best Practices][]
- [FAQ][]

[Tutorial]: ./_tutorial/README.md
[Best Practices]: ./docs/best-practices.md
[FAQ]: ./docs/faq.md
[User Guide]: ./docs/guide.md

## Project status

As of version v0.3.0, Fire is _beta_ and is considered feature complete. It
works well for the tasks it was designed to perform, and we prefer to keep it
as simple as possible.

We'll not be accepting new features at this time, but will gladly accept bug
reports and fixes.

## Community

For questions, please use [GitHub Discussions](https://github.com/google/fire/discussions).

This project is covered by the Go [Code of Conduct][].

[Code of Conduct]: ./CODE_OF_CONDUCT.md
[go-cloud mailing list]: https://groups.google.com/forum/#!forum/go-cloud
