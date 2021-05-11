# Trie
I<a href="https://en.wikipedia.org/wiki/Trie">Wikipedia</a>: "In computer science, a trie, also called digital tree or prefix
tree, is a type of search tree, a tree data structure used for locating specific
keys from within a set. These keys are most often strings, with links between
nodes defined not by the entire key, but by individual characters. In order to
access a key (to recover its value, change it, or remove it), the trie is
traversed depth-first, following the links between nodes, which represent each
character in the key.

Unlike a binary search tree, nodes in the trie do not store their associated
key. Instead, a node's position in the trie defines the key with which it is
associated. This distributes the value of each key across the data structure,
and means that not every node necessarily has an associated key.

All the children of a node have a common prefix of the string associated with
that parent node, and the root is associated with the empty string. This task of
storing data accessible by its prefix can be accomplished in a memory-optimized
way by employing a radix tree.

Though tries can be keyed by character strings, they need not be. The same
algorithms can be adapted for ordered lists of any underlying type, e.g.
permutations of digits or shapes. In particular, a bitwise trie is keyed on the
individual bits making up a piece of fixed-length binary data, such as an
integer or memory address."

# Metrics
<table>
    <tr>
        <th>Name</th>
    </tr>
    <tr>
        <td><a
        href="https://en.wikipedia.org/wiki/C-trie"
        title="C-trie, Wikipedia">C-trie (aka Compressed-ADT)</a></td>
    </tr>
    <tr>
        <td><a
        href="https://en.wikipedia.org/wiki/Ctrie"
        title="Ctrie, Wikipedia">Ctrie</a></td>
    </tr>
    <tr>
        <td><a
        href="https://en.wikipedia.org/wiki/Hash_tree_(persistent_data_structure)"
        title="Hash trie, Wikipedia">Hash trie</a></td>
    </tr>
    <tr>
        <td><a
        href="https://en.wikipedia.org/wiki/Radix_tree"
        title="Radix trie, Wikipedia">Radix trie</a></td>
    </tr>
    <tr>
        <td><a
        href="https://en.wikipedia.org/wiki/Suffix_tree
        title="Suffix tree, Wikipedia">Suffix Tree</a></td>
    </tr>
    <tr>
        <td><a
        href="https://en.wikipedia.org/wiki/Ternary_search_tree
        title="Ternary search tree, Wikipedia">Ternary search tree</a></td>
    </tr>
    <tr>
        <td><a
        href="https://en.wikipedia.org/wiki/X-fast_trie"
        title="X-Fast trie, Wikipedia">X-Fast trie</a></td>
    </tr>
    <tr>
        <td><a
        href="https://en.wikipedia.org/wiki/Y-fast_trie"
        title="Y-Fast trie, Wikipedia">Y-Fast trie</a></td>
    </tr>
</table>
