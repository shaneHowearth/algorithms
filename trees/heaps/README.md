# Heap
Wikipedia: "In computer science, a heap is a specialized tree-based data
structure which is essentially an almost complete[1] tree that satisfies the
heap property: in a max heap, for any given node C, if P is a parent node of C,
then the key (the value) of P is greater than or equal to the key of C. In a min
heap, the key of P is less than or equal to the key of C.[2] The node at the
"top" of the heap (with no parents) is called the root node.

The heap is one maximally efficient implementation of an abstract data type
called a priority queue, and in fact, priority queues are often referred to as
"heaps", regardless of how they may be implemented. In a heap, the highest (or
lowest) priority element is always stored at the root. However, a heap is not a
sorted structure; it can be regarded as being partially ordered. A heap is a
useful data structure when it is necessary to repeatedly remove the object with
the highest (or lowest) priority.

A common implementation of a heap is the binary heap, in which the tree is a
binary tree (see figure). The heap data structure, specifically the binary heap,
was introduced by J. W. J. Williams in 1964, as a data structure for the
heapsort sorting algorithm.[3] Heaps are also crucial in several efficient graph
algorithms such as Dijkstra's algorithm. When a heap is a complete binary tree,
it has a smallest possible height—a heap with N nodes and for each node a
branches always has logٰₐ N height."

# Metrics
- 2-3 Heap
- binary
- binomial
- brodal
- fibonacci
- leftist
- pairing
- skew
- van_emde_boas
- weak
<table>
    <tr>
        <th>Name</th>
        <th colspan="2" style="text-align:center">Space</th>
        <th colspan="2" style="text-align:center">Search</th>
        <th colspan="2" style="text-align:center">Insert</th>
        <th colspan="2" style="text-align:center">Find-min</th>
        <th colspan="2" style="text-align:center">Delete-min</th>
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
        <th style="text-align:center">Average</th>
        <th style="text-align:center">Worst Case</th>
    </tr>
    <tr>
        <td><a href=""https://en.wikipedia.org/wiki/2%E2%80%933_heap
        title="2-3 heap, Wikipedia">2-3 Heap</a></td>
        <td></td>
        <td></td>
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
        <td><a href="https://en.wikipedia.org/wiki/Binary_heap"
        title="Binary heap, Wikipedia">Binary Heap</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(1)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Binomial_heap"
        title="Binomial heap, Wikipedia">Binomial Heap</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(1)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Brodal_queue""
        title="Brodal heap, Wikipedia">Brodal Heap</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(1)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Fibonacci_heap"
        title="Fibonacci heap, Wikipedia">Fibonacci Heap</a></td>
        <td></td>
        <td></td>
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
        <td><a href="https://en.wikipedia.org/wiki/Leftist_tree"
        title="Leftist tree, Wikipedia">Leftist tree</a></td>
        <td></td>
        <td></td>
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
        <td><a href="https://en.wikipedia.org/wiki/Pairing_heap
        title="Pairing heap, Wikipedia">Pairing Heap</a></td>
        <td></td>
        <td></td>
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
        <td><a href="https://en.wikipedia.org/wiki/Skew_heap"
        title="Skew heap, Wikipedia">Skew Heap</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(1)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Binary_heap"
        title="Binary heap, Wikipedia">Binary Heap</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(1)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Van_Emde_Boas_tree
        title="Van Emde Boas tree, Wikipedia">Van Emde Boas tree</a></td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(n)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(1)</td>
        <td>O(1)</td>
        <td>O(log n)</td>
        <td>O(log n)</td>
    </tr>
    <tr>
        <td><a href="https://en.wikipedia.org/wiki/Weak_heap
        title="Weak heap, Wikipedia">Weak Heap</a></td>
        <td></td>
        <td></td>
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
