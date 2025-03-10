# xk6-jsonparser

The `xk6-jsonparser` is a plugin for the k6 load testing tool that allows JSON marshal & unmarshal in Go.

## Install

### Pre-built binaries

```bash
make run
```

### Build from source

```bash
make build
```

## API

The plugin provides two methods:

- `marshal`: convert a JavaScript object to a JSON string. When the operation fails, the output is an empty string.
- `unmarshal`: convert a JSON string to a JavaScript object. When the operation fails, the output is a null value.

### Examples

See the [examples](./examples) folder.
