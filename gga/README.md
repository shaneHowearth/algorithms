# Classic sorting algorithm examples


| Algorithm | Worst Case | Average Case | Best Case | Worst Case Space Complexity | Done |
|:---:|:---:|:---:|:---:|:---|:---|
| Bubble Sort | O(n²) | O(n²) | O(n), list is already sorted. |||
| Insertion Sort | O(n²) | O(n²) | O(n), list is already sorted. | O(n) total, O(1) auxiliary. ||
| Heap Sort | O(n log n) | O(n log n) | O(n log n), distinct keys, O(n) equal keys.| O(n) total, O(1) auxiliary. ||
| Merge Sort | O(n log n) | Θ(n log n) | Ω(n log n), typical, Ω(n) | O(n) total, O(n) auxiliary (O(1) auxiliary with linked lists). ||
| Quicksort | O(n²) | O(n log n) | O(n log n), simple partition, O(n) three way partition and equal keys. | O(n) auxiliary , (O(log n) Hoare, 1962). ||

## Benchmarks
Note: Have fixed the benchmarks, but they are still running
``` shell
$ go test -bench . ./... -timeout 0 -benchmem
```
