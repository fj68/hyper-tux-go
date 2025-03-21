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

To update expected snapshots by replacing them with current ones:

```console
UPDATE_SHAPSHOT=1 go test -tags=guitests
```
