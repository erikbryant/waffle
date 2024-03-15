# Find initial set of Possible Letters

The possible letters for a tile are given as the set:

  pl := g | w + yd - w(row) - w(col) + y(row) + y(col) - ws

The match instruction for a word is given as:

  mi := `[pl1][pl2][pl3][pl4][pl5] | egrep [ye]`

Where:

* g      = self is green
* w      = {all white tiles}
* yd     = {all yellow duplicates}
* w(row) = {all white tiles in this row}
* w(col) = {all white tiles in this col}
* y(row) = {all yellow tiles in this row}
* y(col) = {all yellow tiles in this col}
* s      = self
* ye     = any yellow letters in non-intersection tiles

# Reduce that set by any that have already been assigned

Once the sets of possible letters (pl) for each tile have been found,
take the set of all starting letters and subtract all letters that have
a known position (a set size of one). The remainder will be the letters
that can still be used. Any possible letter that is not in that set can
be eliminated.

# Reduce that set again by valid words

Map each word's sets of possible letters against a broad dictionary.
Reduce the 'pl' sets to the sets of letters in the matches.
