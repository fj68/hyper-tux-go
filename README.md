# hyper-tux-go

## Development

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
