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

When running the k6 scripts for 10 seconds:

- Marshal with Native JS: 401373 iterations
- Marshal with Extension: 332099 iterations
- Unmarshal with Native JS: 289598 iterations
- Unmarshal with Extension: 597902 iterations
