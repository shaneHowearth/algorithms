# Classic sorting algorithm examples


| Algorithm | Worst Case | Average Case | Best Case | Worst Case Space Complexity | Notes |
|:---|:---:|:---:|:---:|:---|:---|
| Built in ||||| Go's built in sorting methods. |
| Bubble Sort | O(n²) | O(n²) | O(n), list is already sorted. |||
| Insertion Sort | O(n²) | O(n²) | O(n), list is already sorted. | O(n) total, O(1) auxiliary. ||
| Heap Sort | O(n log n) | O(n log n) | O(n log n), distinct keys, O(n) equal keys.| O(n) total, O(1) auxiliary. ||
| Merge Sort | O(n log n) | Θ(n log n) | Ω(n log n), typical, Ω(n) | O(n) total, O(n) auxiliary (O(1) auxiliary with linked lists). ||
| Quicksort | O(n²) | O(n log n) | O(n log n), simple partition, O(n) three way partition and equal keys. | O(n) auxiliary , (O(log n) Hoare, 1962). ||

## Benchmarks
``` shell
$ go test -bench . ./... -timeout 0 -benchmem
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/sorting/bubble
BenchmarkSort-4            10000         105230713 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/shanehowearth/algorithms/sorting/bubble  1054.907s
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/sorting/builtin
BenchmarkSort-4            10000           1256226 ns/op              32 B/op          1 allocs/op
BenchmarkCustom-4          10000           1122302 ns/op              64 B/op          2 allocs/op
PASS
ok      github.com/shanehowearth/algorithms/sorting/builtin 28.577s
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/sorting/heap
BenchmarkSort-4            10000           1208035 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/shanehowearth/algorithms/sorting/heap    14.635s
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/sorting/insertion
BenchmarkSort-4            10000          14107052 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/shanehowearth/algorithms/sorting/insertion       143.581s
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/sorting/merge
BenchmarkSort-4            10000           1285258 ns/op         1125253 B/op      19998 allocs/op
PASS
ok      github.com/shanehowearth/algorithms/sorting/merge   15.358s
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/sorting/quick
BenchmarkSort-4            10000            580194 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/shanehowearth/algorithms/sorting/quick   8.218s
```
