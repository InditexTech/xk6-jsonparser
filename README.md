# xk6-jsonparser

An extension for the k6 load testing tool that allows JSON marshal & unmarshal in Go, using the [bytedance/sonic library](https://github.com/bytedance/sonic). This provides better performance than the native JavaScript JSON methods using the underlying Sobek implementation (see the [benchmarks section](examples/benchmark/README.md) for more details).

## Install


To build a `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

Then:

1. Download [xk6](https://github.com/grafana/xk6):
```bash
go install go.k6.io/xk6/cmd/xk6@latest
```

2. [Build](https://github.com/grafana/xk6#command-usage) the k6 binary:
```bash
xk6 build --with github.com/InditexTech/xk6-jsonparser@latest
```

Alternatively, you can build a `k6` binary with the extension from the local code, rather than from GitHub:

```bash
make build
```

### Development

The default target in the [Makefile](Makefile) will download the dependencies, format your code, run the tests and the example.

```bash
git clone git@github.com:InditexTech/xk6-jsonparser.git
cd xk6-jsonparser
make
```

## Usage

This extension provides two methods:

```javascript
import json from "k6/x/json";

export default function() {
  // Convert a JS object to a JSON string.
  // If the operation fails, the output is an empty string.
  const marshalResult = json.marshal({userId: 1, userName: "Lorem ipsum"});
  
  // Convert a JSON string to a JS object.
  // If the operation fails, the output is a null value.
  const unmarshalResult = json.unmarshal('{"userId": 2, "userName": "Dolor sit amet"}');
}
```

See the [examples](./examples) folder for a more detailed usage example.

## Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](CONTRIBUTING.md) file for guidelines on how to contribute to this project.

## License

This project is licensed under the terms of the [AGPL-3.0-only](LICENSE) license.

© 2025 INDUSTRIA DE DISEÑO TEXTIL S.A. (INDITEX S.A.)
