# PyPy3

import sys
from collections import defaultdict

a = int(sys.stdin.readline())
b = defaultdict(int)
max, min = -1000000, 1000001

for _ in range(a):
    c = int(sys.stdin.readline())
    b[c] += 1
    if c > max: max = c
    if c < min: min = c

for i in range(min, max + 1):
    if b[i] is None:
        continue
    else:
        for _ in range(b[i]):
            print(i)
