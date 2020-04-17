import sys
from math import *

maxv = 1000000

p = [True] * maxv
for i in range(2, ceil(sqrt(maxv))):
    if p[i]:
        for j in range(i + i, maxv, i):
            p[j] = False

while True:
    a = int(sys.stdin.readline())
    if a == 0:
        break
    for i in range(3, a//2+1):
        if p[i] and p[a-i]:
            print(f'{a} = {i} + {a-i}')
            break
