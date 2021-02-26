# Classic sorting algorithm examples
## Bubble Sort
* Worst Case Performance O(n²)
* Average Case Performance O(n²)
* Best Case Performance O(n), list is already sorted.

## Insertion Sort //TODO
* Worst Case Performance O(n²)
* Average Case Performance O(n²)
* Best Case Performance O(n), list is already sorted.
* Worst Case Space Complexity O(n) total, O(1) auxiliary.

## Heap Sort // TODO
* Worst Case Performance O(n log n)
* Average Case O(n log n)
* Best Case Performance O(n log n), distinct keys, O(n) equal keys.
* Worst Case Space Complexity O(n) total, O(1) auxiliary.

## Merge Sort
* Worst Case Performance O(n log n)
* Average Case O(n log n)
* Best Case Performance O(n log n), typical, O(n) O(n).
* Worst Case Space Complexity O(n) total, O(n) auxiliary (O(1) auxiliary with linked lists).

## Quicksort
* Worst Case Performance O(n²)
* Average Case Performance O(n log n)
* Best Case Performance O(n log n), simple partition, O(n) three way partition and equal keys.
* Worst Case Space Complexity O(n) auxiliary , (O(log n) Hoare, 1962).

## Benchmarks for existing algorithms
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
