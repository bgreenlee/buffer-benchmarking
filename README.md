# Message Buffer Benchmarking

Benchmarking various data structures for message buffers. Each test fills a buffer with a number of messages (currently set to 1M), then reads them out.

Sample results from my laptop:

```
❯ go test -bench . -benchmem -benchtime 10s
goos: darwin
goarch: arm64
pkg: buffer-benchmarking
BenchmarkLinkedList-8                        154          74571560 ns/op        72000073 B/op    2000001 allocs/op
BenchmarkChannelBuffer-8                     434          27775905 ns/op        24002659 B/op          2 allocs/op
BenchmarkChannelBufferIfLen-8                433          27436656 ns/op        24002660 B/op          2 allocs/op
BenchmarkSlice-8                             339          34684213 ns/op        137727995 B/op        39 allocs/op
BenchmarkSliceMake-8                         388          30968511 ns/op        127860763 B/op        27 allocs/op
BenchmarkDeque-8                            1002          11974208 ns/op        27102012 B/op      47663 allocs/op
BenchmarkDeque2-8                            726          16510112 ns/op        75496709 B/op         33 allocs/op
PASS
ok      buffer-benchmarking     116.834s
```