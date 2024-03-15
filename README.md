Goed is a library written in Go, designed for calculating the edit distance (Levenshtein distance) across various data structures including strings, trees, and graphs. For each of these data structures, Goed offers both a sequential and a parallel version, allowing users to choose the most suitable approach based on their specific requirements and computational resources.

1. String Edit Distance:

  Goed implements a standard dynamic programming-based approach to calculate the edit distance between two strings. Its sequential version is based on the classic Wagner–Fischer algorithm. The parallel version employs wavefront-oriented, tile-based parallelism to adapt to the dynamic programming nature.
  
2. Tree Edit Distance:

4. Graph Edit Distance:
