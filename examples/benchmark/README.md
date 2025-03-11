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

- Native JS methods (Sobek)
  - Marshal: 1442320 iters
  - Unmarshal: 1036531 iters
- Extension with encoding/json library
  - Marshal: 1264760 iters
  - Unmarshal: 2173293 iters
- Extension with goccy/go-json library
  - Marshal: 1426399 iters
  - Unmarshal: 2354553 iters
- Extension with bytedance/sonic library
  - Marshal: 1560107 iters
  - Unmarshal: 2744016 iters
