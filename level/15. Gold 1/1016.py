from math import *

a, b = map(int, input().split())
c = [1 for _ in range(a, b + 1)]
r = []

d = int(sqrt(b))
e = [int(pow(i, 2)) for i in range(2, d + 1)]
for i in e:
    f = (ceil(a / i) * i) - a
    while f <= b - a:
        c[f] = 0
        f += i

print(sum(c))