# hyper-tux-go

## Installation

```console
git clone git@github.com/fj68/hyper-tux-go
```

## Run

```console
go run main.go
```

Or, to make Web version:

```console
go run github.com/eihigh/wasmnow@latest -b -d docs
```

## Development

Recommended to use GitHub Codespaces online or Docker on a local machine using [Dockerfile in this repository](Dockerfile).

Following instructions assume that you are using GitHub Codespaces.

### Unit tests

```console
go test
```

### Visual regression tests

```console
go test -tags=guitests
```

When some errors happen, try to run Xvfb by yourself:

```console
Xvfb :99 -screen 0 1024x768x24 &
```

To update expected snapshots by replacing them with current ones:

```console
UPDATE_SHAPSHOT=1 go test -tags=guitests
```

Note: Currently `TestGameState` always fails because it places actors at random and thus the result differs on each tests.
