## Go Benchmark

benchmark executed by the "go test" command when its -bench flag is provided. Benchmarks are run sequentially.

```go
func BenchmarkRandInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Int()
    }
}
```

> go test -run=TestFib -v

> go test -run=xxx -v -bench=.

> go test -run=xxx -v -bench=. -benchtime=3s

### With Memory test
> go test -benchmem -run=^$ -bench ^BenchmarkSs* graphmysql/utility

> go test -benchmem -bench ^BenchmarkSs* graphmysql/utility

### No testing only benchmark checking
> go test -run=^$ -benchmem -bench ^BenchmarkGetpass* graphmysql/utility

### Withour memory
> go test -bench ^BenchmarkSs* graphmysql/utility
