# hyper-tux-go

## Description

Hyper Tux is a small puzzle-like game played on a 16×16 grid. Players control colored actors by swiping (mouse or touch) to move them across the board and aim to reach target goal cells with the minimum movements. The game uses Ebiten for cross-platform rendering and supports desktop and WebAssembly builds for now.

The project includes utilities for random placement, movement history (undo/redo), and snapshot-based visual regression tests used during development.

### Key features

- Random placement of actors and goals
- Supports swipe by mouse and one-finger touch input
- Undo / Redo with move history visualization
- Ebiten-based crossplatform rendering
- Snapshot-based visual regression tests for UI stability

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

### Release

Assuming on a GitHub Codespaces (Linux).

#### Linux

```console
go build -o bin/main.out
```

#### Windows

```Console
GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o bin/main.exe
```
