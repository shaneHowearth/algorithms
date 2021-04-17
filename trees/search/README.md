# Search
Wikipedia: "In computer science, a search tree is a tree data structure used for
locating specific keys from within a set. In order for a tree to function as a
search tree, the key for each node must be greater than any keys in subtrees on
the left, and less than any keys in subtrees on the right."

# Metrics
<table>
    <tr>
        <th>Name</th>
        <th colspan="2" style="text-align:center">Space</th>
        <th colspan="2" style="text-align:center">Search</th>
        <th colspan="2" style="text-align:center">Insert</th>
        <th colspan="2" style="text-align:center">Delete</th>
    </tr>
    <tr>
        <td></td>
        <th style="text-align:center">Average</th>
        <th style="text-align:center">Worst Case</th>
        <th style="text-align:center">Average</th>
        <th style="text-align:center">Worst Case</th>
        <th style="text-align:center">Average</th>
        <th style="text-align:center">Worst Case</th>
        <th style="text-align:center">Average</th>
        <th style="text-align:center">Worst Case</th>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/2%E2%80%933_tree"
        title="2-3 tree, Wikipedia">2-3 tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/2%E2%80%933%E2%80%934_tree"
        title="2-3-4 tree, Wikipedia">2-3-4 tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/AA_tree"
        title="AA tree, Wikipedia">AA tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/(a,b)-tree"
        title="AB tree, Wikipedia">AB tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/AVL_tree"
        title="AVL tree, Wikipedia">AVL tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/B-tree"
        title="B tree, Wikipedia">B tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/B%2B_tree"
        title="B+ tree, Wikipedia">B+ tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(log n + log L)</td>
        <td>O(log n)</td>
        <td>O(M*log n + log L)</td>
        <td>O(log n)</td>
        <td>O(M*log n + log L)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/B-tree#Variants"
        title="B* tree, Wikipedia">B* tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Bx-tree"
        title="Bˣ tree, Wikipedia">Bˣ tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Binary_search_tree"
        title="Binary tree, Wikipedia">Binary tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Dancing_tree"
        title="Dancing tree, Wikipedia">Dancing tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/HTree"
        title="H tree, Wikipedia">H tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Interval_tree"
        title="Interval tree, Wikipedia">Interval tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Order_statistic_tree"
        title="Order statistic tree, Wikipedia">Order statistic tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Red%E2%80%93black_tree"
        title="Wikipedia">Red Black</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Scapegoat_tree"
        title="Scapegoat tree, Wikipedia">Scapegoat tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
        <td>amortised O(log n)</td>
        <td>O(log n)</td>
        <td>amortised O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Splay_tree"
        title="Splay tree, Wikipedia">Splay tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>amortised O(log n)</td>
        <td>amortised O(log n)</td>
        <td>amortised O(log n)</td>
        <td>amortised O(log n)</td>
        <td>amortised O(log n)</td>
        <td>amortised O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/T-tree"
        title="T tree, Wikipedia">T tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Treap"
        title="Treap, Wikipedia">Treap</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(n)</td>
        <td>O(log n)</td>
        <td>O(n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/UB-tree"
        title="UB tree, Wikipedia">UB tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Weight-balanced_tree"
        title="Weight-balanced tree, Wikipedia">Weight-balanced tree</a></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
        <td></td>
    </tr>
</table>

## Benchmarks
```
goos: linux
goarch: amd64
pkg: github.com/shanehowearth/algorithms/trees/search/binary
cpu: Intel(R) Core(TM) i5-7500 CPU @ 3.40GHz
BenchmarkSearch
BenchmarkSearch-4         769651              7188 ns/op             590 B/op
4 allocs/op
BenchmarkInsertion
BenchmarkInsertion-4       10000           1274600 ns/op          480000 B/op
10000 allocs/op
BenchmarkDeletion
BenchmarkDeletion-4       100000               272.2 ns/op             0 B/op
0 allocs/op
PASS
```
