# Benchmarks

## Running the benchmarks

```bash
# Run from the repository root

# Native (JS) methods
NATIVE=1 xk6 run examples/benchmark/marshal.js --duration=10s

# Extension methods
NATIVE=0 xk6 run examples/benchmark/marshal.js --duration=10s


```

## Results

When running the k6 scripts for 30 seconds:

- encoding/json:
  - Marshal Native: 1442320 iters
  - Marshal Extension: 1264760 iters
  - Unmarshal Native: 1036531 iters
  - Unmarshal Extension: 2173293 iters
- goccy/go-json:
  - Marshal Native: 1438728 iters
  - Marshal Extension: 1426399 iters
  - Unmarshal Native: 1012748 iters
  - Unmarshal Extension: 2354553 iters
