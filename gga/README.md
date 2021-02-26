# Classic sorting algorithm examples


| Algorithm | Worst Case | Average Case | Best Case | Worst Case Space Complexity | Done |
|:---:|:---:|:---:|:---:|:---|:---|
| Bubble Sort | O(n²) | O(n²) | O(n), list is already sorted. |||
| Insertion Sort | O(n²) | O(n²) | O(n), list is already sorted. | O(n) total, O(1) auxiliary. | TODO |
| Heap Sort | O(n log n) | O(n log n) | O(n log n), distinct keys, O(n) equal keys.| O(n) total, O(1) auxiliary. | TODO |
| Merge Sort | O(n log n) | Θ(n log n) | Ω(n log n), typical, Ω(n) | O(n) total, O(n) auxiliary (O(1) auxiliary with linked lists). ||
| Quicksort | O(n²) | O(n log n) | O(n log n), simple partition, O(n) three way partition and equal keys. | O(n) auxiliary , (O(log n) Hoare, 1962). ||

## Benchmarks
``` shell
$ go test -v -bench . ./...
rand seeded
=== RUN   TestBubble
=== RUN   TestBubble/Book_example
=== RUN   TestBubble/Zero_length
=== RUN   TestBubble/Already_sorted
--- PASS: TestBubble (0.00s)
    --- PASS: TestBubble/Book_example (0.00s)
    --- PASS: TestBubble/Zero_length (0.00s)
    --- PASS: TestBubble/Already_sorted (0.00s)
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/gga/bubble
BenchmarkSort
BenchmarkSort-4            10000          15261312 ns/op
PASS
ok      github.com/shanehowearth/algorithms/gga/bubble  152.616s
=== RUN   TestMerge
=== RUN   TestMerge/Book_example
=== RUN   TestMerge/Zero_length
=== RUN   TestMerge/Already_sorted
--- PASS: TestMerge (0.00s)
    --- PASS: TestMerge/Book_example (0.00s)
    --- PASS: TestMerge/Zero_length (0.00s)
    --- PASS: TestMerge/Already_sorted (0.00s)
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/gga/merge
BenchmarkSort
BenchmarkSort-4            10000            390577 ns/op
PASS
ok      github.com/shanehowearth/algorithms/gga/merge   3.909s
=== RUN   TestQuick
=== RUN   TestQuick/Book_example
=== RUN   TestQuick/Zero_length
=== RUN   TestQuick/Already_sorted
--- PASS: TestQuick (0.00s)
    --- PASS: TestQuick/Book_example (0.00s)
    --- PASS: TestQuick/Zero_length (0.00s)
    --- PASS: TestQuick/Already_sorted (0.00s)
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/gga/quick
BenchmarkSort
BenchmarkSort-4            10000          10168884 ns/op
PASS
ok      github.com/shanehowearth/algorithms/gga/quick   101.691s
```
