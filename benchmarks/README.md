# Benchmarks

To run a benchmark like `compress_test.go`, do:

```bash
go test -bench=. -benchtime=10x ./benchmarks/compress_test.go
```

If you want output that's not always in `ns/op`, use `benchstat`:

```bash
go get golang.org/x/perf/cmd/benchstat
go install  golang.org/x/perf/cmd/benchstat

# Then append `| benchstat -` to your command like:
go test -bench=. -benchtime=10x ./benchmarks/compress_test.go | benchstat -
```

## `compress_test.go`

This benchmarks various approaches to compression to see which one is faster.

```bash
go test -bench=. -benchtime=10x ./benchmarks/compress_test.go | benchstat -

goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
                        │      -      │
                        │   sec/op    │
Compress-8                1.154 ± ∞ ¹
CompressFromReader-8      1.255 ± ∞ ¹
ReuseGifsicleCompress-8   1.324 ± ∞ ¹
geomean                   1.243
```
