# Classic sorting algorithm examples


| Algorithm | Worst Case | Average Case | Best Case | Worst Case Space Complexity | Done |
|:---:|:---:|:---:|:---:|:---|:---|
| Bubble Sort | O(n²) | O(n²) | O(n), list is already sorted. |||
| Insertion Sort | O(n²) | O(n²) | O(n), list is already sorted. | O(n) total, O(1) auxiliary. ||
| Heap Sort | O(n log n) | O(n log n) | O(n log n), distinct keys, O(n) equal keys.| O(n) total, O(1) auxiliary. ||
| Merge Sort | O(n log n) | Θ(n log n) | Ω(n log n), typical, Ω(n) | O(n) total, O(n) auxiliary (O(1) auxiliary with linked lists). ||
| Quicksort | O(n²) | O(n log n) | O(n log n), simple partition, O(n) three way partition and equal keys. | O(n) auxiliary , (O(log n) Hoare, 1962). ||

## Benchmarks
Note: the benchmarks shown below suggest a problem with the benchmarking tests
``` shell
$ go test -bench . ./...
rand seeded
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/gga/bubble
BenchmarkSort-4            10000          15430408 ns/op
PASS
ok      github.com/shanehowearth/algorithms/gga/bubble  154.307s
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/gga/heap
BenchmarkSort-4            10000            420515 ns/op
PASS
ok      github.com/shanehowearth/algorithms/gga/heap    4.208s
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/gga/insertion
BenchmarkSort-4           141862            124053 ns/op
PASS
ok      github.com/shanehowearth/algorithms/gga/insertion       17.686s
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/gga/merge
BenchmarkSort-4            10000            400592 ns/op
PASS
ok      github.com/shanehowearth/algorithms/gga/merge   4.016s
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/gga/quick
BenchmarkSort-4            10000          10303087 ns/op
PASS
ok      github.com/shanehowearth/algorithms/gga/quick   103.035s
```
