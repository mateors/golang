## Go Benchmark

benchmark executed by the "go test" command when its -bench flag is provided. Benchmarks are run sequentially.

```go
func BenchmarkRandInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Int()
    }
}
```
