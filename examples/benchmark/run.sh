#!/bin/bash

DURATION=30s
MARSHAL_FILE=examples/benchmark/marshal.js
UNMARSHAL_FILE=examples/benchmark/unmarshal.js

ITERATIONS_MARSHAL_NATIVE=$(NATIVE=1 xk6 run --summary-export=summary.json $MARSHAL_FILE --duration=$DURATION > /dev/null && jq '.metrics.iterations.count' summary.json)
ITERATIONS_MARSHAL_EXTENSION=$(NATIVE=0 xk6 run --summary-export=summary.json $MARSHAL_FILE --duration=$DURATION > /dev/null && jq '.metrics.iterations.count' summary.json)
ITERATIONS_UNMARSHAL_NATIVE=$(NATIVE=1 xk6 run --summary-export=summary.json $UNMARSHAL_FILE --duration=$DURATION > /dev/null && jq '.metrics.iterations.count' summary.json)
ITERATIONS_UNMARSHAL_EXTENSION=$(NATIVE=0 xk6 run --summary-export=summary.json $UNMARSHAL_FILE --duration=$DURATION > /dev/null && jq '.metrics.iterations.count' summary.json)

echo "Marshal Native: $ITERATIONS_MARSHAL_NATIVE iters"
echo "Marshal Extension: $ITERATIONS_MARSHAL_EXTENSION iters"
echo "Unmarshal Native: $ITERATIONS_UNMARSHAL_NATIVE iters"
echo "Unmarshal Extension: $ITERATIONS_UNMARSHAL_EXTENSION iters"
